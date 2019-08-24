package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/**
golang处理请求的一个流程：
listen -> RegisterHandlers -> handlers、
golang的router进行配置一个路由，接收请求，并且分发路由
由于go的微量的虚拟协程技术，使得开启一个是非常的方便，大小只有4k
大小，所以效率较高，并发很厉害，利用计算机的多核资源
*/
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}
func main() {
	//使用闭包将请求的Handler进行一个封装
	r := RegisterHandlers()
	//注册服务器，并设置路由及其监听端口
	http.ListenAndServe("127.0.0.1:8000", r)
}
