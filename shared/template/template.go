package template

import (
	"io"
	"os"
)

type Template struct {
	content string
}

type Fields map[string]string

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

func (t *Template) Build(fields Fields) string {
	return os.Expand(t.content, func(s string) string {
		return fields[s]
	})
}
