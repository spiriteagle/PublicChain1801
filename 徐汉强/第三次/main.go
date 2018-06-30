package main

import (
	"fmt"

	"./lbc"
)

func main() {

	mybc := lbc.NewBlockchain()
	defer mybc.Db.Close()

	cli := lbc.CLI{mybc}
	cli.Run()

	fmt.Println("执行结束!")
}
