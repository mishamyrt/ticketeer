package hook

import (
	"bytes"
	"embed"
	"html/template"
	"runtime"
)

//go:embed *
var templatesFS embed.FS

type hookTmplData struct {
	Extension string
}

// Content returns hook content
func Content() []byte {
	buf := &bytes.Buffer{}
	t := template.Must(template.ParseFS(templatesFS, Name+".tmpl"))
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
