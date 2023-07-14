package main

import (
	"api-git-clone/app"
)

func main() {
	// setup and run app
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
