package main

import (
	"log"
	"math/rand"
)

var APP_NAME string
var APP_VERSION string

func init() {
	if APP_NAME == "" {
		APP_NAME = "random generator"
	}
	if APP_VERSION == "" {
		APP_VERSION = "dev-release"
	}
}

func main() {
	var allowedCharsASCIIStart = 33
	var allowedCharsASCIIEnd = 122
	var randomLength = 16
	generatedString := ""
	for idx := 0; idx < randomLength; idx++ {
		randomASCII := rand.Intn(allowedCharsASCIIEnd - allowedCharsASCIIStart) + allowedCharsASCIIStart
		generatedString += string(rune(randomASCII))
	}
	log.Printf("Generated: %v\n", generatedString)
}
