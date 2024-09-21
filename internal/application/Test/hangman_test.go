package hangman_test

import (
	"testing"

	gamepkg "github.com/backend-academy-2024-go-template/internal/application"
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
	game := gamepkg.NewGame(5)

	category := game.SelectCategory(1)
	assert.Equal(t, "животные", category, "Ожидалась категория 'животные'")

	category = game.SelectCategory(13)
	assert.Contains(t, categories, category, "Выбрана неверная категория")
}

func TestSelectLevel(t *testing.T) {
	game := gamepkg.NewGame(5)

	level := game.SelectLevel(2)
	assert.Equal(t, "средний", level, "Ожидался уровень 'средний'")

	level = game.SelectLevel(14)
	assert.Contains(t, levels, level, "Выбран неверный уровень")
}

func TestSelectWord(t *testing.T) {
	gamepkg.NewGame(5)

	category := "животные"
	level := "легкий"

	_, randomWord := gamepkg.RandomElememt(dictionary[category][level])
	assert.Contains(t, dictionary[category][level], randomWord, "Слово не принадлежит к категории и уровню")
}
