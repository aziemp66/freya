package mail

import (
	"bytes"
	"embed"
	"html/template"
)

//go:embed templates/*.html
var templates embed.FS

var parsedTemplates = template.Must(template.ParseFS(templates, "templates/*.html"))

func RenderEmailVerificationTemplate(token string) (string, error) {
	var buf bytes.Buffer
	err := parsedTemplates.ExecuteTemplate(&buf, "email_verification.html", map[string]string{
		"Token": token,
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func RenderPasswordResetTemplate(token string) (string, error) {
	var buf bytes.Buffer
	err := parsedTemplates.ExecuteTemplate(&buf, "reset_password.html", map[string]string{
		"Token": token,
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
