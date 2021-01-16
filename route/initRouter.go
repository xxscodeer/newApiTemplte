package route

import (
	"XxxxMicroAPI/middleware"
	"XxxxMicroAPI/tools"

	"github.com/kataras/iris/v12"
	"github.com/micro/go-micro/v2"
)

func InitRouter(service micro.Service) *iris.Application {
	app := iris.Default()
	app.Use(middleware.Cors)
	//初始化对象

	//创建服务端调用接口
	tokenRemoteName := tools.GetConfig().UserMicroName.Name

	//token验证
	//编写api接口

	return app
}
