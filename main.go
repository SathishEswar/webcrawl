package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"mvdan.cc/xurls/v2"
)

func main() {
	var result string
	a, _ := os.Create("slightly_smiling_face_emoji.txt")

	filerc, err := os.Open("inputs.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer filerc.Close()
	// create a new collector
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		result = result + e.Text + "\n"
		if strings.Contains(result, "😘") {
			mw := io.MultiWriter(os.Stdout, a)
			fmt.Fprintln(mw, *e.Request.URL)

		}

	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	contents := buf.String()
	rxStrict := xurls.Strict()
	k := rxStrict.FindAllString(contents, -1)

	Q, _ := queue.New(
		2, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000},
	)
	for _, url := range k {
		Q.AddURL(url)
		Q.Run(c)
	}

	log.Printf("Scraping done,\n")

}


hello im new line 67

lne 66
line67
line68


