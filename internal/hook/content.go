package hook

import (
	"bytes"
	"embed"
	"html/template"
	"runtime"
)

//go:embed template/*
var templatesFS embed.FS

const templatePath = "template/prepare-commit-msg.tmpl"

type hookTmplData struct {
	Extension string
}

// Content returns hook content
func Content() ([]byte, error) {
	buf := &bytes.Buffer{}
	t := template.Must(template.ParseFS(templatesFS, templatePath))
	err := t.Execute(buf, hookTmplData{
		Extension: getExtension(),
	})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func getExtension() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}
	return ""
}
