package lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jtthink/framework/gin_"
	"jtthink/src/Boot"
)

//放 通用中间件

//检查服务是否可用
func Check_Middleware()  gin_.Middleware {
	return func(next gin_.Endpoint) gin_.Endpoint {
		return func(context *gin.Context, request interface{}) (response interface{}, err error) {
			if  !Boot.ServerIsReady(){
				return nil,fmt.Errorf("server is loading...")
			}
			return next(context,request)
		}
	}
}