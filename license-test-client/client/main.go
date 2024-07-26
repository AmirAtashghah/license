package main

import (
	"client/apicaller"
	"client/mockrepo"
	"fmt"
	"github.com/robfig/cron"
	"log"
)

//var SudoPassword string

func main() {

	//if len(os.Args) != 2 {
	//	fmt.Println("Usage: go run main.go <sudo_password>")
	//	return
	//}
	//
	//SudoPassword = os.Args[1]

	mRepo := mockrepo.NewMockRepo()

	checker := apicaller.NewService(mRepo)

	c := cron.New()
	c.AddFunc("*/10 * * * *",
		func() {
			if err := checker.Check(); err != nil {
				log.Println(err)
			} else {
				fmt.Println("checked license, every things ok")
			}

		})

	// Start cron with one scheduled job
	log.Println("Start cron")
	c.Start()

	select {}
}
