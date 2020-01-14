package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("Time", api.Time)
	r.Run() // localhost: 8080
}

func ConnectToDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=유저네임 password=비밀번호
	dbname=디비명 host=IP주소 sslmode=disable")
	if db != nil {
		db.SetMaxOpenConns(100)  // 최대 커넥션 수 제한
		db.SetMaxIdleConns(10)	 // 대기 커넥션 최대 개수를 설정
	}
	if err != nil {
		return nil, err
	}
	return db, err
}

func Time(c *gin.Context) {
	db, err := model.ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Time 함수가 종료되기 직전에 실행

	var time stringgin
	err = db.QueryRow("" +
		"SELECT now()").Scan(&time)	// Scan()으로 time변수에 값을 넣어줌(&)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, map[string]string{
			"Time": time,
	})
}