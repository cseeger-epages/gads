package main

import (
	"context"
	"fmt"
	"github.com/sakari-ai/gads"
	"log"
)

func main() {
	config, err := gads.NewCredentials(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	bs := gads.NewCampaignService(config.Auth)

	selector := gads.Selector{
		Fields: []string{
			"Id",
			"Name",
			"Status",
		},
		Paging: &gads.Paging{Offset: 0, Limit: 10},
	}
	adGroup, total, e := bs.Get(selector)
	fmt.Println(e)
	fmt.Println(adGroup, total)
}
