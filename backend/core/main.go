package main

import (
	_ "github.com/lib/pq"
	"github.com/root9464/Ton-students/core/app"
)

func main() {
	a := app.NewApp()
	err := a.Run()
	if err != nil {
		panic(err)
	}

}
