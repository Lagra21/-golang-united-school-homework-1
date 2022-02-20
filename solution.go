package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func GetWorldTxt() string {
	return emoji.Sprint(":map:")
}

func ShowHelloWorld() {
	worldTxt := GetWorldTxt()
	fmt.Println("Hello " + worldTxt + "!")
}

func main() {
	ShowHelloWorld()
}
