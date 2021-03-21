package main

import (
    "github.com/gin-gonic/gin" //第三方库
    "go-module-example/pkg/file" //自定义本地库
)

func main() {
    r := gin.Default()
    r.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello world!",
        })
    })
    r.GET("/package", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": file.IsFile(),
        })
    })
    r.Run()
}
