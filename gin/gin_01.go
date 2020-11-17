package main

import "github.com/gin-gonic/gin"

func main() {
	users := []User{{
		ID:   123,
		Name: "张三",
	}, {
		ID:   456,
		Name: "李四",
	}}

	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "The user id is %s", id)
	})

	_ = r.Run(":8080")
}

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
