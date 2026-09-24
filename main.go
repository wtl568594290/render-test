package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	ImgDir = "images"
)

var ImgUuidMap = make(map[string]string)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.File("./index.html")
	})
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	router.POST("/doc", func(c *gin.Context) {
		var req []FileInfoShortT
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		DocList = req

		// 清空 ImgUuidMap,图片目录
		ImgUuidMap = make(map[string]string)
		// 清空图片目录
		if err := os.RemoveAll(ImgDir); err != nil {
			log.Printf("remove img dir %s failed: %v", ImgDir, err)
		}

		c.JSON(200, gin.H{
			"message": "success",
		})

	})
	router.GET("/doc", func(c *gin.Context) {
		if len(DocList) == 0 {
			c.JSON(400, gin.H{
				"message": "doc list is empty",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "success",
			"doc":     DocList,
		})
	})
	router.POST("/img", func(c *gin.Context) {
		name := c.PostForm("name")
		if name == "" {
			c.JSON(400, gin.H{
				"message": "name is empty",
			})
			return
		}
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		ext := filepath.Ext(file.Filename)

		filename := uuid.NewString() + ext
		path := filepath.Join(ImgDir, filename)

		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		ImgUuidMap[name] = filename
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"img":     filename,
		})
	})
	router.GET("/img", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			c.JSON(400, gin.H{
				"message": "name is empty",
			})
			return
		}
		img, ok := ImgUuidMap[name]
		if !ok {
			c.JSON(400, gin.H{
				"message": "img not found",
			})
			return
		}
		c.File(filepath.Join(ImgDir, img))
	})
	router.Run() // 默认监听 0.0.0.0:8080
}
