package main

import (
	controllers "gitee.com/ethancheng/regular_demo/internal/pkg/controler"
	"github.com/gin-gonic/gin"
	"time"
)

// Person 注意这里的struct里面的字段不要小写，否则无法更改
// 另外注意这里的date_format "2006-01-02 15:04:05"
type Person struct {
	Name     string    `form:"name" binding:"required"`
	Age      int       `form:"age" binding:"required,gt=11"`
	Birthday time.Time `form:"birthday" date_format:"2006-01-02"`
}

func bind(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.String(200, "%s", err)
		return
	} else {
		c.String(200, "values is %v", person)
	}

	//if error :=c.ShouldBind(&person) ;error!=nil{
	//	c.String(200,"%v",error)
	//} else {
	//	c.String(200,"values is %v",person)
	//}
}

func getvideo(c *gin.Context) {
	c.String(200, "%s", "get")
}

func checkClientIp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{"127.0.0.2"}
		flag := false
		for _, i := range ipList {
			if i == c.ClientIP() {
				flag = true
				break
			}
			if !flag {
				c.String(401, "%s in not in ipList", c.ClientIP())
				c.Abort()
			}
		}
	}
}

func main() {
	r := gin.Default()
	//r := gin.New()
	//r.Use(checkClientIp())
	//r.GET("/ping", func(c *gin.Context) {
	//   c.JSON(200,gin.H{
	//   	"message":"pong",
	//   })
	//})
	////get方法
	//r.GET("/get", func(c *gin.Context) {
	//	c.String(200,"get")
	//})
	////POST方法
	//r.POST("/post", func(c *gin.Context) {
	//	c.String(200,"post")
	//})
	//
	////注意这里httpMethod是大写DELETE
	//r.Handle("DELETE","/delete", func(c *gin.Context) {
	//	c.String(200,"delete")
	//})
	//
	//r.Any("/any", func(c *gin.Context) {
	//	c.String(200,"any")
	//})
	//
	//r.Static("/assets","/assets")
	//
	//// URL参数
	//r.GET("/:name/:id", func(c *gin.Context) {
	//	c.JSON(200,gin.H{
	//		"name":c.Param("name"),
	//		"id": c.Param("id"),
	//	})
	//})
	//
	////泛绑定
	//r.GET("/user/*id", func(c *gin.Context) {
	//	c.String(200,"泛绑定")
	//})
	//
	////获取get参数
	//r.GET("/getParam", func(c *gin.Context) {
	//	c.String(http.StatusOK,"%s,%s",c.Query("name"),c.DefaultQuery("id","0"))
	//})
	//
	////获取body参数
	//r.POST("/getBody", func(c *gin.Context) {
	//	body,error := io.ReadAll(c.Request.Body)
	//	if error!= nil{
	//		c.String(http.StatusBadRequest,error.Error())
	//		c.Abort()
	//	}
	//	c.String(200,"%s",body)
	//})

	control := controllers.NewController()

	group := r.Group("/videos")
	group.GET("/", control.GetAll)
	group.POST("/:id", control.Update)
	group.POST("/", control.Create)

	//

	//bind参数

	r.GET("/middleware", func(c *gin.Context) {
		c.String(200, "%s", c.ClientIP())
	})
	r.POST("/bind", bind)
	r.GET("/bind", bind)

	// 启动gin服务，默认情况下是8080端口，当然也可以自定义端口号
	r.Run("127.0.0.1:8100")
	// 测试git pull远程操作，继续编辑
}
