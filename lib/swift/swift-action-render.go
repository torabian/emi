package swift

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func SwiftActionRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {
	realms, deps, err := GetActionRealms(action, ctx, complexes)
	if err != nil {
		return nil, err
	}

	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  core.TOKEN_ORIGINAL_NAME,
				Value: realms.ActionName,
			},
		},
	}

	const tmpl = `import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif

/**
 Action to communicate with the action {{ .realms.ActionName }}
 */
struct {{ .realms.ActionName }}Meta {
    let name: String = "{{ .realms.ActionName }}"
    let url: String = "{{ .action.Url }}"
    let method: String = "{{ .action.Method | upper }}"
}

/*
 struct {{ .realms.ActionName }}Request {
     // reserved
 }
 */

struct {{ .realms.ActionName }}Response {
    let statusCode: Int
    let headers: [String: String]
    let payload: Data?

    init(
        statusCode: Int = 200,
        headers: [String: String] = [:],
        payload: Data? = nil
    ) {
        self.statusCode = statusCode
        self.headers = headers
        self.payload = payload
    }
}

{{ if .realms.PathParameter }}
{{ b2s .realms.PathParameter.ActualScript }}
{{ end }}

final class {{ .realms.ActionName }}Client {

    private static let session: URLSession = .shared

    private static func buildUrl(
        base: String,
        path: String,
        query: [String: String]
    ) -> URL? {
        guard var components = URLComponents(string: base) else {
            return nil
        }

        components.path = path
        components.queryItems = query.map {
            URLQueryItem(name: $0.key, value: $0.value)
        }

        return components.url
    }

    static func compute(
        {{ if .realms.PathParameter }}
        path: {{ .realms.ActionName }}PathParameter,
        {{ end }}
        query: [String: String] = [:],
        headers: [String: String] = [:],
        body: Data? = nil
    ) async throws -> {{ .realms.ActionName }}Response {

        let meta = {{ .realms.ActionName }}Meta()
        let baseUrl = ""

        guard var url = buildUrl(
            base: baseUrl,
            path: meta.url,
            query: query
        ) else {
            throw URLError(.badURL)
        }

        {{ if .realms.PathParameter }}
        url = {{ .realms.ActionName }}PathParameterApply(path, url)
        {{ end }}

        var request = URLRequest(url: url)
        request.httpMethod = meta.method
        request.httpBody = body
        request.setValue("application/json", forHTTPHeaderField: "Accept")

        headers.forEach {
            request.setValue($0.value, forHTTPHeaderField: $0.key)
        }

        let (data, response) = try await session.data(for: request)

        guard let http = response as? HTTPURLResponse else {
            throw URLError(.badServerResponse)
        }

        return {{ .realms.ActionName }}Response(
            statusCode: http.statusCode,
            headers: http.allHeaderFields.reduce(into: [:]) { acc, item in
                if let k = item.key as? String,
                   let v = item.value as? String {
                    acc[k] = v
                }
            },
            payload: data
        )
    }
}

{{ if .realms.RequestClass }}
{{ b2s .realms.RequestClass.ActualScript }}
{{ end }}

{{ if .realms.ResponseClass }}
{{ b2s .realms.ResponseClass.ActualScript }}
{{ end }}
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action": action,
		"realms": realms,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = realms.ActionName
	res.SuggestedExtension = ".swift"
	res.CodeChunkDependensies = deps

	return res, nil
}
