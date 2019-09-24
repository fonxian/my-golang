package main

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
)

//https://mholt.github.io/json-to-go/
type DoubanMovie struct {
	Subjects []struct {
		Rate     string `json:"rate"`
		CoverX   int    `json:"cover_x"`
		Title    string `json:"title"`
		URL      string `json:"url"`
		Playable bool   `json:"playable"`
		Cover    string `json:"cover"`
		ID       string `json:"id"`
		CoverY   int    `json:"cover_y"`
		IsNew    bool   `json:"is_new"`
	} `json:"subjects"`
}

func main() {
	pageUrl := "https://movie.douban.com/j/search_subjects?type=movie&tag=%E8%B1%86%E7%93%A3%E9%AB%98%E5%88%86&sort=rank&page_limit=20&page_start=0"

	req := gorequest.New()

	resp, body, err := req.Get(pageUrl).
		Set("Host", "movie.douban.com").
		Set("Referer", "https://movie.douban.com/explore").
		End()

	fmt.Println(err)
	fmt.Println(resp)
	fmt.Println(body)

	var doubanMovie DoubanMovie
	json.Unmarshal([]byte(body), &doubanMovie)
	for _, movie := range doubanMovie.Subjects {
		log.Printf("name:%s,rate:%s", movie.Title, movie.Rate)
	}

}
