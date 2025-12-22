package querypredict

import (
	"embed"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func LoadQueriesFromReader(r io.Reader) (*QueryDocument, error) {
	var queries *QueryDocument
	decoder := yaml.NewDecoder(r)
	if err := decoder.Decode(&queries); err != nil {
		return nil, err
	}
	return queries, nil
}

func LoadQueriesFromFile(path string) (*QueryDocument, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LoadQueriesFromReader(f)
}

func LoadEmbeddedQuery(embed embed.FS, name string) (*QueryDocument, error) {
	f, err := embed.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LoadQueriesFromReader(f)
}

func LoadQueriesFromString(s string) (*QueryDocument, error) {
	return LoadQueriesFromReader(strings.NewReader(s))
}
