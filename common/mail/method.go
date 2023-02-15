package mail

import (
	"bytes"
	"embed"
	"html/template"
)

//go:embed templates/*.html
var templates embed.FS

var parsedTemplates = template.Must(template.ParseFS(templates, "templates/*.html"))

func ParseTemplates[T EmailVerification | PasswordReset](data T, templateTitle string) (string, error) {
	buff := new(bytes.Buffer)

	err := parsedTemplates.ExecuteTemplate(buff, templateTitle, data)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
