package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Word struct {
	Word string
	Pos  int
}

// filterWords Фильтрует текст, заменяя цензурные и повторяющиеся слова
func filterWords(text string, censorMap map[string]string) string {
	// Разделение текста на предложения с помощью splitSentences
	sentences := splitSentences(text)
	// Если предложений больше одного, то обработка каждого предложения рекурсивно
	if len(sentences) > 1 {
		// Обработка каждого предложения в цикле
		for i := 0; i < len(sentences); i++ {
			// Рекурсивный вызов функции filterWords
			sentences[i] = filterWords(sentences[i], censorMap)
		}
		// Прерывание блока условия "если предложений больше одного" c помощью return strings.Join(sentences, " ")
		return strings.Join(sentences, " ")
	}
	// Разделение текста на отдельные слова с помощью strings.Fields(text)
	words := strings.Fields(text)
	// Создание пустой карты уникальных слов с помощью make(map[string]Word)
	unique := make(map[string]Word)
	// Обработка каждого слова в цикле
	// Обработка каждого слова в цикле
	for i := 0; i < len(words); i++ {
		// Если слово содержится в карте цензурных слов, то
		val1, ok1 := censorMap[strings.ToLower(words[i])]

		if ok1 == true {
			// Замена слова на значение из карты, используя CheckUpper
			words[i] = CheckUpper(words[i], val1)
		}
		// Если слово не содержится в карте уникальных слов, то (для проверки ключа в карте уникальных слов, используйте функцию strings.ToLower)
		if _, ok2 := unique[strings.ToLower(words[i])]; !ok2 {
			unique[strings.ToLower(words[i])] = Word{Word: words[i], Pos: i}
			continue
		}
		// Если слово содержится в карте уникальных слов, то нужно его очистить
		_, ok := unique[strings.ToLower(words[i])]
		if ok == true {
			words[i] = ""
		}
	}
	// Замена в слайсе слов при помощи карты уникальных слов и их индекса
	var res []string
	for _, value := range unique {
		res = append(words[:value.Pos], value.Word)
		res = append(res, words[value.Pos:]...)
	}
	// Возвращение предложения из слайса слов, используйте функцию WordsToSentence
	return WordsToSentence(words)
}

// WordsToSentence Удаляет пустые слова из слайса и объединяет их в предложение, добавляя в конце восклицательный знак
func WordsToSentence(words []string) string {
	filtered := make([]string, 0, len(words))

	for _, word := range words {
		if word != "" {
			filtered = append(filtered, word)
		}
	}

	return strings.ReplaceAll(strings.Join(filtered, " ")+"!", "!!", "!")
}

// CheckUpper Проверяет, нужно ли заменять первую букву на заглавную
func CheckUpper(old, new string) string {
	if len(old) == 0 || len(new) == 0 {
		return new
	}

	chars := []rune(old)

	if unicode.IsUpper(chars[0]) {
		runes := []rune(new)
		new = string(append([]rune{unicode.ToUpper(runes[0])}, runes[1:]...))
	}

	return new
}

// splitSentences Разделяет текст на предложения
func splitSentences(message string) []string {
	// Создание регулярного выражения для поиска знаков препинания
	originSentences := strings.Split(message, "!")
	var orphan string
	var sentences []string

	for i, sentence := range originSentences {
		words := strings.Split(sentence, " ")

		if len(words) == 1 {
			if len(orphan) > 0 {
				orphan += " "
			}

			orphan += words[0] + "!"
			continue
		}

		if orphan != "" {
			originSentences[i] = strings.Join([]string{orphan, " ", sentence}, " ") + "!"
			orphan = ""
		}

		sentences = append(sentences, originSentences[i])
	}

	return sentences
}

func main() {
	text := "Внимание! Внимание! Покупай срочно срочно крипту только у нас! Биткоин лайткоин эфир по низким ценам! Беги, беги, успевай стать финансово независимым с помощью крипты! Крипта будущее финансового мира!"
	censorMap := map[string]string{
		"крипта":   "фрукты",
		"крипту":   "фрукты",
		"крипты":   "фруктов",
		"биткоин":  "яблоки",
		"лайткоин": "яблоки",
		"эфир":     "яблоки",
	}

	filteredText := filterWords(text, censorMap)
	fmt.Println(filteredText) // Внимание! Покупай срочно фрукты только у нас! Яблоки по низким ценам! Беги, успевай стать финансово независимым с помощью фруктов! Фрукты будущее финансового мира!
}
