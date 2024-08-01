package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	link := "https://old.reddit.com/r/memes/top/?sort=top&t=month"
	c := colly.NewCollector()

	c.OnHTML("p.title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit(link)
	c.Wait()

}
