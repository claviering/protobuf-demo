package main

import (
	"fmt"
	"io/ioutil"
	pb "main/protobuf/public/proto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

func main() {
	r := gin.Default()
	r.Static("/public", "./public")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/api/awesome", func(c *gin.Context) {
		req := &pb.AwesomeMessage{}
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			proto.UnmarshalMerge(body, req)
			fmt.Println(*&req.AwesomeField)
		}
		fmt.Print(req.AwesomeField)
		data := &pb.AwesomeMessage{
			AwesomeField: "hello world",
		}
		c.ProtoBuf(http.StatusOK, data)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "http://localhost:8080")
}
