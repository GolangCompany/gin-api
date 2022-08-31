package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type News struct {
	Date int    `json:"date" binding: "required"`
	Name string `json:"name" binding: "required"`
}

func main() {
	w := gin.Default()
	w.GET("", func(s *gin.Context) {
		s.JSON(http.StatusOK, gin.H{
			"responseData": "Misha to sell exquisite gear",
		})
		//s.String(http.StatusOK, "Misha to sell exquisite gear")
	})
	w.GET("/:newsupdate", func(s *gin.Context) {
		var val = s.Param("newsupdate")
		s.JSON(http.StatusOK, gin.H{
			"newsupdate": val,
		})
	})
	w.POST("/add", func(s *gin.Context) {
		var data News
		if err := s.ShouldBind(&data); err != nil {
			fmt.Println(err)
			s.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
		} else {
			s.JSON(http.StatusOK, gin.H{
				"data": data,
			})
		}
	})

	w.Run("localhost:8080")
}
