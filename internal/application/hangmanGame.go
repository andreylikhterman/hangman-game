package hangman

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	models "github.com/backend-academy-2024-go-template/internal/domain"
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

var hints = map[string]string{
	"кот":           "Ловит мышей",
	"пёс":           "Друг человека",
	"лев":           "Царь зверей",
	"лось":          "Животное с рогами",
	"рак":           "Водное существо",
	"зебра":         "Полосатая лошадь",
	"акула":         "Морской хищник",
	"тигр":          "Большая кошка",
	"аист":          "Птица с длинными ногами",
	"медведь":       "Любит мёд",
	"гиппопотам":    "Огромное животное, живёт в воде",
	"ягуар":         "Пятнистая кошка",
	"коала":         "Ест эвкалипт, живёт на деревьях",
	"дикобраз":      "У него есть иглы для защиты",
	"дятел":         "Стучит по деревьям",
	"суп":           "Первое блюдо",
	"сок":           "Напиток",
	"сыр":           "Молочный продукт",
	"чай":           "Горячий напиток",
	"торт":          "Десерт на праздник",
	"пирог":         "Выпечка с начинкой",
	"омлет":         "Блюдо из яиц",
	"борщ":          "Красный суп",
	"салат":         "Нарезанные овощи",
	"оладьи":        "Жареные лепёшки",
	"спагетти":      "Макароны",
	"вареники":      "Тесто с начинкой",
	"гаспачо":       "Холодный суп",
	"равиоли":       "Маленькие изделия из теста",
	"тирамису":      "Десерт с кофе",
	"плеер":         "С помощью него слушают музыку",
	"утюг":          "Разглаживает одежду",
	"лампа":         "Есть на каждом столе",
	"миксер":        "Смешивает ингредиенты",
	"пылесос":       "Чистит полы",
	"планшет":       "Устройство с экраном",
	"мотор":         "Главная часть в машине",
	"принтер":       "Печатает документы",
	"камера":        "Снимает фото и видео",
	"дрель":         "Сверлит отверстия",
	"холодильник":   "Хранит продукты",
	"фотоаппарат":   "Делает снимки",
	"кондиционер":   "Охлаждает воздух",
	"электромобиль": "Машина",
	"навигатор":     "Определяет маршрут",
}

var categories = []string{"животные", "еда", "техника"}
var levels = []string{"легкий", "средний", "сложный"}

type Game struct {
	words   map[string]map[string][]string
	hints   map[string]string
	word    models.Word
	player  models.Player
	windows models.Windows
}

func NewGame(attempts int) *Game {
	game := &Game{
		words: dictionary,
		hints: hints,
		player: models.Player{
			CountAttempts: attempts,
		},
		windows: models.Windows{
			HangmanStages: models.Stages,
		},
	}
	game.player.Window = &(game.windows)

	return game
}

func RandomElememt(array []string) (index int64, element string) {
	index = time.Now().Unix() % int64(len(array))
	return index, array[index]
}

func (game *Game) SelectCategory(category int) string {
	var index int64
	if category > 0 && category <= len(categories) {
		index = int64(category - 1)
	} else {
		index, _ = RandomElememt(categories)
	}

	game.player.CountAttempts += int(1 - index)

	return categories[index]
}

func (game *Game) SelectLevel(level int) string {
	if level > 0 && level <= len(levels) {
		return levels[level-1]
	}

	index, _ := RandomElememt(levels)

	return levels[index]
}

func contains(array []rune, element rune) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}

	return false
}

func (game *Game) selectLevelAndCategory() {
	game.windows.SelectLevel()
	level := game.player.ChooseLevel()
	game.windows.SelectCategory()
	category := game.player.ChooseCategory()

	correctLevel := game.SelectLevel(level)
	correctCategory := game.SelectCategory(category)

	_, randomWord := RandomElememt(game.words[correctCategory][correctLevel])
	game.word = game.word.NewWord(randomWord, correctCategory, correctLevel)
}

func (game *Game) displayInitialData() {
	word := game.word.Word
	game.windows.MainWindow()
	game.windows.Cursor.ToWord()
	fmt.Print(strings.Repeat("_", len(word)))
	game.windows.Cursor.ToAttempts()
	fmt.Print(game.player.CountAttempts)
	game.windows.Cursor.ToInput()
}

func (game *Game) handleUserInput(char string) {
	game.windows.Cursor.ToInput()
	game.windows.ClearText(char)

	if len([]rune(char)) != 1 {
		game.windows.Cursor.ToInput()
		return
	}

	letter, _ := utf8.DecodeRuneInString(strings.ToLower(char))

	game.windows.Cursor.ToWord()

	switch {
	case contains(game.word.Word, letter):
		game.word.UpdateGuessedLetters(letter)
	case char == "?":
		hint := game.hints[string(game.word.Word)]
		game.windows.ShowHint(hint)
	default:
		game.decrementAttempts()
	}

	game.windows.CrossOutLetter(letter)
	game.windows.Cursor.ToInput()
}

func (game *Game) decrementAttempts() {
	game.player.CountAttempts--
	game.windows.Cursor.ToAttempts()
	fmt.Print(game.player.CountAttempts)
	game.windows.DrawHangman(game.player.CountAttempts)
}

func (game *Game) endGame() {
	if game.player.CountAttempts == 0 {
		game.windows.Loss()
	} else {
		game.windows.Win()
	}

	time.Sleep(2 * time.Second)

	game.windows.CleanScreen()
}

func (game *Game) Play() {
	game.windows.Start()
	game.selectLevelAndCategory()
	game.displayInitialData()

	for game.player.CountAttempts > 0 && int(game.word.CountGuessedLetters) != len(game.word.Word) {
		var char string
		if _, err := fmt.Scan(&char); err != nil {
			fmt.Print("")
		}

		game.handleUserInput(char)
	}

	game.endGame()
}
