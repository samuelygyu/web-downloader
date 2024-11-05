package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"

	"samuel/you-get-web/tools"
)

type DownloadReq struct {
	Url   string
	Proxy string
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},  // 允许的来源，可以是多个
		AllowMethods:     []string{"GET", "POST", "OPTIONS"}, // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Content-Type"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},         // 允许浏览器访问的响应头
		AllowCredentials: true,                               // 是否允许发送 Cookie
		MaxAge:           12 * 3600,                          // 预检请求的缓存时间
	}))

	r.POST("/download", func(c *gin.Context) {
		iflag := randstr.String(16)
		var req DownloadReq
		req.Url = c.PostForm("url")
		req.Proxy = c.PostForm("proxy")
		if req.Url == "" {
			c.JSON(http.StatusNoContent, gin.H{"error": "url is empty"})
			return
		}

		file, err1 := c.FormFile("cookies")
		var cookiesName string
		if err1 == nil {
			log.Println(file.Filename)
			cookiesName = fmt.Sprintf("/Temp/cookies/%s.txt", iflag)
			c.SaveUploadedFile(file, cookiesName)
		}

		var args []string
		args = append(args, req.Url, "-o", "/Temp/you-get/download")
		args = append(args, "-O", iflag)
		if req.Proxy != "" {
			args = append(args, "-x", req.Proxy)
		}
		if cookiesName != "" {
			args = append(args, "-c", cookiesName)
		}

		msg := exec.Command("you-get", args...)
		log.Printf("cmd: %v", msg.Args)
		output, err2 := msg.Output()

		if err2 != nil {
			fmt.Println("Error executing command:", err2)
			c.JSON(http.StatusExpectationFailed, gin.H{"error": "fail, please try again"})
			return
		}

		suffix, title, err3 := tools.SplitFileSuffix(string(output))
		if err3 != nil {
			fmt.Println("Error cannot find the file suffix", err3)
			c.JSON(http.StatusExpectationFailed, gin.H{"error": "fail, please try again"})
			return
		}

		file1 := fmt.Sprintf("%s[00].%s", iflag, suffix)
		file2 := fmt.Sprintf("%s[01].%s", iflag, suffix)

		if _, err3 := os.Stat(file1); os.IsNotExist(err3) {
			fmt.Printf("file1 %s not exist", file1)
		} else {
			c.FileAttachment("/Temp/you-get/download/" + file1, title)
			return
		}
		if _, err3 := os.Stat(file2); os.IsNotExist(err3) {
			fmt.Printf("file2 %s not exist", file2)
		} else {
			c.File("/Temp/you-get/download/" + file2)
		}
		c.JSON(http.StatusOK, gin.H{"status": req.Url + iflag})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
