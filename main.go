package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func add(context *gin.Context) {
	num1Str := context.Param("num1")
	num2Str := context.Param("num2")

	num1, err := strconv.Atoi(num1Str)
	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": "num1 must be an integer"})
		return
	}

	num2, err := strconv.Atoi(num2Str)
	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": "num2 must be an integer"})
		return
	}

	result := num1 + num2
	context.JSON(http.StatusOK, gin.H{"result": result})
}

func subtract(context *gin.Context) {
	num1Str := context.Param("num1")
	num2Str := context.Param("num2")

	num1, err := strconv.Atoi(num1Str)
	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": "num1 must be an integer"})
		return
	}

	num2, err := strconv.Atoi(num2Str)
	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": "num2 must be an integer"})
		return
	}

	result := num1 - num2
	context.JSON(http.StatusOK, gin.H{"result": result})
}

func main() {
	router := gin.Default()
	router.GET("/add/:num1/:num2", add)
	router.GET("/subtract/:num1/:num2", subtract)
	router.Run("localhost:3000")
}
