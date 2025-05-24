package service

import (
	"bot/llm"
	"bot/prompt"

	"github.com/HazelnutParadise/Go-Utils/conv"
	"github.com/TimLai666/zaipra"
	"github.com/tmc/langchaingo/llms"
)

func AskInfo(msg string) (string, error) {
	sysPrompt, err := prompt.ReplySystemPrompt(msg)
	if err != nil {
		return "", err
	}

	return zaipra.Answer("民眾傳來的訊息", msg, sysPrompt, []zaipra.Info{
		{Title: "社區活動資訊", Description: "好玩的活動、講座等", Content: conv.ToString(prompt.ActivitiesInfoMap)},
		{Title: "社區故事和特色", Description: "社區長者的故事和習俗特色", Content: conv.ToString(prompt.StoriesInfoMap)},
		{Title: "社區獎項資訊", Description: "社區獲得的獎項", Content: conv.ToString(prompt.AwardsInfoMap)},
	}, llm.LLMClient, llms.WithTemperature(1.2))
}
