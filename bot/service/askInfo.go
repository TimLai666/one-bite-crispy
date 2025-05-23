package service

import (
	"bot/llm"
	"bot/prompt"
)

func AskInfo(msg string) (string, error) {
	prompt, err := prompt.ReplyPrompt(msg)
	if err != nil {
		return "", err
	}
	return llm.CallLLM(prompt)
}
