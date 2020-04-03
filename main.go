package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"iris_study/CmsProject/config"
	"iris_study/CmsProject/controller"
	"iris_study/CmsProject/datasource"
	"iris_study/CmsProject/service"
	"time"
)

func main() {
	app := newApp()
	configation(app)
	mvcHandle(app)
	config := config.InitConfig()
	addr := ":" + config.Port
	app.Run(
		iris.Addr(addr),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}

func mvcHandle(app *iris.Application) {
	
	//todo 启用session
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookie",
		Expires: 24 * time.Hour,
	})

	engine := datasource.NewMysqlEngine()

	//todo 管理员模块功能
	adminService := service.NewAdminService(engine)
	admin := mvc.New(app.Party("/admin"))
	admin.Register(
		adminService,
		sessManager.Start)
	admin.Handle(new(controller.AdminController))
}

func configation(app *iris.Application) {
	app.Configure(iris.WithConfiguration(iris.Configuration{Charset: "UTF-8"}))
	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg":    "not found",
			"data":   iris.Map{},
		})
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg":    "interal error",
			"data":   iris.Map{},
		})
	})
}

func newApp() *iris.Application {
	app := iris.New()

	//todo：设置日志级别，开发阶段为debug
	app.Logger().SetLevel("debug")

	//todo：注册静态资源
	app.HandleDir("/static", "./CmsProject/static")
	app.HandleDir("/manage/static", "./CmsProject/static")

	//todo:注册视图文件
	app.RegisterView(iris.HTML("./CmsProject/static", ".html"))
	app.Get("/", func(context context.Context) {
		context.View("index.html")
	})

	return app
}
