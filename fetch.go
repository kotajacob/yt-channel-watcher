package main

import (
	"context"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

func fetch(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(baseFeedURL+id, ctx)
	if err != nil {
		return err
	}

	log.Println(feed.Title)
	for _, item := range feed.Items {
		published, err := time.Parse(
			"2006-01-02T15:04:05Z07:00",
			item.Published,
		)
		if err != nil {
			log.Printf("failed parsing time for %v: %v\n", item.Title, err)
		}

		log.Println(published.Format("20060102 -"), item.Title, item.Link)
	}
	return nil
}
