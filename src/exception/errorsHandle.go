package exception

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CheckErrors(err error, message string) {
	if err != nil {
		log.Println(message)
		log.Println(err)
	}
}

func ThrowExceptionBadArgument(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": http.StatusBadRequest,
		"error":  "Bad arguments",
	})
	log.Println(err)
}

func ThrowExceptionSQLError(c *gin.Context, err error, resp any) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status": http.StatusInternalServerError,
		"error":  "We could not execute your query",
		"data":   resp,
	})
	log.Println(err)
}
