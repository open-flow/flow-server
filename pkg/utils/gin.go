package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
)

func BindGinFactory(e *gin.Engine) func(httpMethod string, relativePath string, fn interface{}) {
	return func(httpMethod string, relativePath string, fn interface{}) {
		reqType := reflect.TypeOf(fn).In(1).Elem()
		fnValue := reflect.ValueOf(fn)

		e.Handle(httpMethod, relativePath, func(c *gin.Context) {
			respond := func(val interface{}, err interface{}) {
				if err == nil {
					c.JSON(http.StatusOK, val)
				} else {
					c.AbortWithStatus(http.StatusInternalServerError)
					log.Println("function returned error", err)
				}
			}

			reqValue := reflect.New(reqType)
			err := c.ShouldBind(reqValue.Interface())
			if err != nil {
				_ = c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			args := []reflect.Value{reflect.ValueOf(c), reqValue}
			resArr := fnValue.Call(args)
			resVal := resArr[0].Interface()
			errVal := resArr[1].Interface()

			respond(resVal, errVal)
		})
	}
}
