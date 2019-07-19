package main

import (
	"fmt"

	"github.com/varmamsp/cello/crawler"
	"github.com/varmamsp/cello/store/sqlstore"
)

func main() {
	c := crawler.NewCrawler(sqlstore.NewSqlStore())

	c.Start()
	fmt.Println("Mmmm")

	var cx chan int
	<-cx
}
