package prompt

import "github.com/HazelnutParadise/Go-Utils/conv"

var communityInfoRootMap = map[string]string{
	"社區名稱": "新北市永和區民權社區",
	"產品":   conv.ToString([]map[string]string{oneBiteCrispyInfo}),
	"範圍":   conv.ToString([4]string{"民族里", "民權里", "民治里", "智光里"}),
	"社區發展協會": conv.ToString(map[string]string{
		"地址":   "新北市永和區民族里民權路60號4樓",
		"電話號碼": "02-29433377",
		"開放時間": "周一至周五的8:30到16:00",
		"備註":   "民權社區發展協會同時也是市民活動中心",
	}),
	"舉辦活動的地點": conv.ToString([2]string{"民族公園（地址：新北市永和區民族街33號）", "社區發展協會"}),
	"活動":      conv.ToString(activitiesInfo),
	"獎項":      conv.ToString(awardsInfo),
	"官方網站":    "https://timlai666.github.io/one-bite-crispy/",
	"粉絲專頁":    "https://www.facebook.com/profile.php?id=100069389774627&locale=zh_TW",
}
