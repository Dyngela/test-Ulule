package exception

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CheckErrors Utils function to log conveniently error message and root cause
func CheckErrors(err error, message string) {
	if err != nil {
		log.Println(message)
		log.Println(err)
	}
}

// ThrowExceptionBadArgument Utils function to log error and return API response in error in case given parameters
// from client request has error.
func ThrowExceptionBadArgument(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": http.StatusBadRequest,
		"error":  "Bad arguments",
	})
	log.Println(err)
}

// ThrowExceptionSQLError Utils function to log error and return API response in error in case SQL request failed.
// Return a generic message and an empty json
func ThrowExceptionSQLError(c *gin.Context, err error, resp any) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status": http.StatusInternalServerError,
		"error":  "We could not execute your query",
		"data":   resp,
	})
	log.Println(err)
}
