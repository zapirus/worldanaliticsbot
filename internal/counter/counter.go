package counter

import (
	"Projects/WordAnalytics/pkg/logger"
	"log"
	"regexp"
	"strings"
	"unicode"
)

type Word struct {
	WordID int
	Word   string
	Count  int
}

func Count(parsedUrl string) []Word {
	logging := logger.GetLogger()

	logging.Info("Formatting text")
	parsedUrl = formatText(parsedUrl)
	strArr := strings.Split(parsedUrl, " ")
	var strMap = map[string]int{}

	logging.Info("Checking result")
	for _, value := range strArr {
		if isWord(value) {
			if _, isExist := strMap[value]; isExist {
				strMap[value] += 1
			} else {
				strMap[value] = 1
			}
		}

	}

	logging.Info("Result out")
	return toObjectArr(strMap)
}

func toObjectArr(mapForWords map[string]int) []Word {
	var arr []Word
	for value, key := range mapForWords {
		if value == "" {
			continue
		} else {
			arr = append(
				arr,
				Word{
					Word:  value,
					Count: key,
				})
		}
	}

	return arr
}

func isWord(wordFromUrl string) bool {
	for _, value := range wordFromUrl {
		if unicode.IsDigit(rune(value)) {
			return false
		}
	}
	return true
}

func formatText(parsedUrl string) string {
	parsedUrl = strings.ReplaceAll(parsedUrl, "/", "")
	parsedUrl = strings.ReplaceAll(parsedUrl, ",", " ")
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	parsedUrl = re.ReplaceAllString(parsedUrl, " ")

	return parsedUrl
}
