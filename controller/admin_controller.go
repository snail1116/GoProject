package controller

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"iris_study/CmsProject/service"
)

//todo 管理员控制器

type AdminController struct {
	//todo:iris框架自动为每个请求绑定上下文对象
	Ctx iris.Context

	//todo:admin功能实体
	Service service.AdminService

	//todo:session对象
	Session *sessions.Session
}

const (
	ADMINTABLENAME = "admin"
	ADMIN          = "admin"
)

type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//todo 管理员登录功能  admin/login
func (ac *AdminController) PostLogin(context iris.Context) mvc.Result {
	iris.New().Logger().Info("admin login ")

	var adminLogin AdminLogin
	ac.Ctx.ReadJSON(&adminLogin)

	if adminLogin.UserName == "" || adminLogin.Password == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "0",
				"success": "登录失败",
				"message": "用户名或密码为空，请填写后尝试登录",
			},
		}
	}

	//todo:根据用户名密码到数据库查询
	admin, exist := ac.Service.GetByAdminNameAndPassword(adminLogin.UserName, adminLogin.Password)

	if !exist {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "0",
				"success": "登录失败",
				"message": "用户名或密码错误，请重新登录",
			},
		}
	}

	//todo:管理员存在，设置session
	userByte, _ := json.Marshal(admin)
	ac.Session.Set(ADMIN, userByte)
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  "1",
			"success": "登录成功",
			"message": "管理员登录成功",
		},
	}
}
