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
	{Emoji: "😊", Content: "Улыбайся своему партнеру"},
	{Emoji: "💪", Content: "Держи рамку"},
	{Emoji: "🥁", Content: "Ходи в тайминг"},
	{Emoji: "👩‍🎓", Content: "Не пропускай занятия"},
	{Emoji: "🐣", Content: "Приглашай начинашек"},
	{Emoji: "🏆", Content: "Танцуй на конурсах хорошо, плохо не танцуй"},
	{Emoji: "🏋️‍♀️", Content: "Чаще ходи на сампо"},
	{Emoji: "💃", Content: "Больше танцуй на вечеринках"},
	{Emoji: "📚", Content: "Отрабатывай материалы индив"},
	{Emoji: "🙅‍♂️", Content: "Веди корпусом, а не руками"},
	{Emoji: "📽", Content: "Вдохновляйся танцевальными видео"},
	{Emoji: "🎭", Content: "Поставь, наконец, рутину"},
	{Emoji: "💅🏽", Content: "Выгляди достойно на танцполе"},
	{Emoji: "🤝", Content: "Обещал танцевать — танцуй"},
	{Emoji: "🎸", Content: "Блюз ты можешь не любить, но танцевать его обязан"},
	{Emoji: "🤸‍♀️", Content: "Не делай акробатику в ДнД"},
	{Emoji: "🚶‍♀️", Content: "Выйди, наконец, из новисов"},
	{Emoji: "🕺", Content: "Одевайся красиво на конкурс"},
	{Emoji: "⭐️", Content: "Выйди в оллстары"},
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetAdvice returns a random advice
func GetAdvice() Advice {
	return advices[rand.Intn(len(advices))]
}
