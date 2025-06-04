package hangman_test

import (
	"testing"

	game "hangman/internal/application"

	"github.com/stretchr/testify/assert"
)

var dictionary = map[string]map[string][]string{
	"животные": {
		"легкий":  []string{"кот", "пёс", "лев", "лось", "рак"},
		"средний": []string{"зебра", "акула", "тигр", "аист", "медведь"},
		"сложный": []string{"гиппопотам", "ягуар", "коала", "дикобраз", "дятел"},
	},
	"еда": {
		"легкий":  []string{"суп", "сок", "сыр", "чай", "торт"},
		"средний": []string{"пирог", "омлет", "борщ", "салат", "оладьи"},
		"сложный": []string{"спагетти", "вареники", "гаспачо", "равиоли", "тирамису"},
	},
	"техника": {
		"легкий":  []string{"плеер", "утюг", "лампа", "миксер", "пылесос"},
		"средний": []string{"планшет", "мотор", "принтер", "камера", "дрель"},
		"сложный": []string{"холодильник", "фотоаппарат", "кондиционер", "электромобиль", "навигатор"},
	},
}

var categories = []string{"животные", "еда", "техника"}
var levels = []string{"легкий", "средний", "сложный"}

func TestSelectCategory(t *testing.T) {
	game := game.NewGame(5)

	category := game.SelectCategory(1, nil)
	assert.Equal(t, "животные", category, "Ожидалась категория 'животные'")

	category = game.SelectCategory(13, nil)
	assert.Contains(t, categories, category, "Выбрана неверная категория")
}

func TestSelectLevel(t *testing.T) {
	game := game.NewGame(5)

	level := game.SelectLevel(2, nil)
	assert.Equal(t, "средний", level, "Ожидался уровень 'средний'")

	level = game.SelectLevel(14, nil)
	assert.Contains(t, levels, level, "Выбран неверный уровень")
}
