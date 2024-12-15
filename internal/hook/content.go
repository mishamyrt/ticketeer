package hook

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"runtime"
)

//go:embed template/*
var templatesFS embed.FS

type hookTmplData struct {
	Extension string
}

func getTemplatePath() string {
	return fmt.Sprintf("template/%s.tmpl", Name)
}

// Content returns hook content
func Content() []byte {
	buf := &bytes.Buffer{}
	t := template.Must(template.ParseFS(templatesFS, getTemplatePath()))
	err := t.Execute(buf, hookTmplData{
		Extension: getExtension(),
	})
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func getExtension() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}
	return ""
}
