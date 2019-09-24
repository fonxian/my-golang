package main

// import "fmt"
// import "net/http"
// import "io/ioutil"

import (
	"fmt"
	// "github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	// "github.com/tidwall/gjson"
)

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

}
