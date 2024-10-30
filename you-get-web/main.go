package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
)

type DownloadReq struct {
	Url   string
	Proxy string
}

func generateCommand(url string, proxy string, cookiesName string) string {
	fomat := "you-get %s%s%s -o /Temp/you-get/download"
	var p string
	var c string
	if proxy != "" {
		p = fmt.Sprintf(" -x %s", proxy)
	}
	if cookiesName != "" {
		c = fmt.Sprintf(" -c %s", cookiesName)
	}

	return fmt.Sprintf(fomat, url, p, c)
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.POST("/download", func(c *gin.Context) {
		iflag := randstr.String(16)
		var req DownloadReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		file, err1 := c.FormFile("cookies")
		var cookiesName string
		if err1 == nil {
			log.Println(file.Filename)
			cookiesName = fmt.Sprintf("/Temp/cookies/%s.txt", iflag)
			c.SaveUploadedFile(file, cookiesName)
		}

		msg := exec.Command("you-get", generateCommand(req.Url, req.Proxy, cookiesName))
		_, err2 := msg.Output()
		if err2 != nil {
			fmt.Println("Error executing command:", err2)
		}

		c.JSON(http.StatusOK, gin.H{"status": req.Url + iflag})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
