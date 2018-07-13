package main

import (
	"net/http"
	"xinyun/handler"

	"github.com/gin-gonic/gin"
)

/* type postRequst struct {
	accessToken     string
	gaugeID         string
	subjectsAnswers string
}
*/

func main() {
	router := gin.Default()
	router.POST("/first", handler.PostHandler)

	http.ListenAndServe(":8000", router)
}
