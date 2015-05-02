package main

import (
	"github.com/timakin/ts/loader"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandAll,
	commandBiz,
	commandHack,
}

var commandAll = cli.Command{
	Name:  "pop",
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

func doAll(c *cli.Context) {
	hn := make(chan loader.ResultData)
	ph := make(chan loader.ResultData)
	re := make(chan loader.ResultData)
	go loader.GetHNFeed(hn)
	go loader.GetPHFeed(ph)
	go loader.GetRedditFeed(re)
	hnres := <- hn
	phres := <- ph
	reres := <- re
	var HNData loader.Feed = &hnres
	var PHData loader.Feed = &phres
	var REData loader.Feed = &reres
	HNData.Display()
	PHData.Display()
	REData.Display()
}

func doBiz(c *cli.Context) {
}

func doHack(c *cli.Context) {
}
