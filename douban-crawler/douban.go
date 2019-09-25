package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/parnurzeal/gorequest"
	"log"
	"strconv"
	"time"
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

//数据库连接池
//分模块

func main() {

	var db, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/build?charset=utf8")
	for i := 0; i < 200; i++ {
		crawler(db, 20*i)
		//休眠2秒
		time.Sleep(time.Duration(2) * time.Second)
	}

	db.Close()

	// dbConnect(db)
}

func crawler(db *sql.DB, pageStart int) {

	pageUrl := "https://movie.douban.com/j/search_subjects?type=movie&tag=%E8%B1%86%E7%93%A3%E9%AB%98%E5%88%86&sort=rank&page_limit=20&page_start=" + strconv.Itoa(pageStart)

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
		insert(db, movie.Title, movie.Rate)
	}
}

//连接数据库
//https://blog.csdn.net/sinat_26682309/article/details/89738546
// func dbConnect(db *sql.DB) {

// 	var id, name string
// 	rows, err := db.Query("select company_id,credit_code from t_build_main limit 100")
// 	if err == nil {
// 		for rows.Next() {
// 			err := rows.Scan(&id, &name)
// 			if err == nil {
// 				log.Printf("id:%s,name:%s", id, name)
// 			}

// 		}
// 	} else {
// 		log.Println(err)
// 	}

// }

//入库
//https://github.com/go-sql-driver/mysql
func insert(db *sql.DB, title string, rate string) {
	stmt, _ := db.Prepare("INSERT t_douban SET title=?,rate=?")
	stmt.Exec(title, rate)

}
