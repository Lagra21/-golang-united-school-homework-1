package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	msgStr := emoji.Sprint("Hello, world! :balloon:")
	fmt.Println(msgStr)
}
