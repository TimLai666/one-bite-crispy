package llm

import (
	"bot/config"

	"github.com/tmc/langchaingo/llms/ollama"
)

var (
	LLMClient *ollama.LLM
	model     string = "gemma3:12b"
)

func init() {
	llm, err := ollama.New(ollama.WithServerURL(config.OLLAMA_BASE_URL), ollama.WithModel(model))
	if err != nil {
		panic(err)
	}
	LLMClient = llm
}
