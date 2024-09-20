package hangman

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

type HangmanGame struct {
	words   map[string]map[string][]string
	hints   map[string]string
	word    models.Word
	player  models.Player
	windows models.HangmanWindows
}

func NewHangmanGame(attempts int) *HangmanGame {
	return &HangmanGame{
		words: dictionary,
		hints: hints,
		player: models.Player{
			CountAttempts: attempts,
		},
		windows: models.HangmanWindows{
			HangmanStages: models.Stages,
		},
	}
}

func RandomElememt(array []string) (int64, string) {
	index := time.Now().Unix() % int64(len(array))
	return index, array[index]
}

func (game *HangmanGame) SelectCategory(category string) string {
	var index int64
	if number, err := strconv.Atoi(category); err == nil && number > 0 && number <= len(categories) {
		index = int64(number - 1)
	} else {
		index, _ = RandomElememt(categories)
	}
	game.player.CountAttempts += int(1 - index)
	return categories[index]
}

func (game *HangmanGame) SelectLevel(level string) string {
	if number, err := strconv.Atoi(level); err == nil && number > 0 && number <= len(levels) {
		return levels[number-1]
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

func (game *HangmanGame) selectLevelAndCategory() {
	game.windows.SelectLevel()
	level := game.player.ChooseLevel()
	game.windows.SelectCategory()
	category := game.player.ChooseCategory()

	level = game.SelectLevel(level)
	category = game.SelectCategory(category)

	_, randomWord := RandomElememt(game.words[category][level])
	game.word = game.word.NewWord(randomWord, level, category)
}

func (game *HangmanGame) displayInitialData() {
	word := []rune(game.word.Word)
	game.windows.MainWindow()
	game.windows.Cursor.ToWord()
	fmt.Print(strings.Repeat("_", len(word)))
	game.windows.Cursor.ToAttempts()
	fmt.Print(game.player.CountAttempts)
	game.windows.Cursor.ToInput()
}

func (game *HangmanGame) handleUserInput(char string) {
	game.windows.Cursor.ToInput()
	game.windows.ClearText(char)

	if len([]rune(char)) != 1 {
		game.windows.Cursor.ToInput()
		return
	}

	letter := []rune(strings.ToLower(char))[0]
	game.windows.Cursor.ToWord()

	if contains(game.word.Word, letter) {
		game.word.UpdateGuessedLetters(letter)
	} else if char == "?" {
		hint := game.hints[string(game.word.Word)]
		game.windows.ShowHint(hint)
	} else {
		game.decrementAttempts()
	}

	game.windows.CrossOutLetter(letter)
	game.windows.Cursor.ToInput()
}

func (game *HangmanGame) decrementAttempts() {
	game.player.CountAttempts--
	game.windows.Cursor.ToAttempts()
	fmt.Print(game.player.CountAttempts)
	game.windows.DrawHangman(game.player.CountAttempts)
}

func (game *HangmanGame) endGame() {
	if game.player.CountAttempts == 0 {
		game.windows.Loss()
	} else {
		game.windows.Win()
	}
	time.Sleep(2 * time.Second)
	game.windows.CleanScreen()
}

func (game *HangmanGame) Play() {
	game.windows.Start()
	game.selectLevelAndCategory()
	game.displayInitialData()

	for game.player.CountAttempts > 0 && game.word.CountGuessedLetters != int8(len(game.word.Word)) {
		var char string
		fmt.Scan(&char)
		game.handleUserInput(char)
	}

	game.endGame()
}
