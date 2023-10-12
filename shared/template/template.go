package template

import (
	"io"
	"os"
)

// Template is a string template. It may contain variable values represented as ${var} or $var that can be
// later replaced with values via Build method.
type Template struct {
	content string
}

// FromFile creates a Template from a given file.
func FromFile(path string) (*Template, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return &Template{
		content: string(content),
	}, nil
}

// Content returns the raw content of the Template. Can be used safely in case the template source does not
// contain any variable values.
func (t *Template) Content() string {
	return t.content
}

type Fields map[string]string

// Build replaces all variable values in the template string with the fields provided. Variables not mapped in
// the fields should be replaced with empty strings.
func (t *Template) Build(fields Fields) string {
	return os.Expand(t.content, func(s string) string {
		return fields[s]
	})
}
