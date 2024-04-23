package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"unicode"
)

func GetAllBeefName() []string {
	var url = "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

	response, err := http.Get(url)
	if err != nil {
		return nil
	}
	if response.StatusCode != 200 {
		return nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var tokens = tokenize(string(body))
	return tokens
}

func tokenize(text string) []string {
	delimiterFunc := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '-'
	}

	tokens := strings.FieldsFunc(strings.ToLower(text), delimiterFunc)
	return tokens
}
