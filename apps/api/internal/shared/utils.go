package shared

import (
	"strconv"

	"github.com/rs/zerolog/log"
)

func ParseStringArrayIntoIntArray(strArr []string) []int {
	var intArr []int
	for _, code := range strArr {
		intCode, err := strconv.Atoi(code)
		if err == nil {
			intArr = append(intArr, intCode)
		} else {
			log.Warn().Err(err).Msgf("Failed to parse int from string '%s', skipping", code)
		}
	}
	return intArr
}

func ParseStringIntoInt(str string) int {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		log.Warn().Err(err).Msgf("Failed to parse int from string '%s', using 0", str)
		return 0
	}
	return intValue
}

func CleanStringArray(strArr ...*[]string) {
	for _, arr := range strArr {
		*arr = []string{}
	}
}

func IntAsStrings(intArr []int) []string {
	var strArr []string
	for _, num := range intArr {
		strArr = append(strArr, strconv.Itoa(num))
	}
	return strArr
}
