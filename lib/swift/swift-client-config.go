package swift

import "github.com/torabian/emi/lib/core"

// emiClientConfigSource is the one piece of runtime configuration every generated
// action's Client.compute() needs and none of them had a way to receive: the actual
// server host to call. Every {{ActionName}}Client previously hardcoded `let baseUrl = ""`
// - meaning every request resolved against an empty host and the generated SDK could
// never actually reach a real server, regardless of how correct its types were. A caller
// sets EmiClientConfig.baseUrl once (e.g. at app startup); every generated client reads
// it from here instead.
const emiClientConfigSource = `/**
 * EmiClientConfig - shared configuration for every generated action client. Set
 * EmiClientConfig.baseUrl once (e.g. at app startup) before calling any action.
 */
enum EmiClientConfig {
	static var baseUrl: String = ""
}
`

// SwiftClientConfigFile is EmiClientConfig rendered as a standalone generated file -
// unconditionally appended by SwiftFullModule (see swift-public-api.go), same as
// SwiftAnyCodableFile.
func SwiftClientConfigFile() core.VirtualFile {
	return core.VirtualFile{
		Name:         "EmiClientConfig",
		Extension:    ".swift",
		ActualScript: emiClientConfigSource,
	}
}
