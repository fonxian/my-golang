package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/parnurzeal/gorequest"
	"log"
)

//https://mholt.github.io/json-to-go/
//https://www.cnblogs.com/mafeng/p/8429983.html
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

//TODO
//翻页
//休眠

//入库
//https://github.com/go-sql-driver/mysql

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
	db()
}

//连接数据库
//https://blog.csdn.net/sinat_26682309/article/details/89738546
func db() {
	var id, name string
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/build")
	rows, err := db.Query("select company_id,credit_code from t_build_main limit 100")
	if err == nil {
		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err == nil {
				log.Printf("id:%s,name:%s", id, name)
			}

		}
	} else {
		log.Println(err)
	}

}
