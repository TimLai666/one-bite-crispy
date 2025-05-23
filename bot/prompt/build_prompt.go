package prompt

import (
	"bytes"
	"log"
	"text/template"
	"time"
)

var tmpl *template.Template

func init() {
	filepath := "prompt/base_prompt.md"
	var err error
	tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		log.Fatalf("failed to parse template file: %v", err)
	}
}

func ReplyPrompt(msg string) (string, error) {
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, struct {
		Info    map[string]string
		Msg     string
		NowTime string
	}{Info: communityInfoRootMap, Msg: msg, NowTime: time.Now().Format("2006/01/02")}); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
