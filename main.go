package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
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
		addList := StorageMap.sync(req)
		c.JSON(200, gin.H{
			"message": "success",
			"data":    addList,
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
			"data":    DocList,
		})
	})

	router.POST("/img", func(c *gin.Context) {
		name := c.PostForm("name")
		hash := c.PostForm("hash")
		ext := c.PostForm("ext")
		if name == "" || hash == "" || ext == "" {
			c.JSON(400, gin.H{
				"message": "some params are empty",
			})
			return
		}
		if StorageMap.has(hash + name) {
			c.JSON(200, gin.H{
				"message": "img already exists",
			})
			return
		}

		filename := hash + ext
		path := filepath.Join(ImgDir, filename)
		// 如果文件存在，跳过c.SaveUploadedFile
		if _, err := os.Stat(path); os.IsNotExist(err) {
			// 文件不存在，继续执行
			file, err := c.FormFile("file")
			if err != nil {
				c.JSON(400, gin.H{
					"message": err.Error(),
				})
				return
			}

			if err := c.SaveUploadedFile(file, path); err != nil {
				c.JSON(400, gin.H{
					"message": err.Error(),
				})
				return
			}
		}
		StorageMap.add(hash+name, path)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    filename,
		})
	})

	router.GET("/img", func(c *gin.Context) {
		name := c.Query("name")
		hash := c.Query("hash")
		if name == "" || hash == "" {
			c.JSON(400, gin.H{
				"message": "name or hash is empty",
			})
			return
		}
		path, err := StorageMap.get(hash + name)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.File(path)
	})
	router.Run() // 默认监听 0.0.0.0:8080
}
