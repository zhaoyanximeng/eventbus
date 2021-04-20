package main

import (
	"eventbus/eventbus"
	"github.com/gin-gonic/gin"
	"time"
)

var eBus = eventbus.NewEventBus()

func getPrice() int {
	return 3000
}

func getProdInfo() interface{} {
	ch := eBus.Sub("price")
	eBus.Pub("price",getPrice())
	prod := gin.H{"id":10,"name":"golang实践","price":ch.Data(time.Second)}

	return prod
}

func GetList() interface{} {
	time.Sleep(time.Second * 5)
	return gin.H{
		"message": "商品列表",
	}
}

func main() {
	r := gin.New()
	r.GET("/prods", func(c *gin.Context) {
		ch := eBus.Sub("prods") // 订阅

		go func() {
			eBus.Pub("prods",getProdInfo()) // 发布
		}()

		c.JSON(200,ch.Data(time.Second * 1))
	})

	r.Run(":8080")
}
