package advices

import (
	"math/rand"
	"time"
)

// Advice - Emoji and Content
type Advice struct {
	Emoji   string
	Content string
}

var advices = []Advice{
	{
		Emoji:   "😊",
		Content: "Улыбайся своему партнеру",
	},
	{
		Emoji:   "💪",
		Content: "Держи рамку",
	},
	{
		Emoji:   "🥁",
		Content: "Ходи в тайминг",
	},
	{
		Emoji:   "👩‍🎓",
		Content: "Не пропускай занятия",
	},
	{
		Emoji:   "🐣",
		Content: "Приглашай начинашек",
	},
	{
		Emoji:   "🏆",
		Content: "Танцуй на конурсах хорошо, плохо не танцуй",
	},
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetAdvice returns a random advice
func GetAdvice() Advice {
	return advices[rand.Intn(len(advices))]
}
