package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type HelloReq struct {
	g.Meta `path:"/hello" method:"get"`
	Name   string `v:"required" dc:"姓名"`
	Age    int    `v:"required" dc:"年龄"`
}

type Hello struct{}

type HelloRes struct{}

func (Hello) Say(ctx context.Context, req *HelloReq) (res *HelloRes, err error) {
	r := g.RequestFromCtx(ctx)
	r.Response.Writef(
		"Hello %s! Your Age is %d",
		req.Name,
		req.Age,
	)
	return
}

// middleware 中间件处理函数
func ErrorHandler(r *ghttp.Request) {
	// use 路由回调函数
	// 前置中间件处理逻辑
	// 在请求到达路由之前执行的一些操作
	g.Log().Infof(r.Context(), "Incoming request: %s %s", r.Method, r.URL.Path)

	r.Middleware.Next() // 调用Next()，让请求继续传递到下一个中间件或最终的路由处理函数

	// 后置中间件处理逻辑
	// 在路由处理完后，可以对响应进行一些操作
	// 进入下一个流程
	// 判断是否产生 错误
	err := r.GetError()
	if err != nil {
		r.Response.Write("error occurs", err.Error())
		return
	}
}
func main() {
	s := g.Server()
	//s.BindHandler("/", func(r *ghttp.Request) {
	//	var req HelloReq
	//	if err := r.Parse(&req); err != nil {
	//		r.Response.Write(err.Error())
	//		return
	//	}
	//	if req.Name == "" {
	//		r.Response.Write("name is empty")
	//		return
	//	}
	//	if req.Age <= 0 {
	//		r.Response.Write("age > 0 pl")
	//		return
	//	}
	//	r.Response.Writef(
	//		"Hello %s! Your Age is %d",
	//		req.Name,
	//		req.Age)
	//})
	//s.SetPort(8010)
	//s.Run()
	s.Group("/", func(group *ghttp.RouterGroup) {
		//我们定义了一个错误处理的中间件ErrorHandler，
		//在该中间件中我们先通过r.Middleware.Next()执行路由函数流程，
		//随后通过r.GetError()获取路由回调函数是否有错误产生。
		//如果产生错误，那么直接输出错误信息。
		group.Middleware(ErrorHandler)
		group.Bind(
			new(Hello),
		)
	})

	// eg :
	// 定义路由组，并应用中间件
	//s.Group("/", func(group *ghttp.RouterGroup) {
	//	group.Middleware(LogMiddleware) // 记录日志的中间件  前置中间件
	//	group.Middleware(AuthMiddleware) // 权限验证的中间件
	//
	//	// 路由处理函数
	//	group.GET("/hello", HelloHandler)
	//})
	s.SetPort(8000)
	s.Run()
}
