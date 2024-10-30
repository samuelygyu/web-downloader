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

		var args []string
		args = append(args, req.Url, "-o", "/Temp/you-get/download")
		if req.Proxy!= "" { 
			args = append(args, "-x", req.Proxy)
		}
		if cookiesName!= "" {
			args = append(args, "-c", cookiesName)
		}

		msg := exec.Command("you-get", args...)
		log.Printf("cmd: %v", msg.Args)
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
