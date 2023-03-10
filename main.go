package main

import (
	"encoding/json"
	"strconv"
	"os"
	"log"
	"fmt"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {

	// .envファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return
	}

	// 環境変数からエンドポイントの値を取得する
	scheme := os.Getenv("SCHEME")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	engine:= gin.Default()

	engine.GET("/:n", func(c *gin.Context) {
		num_str := c.Param("n")
		num, err := strconv.Atoi(num_str)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number."})
			return
		}

		if (num < 1) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Negative number."})
			return
		}

		if (num == 1) {
			c.JSON(http.StatusBadRequest, gin.H{"n": 1, "error": nil})
			return
		}

    // HTTP GETリクエストを送信する
    resp, err := http.Get(fmt.Sprintf("%s://%s:%s/%d", scheme, host, port, num - 1))
    if err != nil {
      log.Fatal(err)
    }
    defer resp.Body.Close()

    // レスポンスJSONをデコードする
    var data map[string]interface{}
    err = json.NewDecoder(resp.Body).Decode(&data)
    if err != nil {
      log.Fatal(err)
    }

    // "n"プロパティの値を整数として取得する
    n, ok := data["n"].(float64)
    if ok == false {
      log.Fatal(`Property "n" is not a number.`)
    }

		// idを整数として使用する
		c.JSON(http.StatusOK, gin.H{"n": num * int(n), "error": nil})
	})
	engine.Run(fmt.Sprintf("0.0.0.0:%s", port))
}
