package prompt

import (
	"bytes"
	"log"
	"text/template"
	"time"

	"github.com/HazelnutParadise/Go-Utils/conv"
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

func ReplySystemPrompt(msg string) (string, error) {
	var tpl bytes.Buffer
	now := time.Now()
	if err := tmpl.Execute(&tpl, struct {
		CommunityBaseInfo string
		Today             string
		NowTime           string
	}{CommunityBaseInfo: conv.ToString(CommunityBaseInfoMap), Today: now.Format("2006/01/02"), NowTime: now.Format("15:04")}); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
