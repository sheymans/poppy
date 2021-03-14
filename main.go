package main

import (
	"fmt"
	"github.com/sheymans/poppy/app"
	"os"
)

func main() {
	err := app.Run()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
