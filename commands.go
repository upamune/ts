package main

import (
	"log"
	"os"
	"github.com/timakin/ts/loader"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandAll,
	commandBiz,
	commandHack,
}

var commandAll = cli.Command{
	Name:  "all",
	Usage: "",
	Description: `
`,
	Action: doAll,
}

var commandBiz = cli.Command{
	Name:  "biz",
	Usage: "",
	Description: `
`,
	Action: doBiz,
}

var commandHack = cli.Command{
	Name:  "hack",
	Usage: "",
	Description: `
`,
	Action: doHack,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doAll(c *cli.Context) {
	hn := make(chan loader.ResultData)
	go loader.GetHNFeed(hn)
	phres := <- hn
	var HNData loader.Feed = &phres
	HNData.Display()
	loader.GetPHFeed()
}

func doBiz(c *cli.Context) {
}

func doHack(c *cli.Context) {
}
