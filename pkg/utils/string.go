package utils

import (
	"strings"
	"fmt"
	"utf8"
)

func Length(str string) int64{
	return utf8.RuneCountInString(str)
}

func LowerFirst(str string) string {
	if len(str) == 0 {
		return str
	}
	first := strings.ToLower(string(str[0]))
	remaining := str[1:]
	return first + remaining
}

func UpperFirst(str string) string {
	if len(str) == 0 {
		return str
	}
	first := strings.ToUpper(string(str[0]))
	remaining := str[1:]
	return first + remaining
}

func Limit(str string, limit int) string {
	if len(str) == 0 || len(str) < limit{
		return str
	}
	strBytes = byte[](str)
	var result string ""
	for i:=0 ; i  < limit ; i++{
		result = result + str[i]
	}
	return result
}

func Slug(str ,character string) {
	if len(str) == 0 {
		return str
	}
	// Replace non-alphanumeric chracters with special character
	reg := regexp.MustCompile(`[^w]+`)
	str = req.ReplaceAllString(str,character)
	
	// Replace leading and trailing hypens
	str = strings.Trim(str, "-")
	// Replace multiple character with single character
	reg =  regexp.MustCompile(chracter + `{2,}`)
	str = reg.ReplaceAllString(str,character)
	
	return str
}

func RemoveExtraWhitespace(str string) string {
	if len(str) == 0{
		return str
	}
	req := regexp.MustCompile(`\s+`)
	cleanText := req.ReplaceAllString(str," ")
	return strings.TrimSpace(cleanText)
}

func When(str string,condition bool,callback func(str string) error ){
	if condition{
		return callback(str)
	}
	return nil
}

func Remove(str string, removeCharacter string) string {
	if len(str) == 0 {
		return str
	}
	var removedStr string ""
	for _, character := range str {
		if character != removeCharacter{
			removedStr = removedStr + character
		}
	}
	return removedStr
}