package main

import (
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
	endpoint := os.Getenv("ENDPOINT")

	engine:= gin.Default()

	engine.GET("/:num", func(c *gin.Context) {
		num_str := c.Param("num")
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

		// idを整数として使用する
		c.JSON(http.StatusOK, gin.H{"n": n})
	})
	engine.Run("0.0.0.0:8080")
}
