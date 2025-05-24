package llm

import (
	"bot/config"
	"context"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

var (
	client *ollama.LLM
	model  string = "gemma3:12b"
)

func init() {
	llm, err := ollama.New(ollama.WithServerURL(config.OLLAMA_BASE_URL), ollama.WithModel(model))
	if err != nil {
		panic(err)
	}
	client = llm
}

func CallLLM(prompt string) (string, error) {
	prompt = "使用繁體中文回答！\n禁止使用粗話或髒話！\n" + prompt
	return llms.GenerateFromSinglePrompt(context.Background(), client, prompt, llms.WithTemperature(1.5))
}
