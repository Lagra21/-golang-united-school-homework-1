package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func ShowWorld() {
	msgStr := emoji.Sprint("Hello :map:!")
	fmt.Println(msgStr)
}

func main() {
	ShowWorld()
}
