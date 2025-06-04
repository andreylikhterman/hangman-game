package hangman

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"hangman/internal/domain"
	"hangman/pkg/random"
	"hangman/pkg/slice"
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
	word    domain.Word
	player  domain.Player
	windows domain.Windows
}

func NewGame(attempts int) *Game {
	game := &Game{
		words: dictionary,
		hints: hints,
		player: domain.Player{
			CountAttempts: attempts,
		},
		windows: domain.Windows{
			HangmanStages: domain.Stages,
		},
	}

	return game
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

func (game *Game) SelectCategory(category int, err error) string {
	if err == nil && category > 0 && category <= len(categories) {
		return categories[category-1]
	}

	_, element := random.Elememt(categories)

	return element
}

func (game *Game) SelectLevel(level int, err error) string {
	var index int64
	if err == nil && level > 0 && level <= len(levels) {
		index = int64(level) - 1
	} else {
		index, _ = random.Elememt(levels)
	}

	game.player.CountAttempts += int(1 - index)

	return levels[index]
}

func (game *Game) selectLevelAndCategory() {
	game.windows.SelectLevel()
	level, err := game.player.ChooseLevel()
	correctLevel := game.SelectLevel(level, err)
	game.windows.SelectCategory()
	category, err := game.player.ChooseCategory()
	correctCategory := game.SelectCategory(category, err)

	_, randomWord := random.Elememt(game.words[correctCategory][correctLevel])
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
	case slice.Contains(game.word.Word, letter):
		game.word.UpdateGuessedLetters(letter)
	case letter == '?':
		hint := game.hints[string(game.word.Word)]
		game.windows.ShowHint(hint)
	case letter-'а' >= 0 && letter-'а' <= 32:
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
