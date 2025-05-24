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
	now := time.Now()
	if err := tmpl.Execute(&tpl, struct {
		Info    map[string]string
		Msg     string
		Today   string
		NowTime string
	}{Info: communityInfoRootMap, Msg: msg, Today: now.Format("2006/01/02"), NowTime: now.Format("15:04")}); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
