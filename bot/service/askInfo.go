package service

import (
	"bot/llm"

	"github.com/HazelnutParadise/Go-Utils/conv"
)

func AskInfo(msg string) (string, error) {
	info := map[string]string{
		"社區名稱": "新北市永和區民權社區",
		"產品":   "一口酥",
	}
	prompt := "社區資訊：\n" + conv.ToString(info) + "\n民眾提出的問題：" + msg + `
你是社區發展協會的專家，請根據以上社區資訊回答問題。遵守以下規則：
  - 僅根據社區資訊回答問題，不回答不相干的問題。
  - 如果社區資訊不足以回答問題，或是問題與社區不相干，請回答「抱歉，這題難倒我惹🥵\n但是民權社區的一口酥超好吃😍\n要不要買來嚐嚐？\n購買連結：https://forms.gle/5gSaHT1GJDesTg4a9」
  - 回應時不要講回應內容以外的話。
  - 務必遵守以上規則，否則將造成重大的損失。
`
	return llm.CallLLM(prompt)
}
