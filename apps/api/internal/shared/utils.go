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
