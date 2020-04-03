package service

import (
	"github.com/go-xorm/xorm"
	"iris_study/CmsProject/model"
)

//todo 管理员服务，标准的开发模式将每个实体提供的功能以接口标准的形式定义，供控制层进行调用。

type AdminService interface {
	//todo 通过管理员名+密码获取管理员实体，如查询到，返回管理员实体，并返回true，否则返回nil，false
	GetByAdminNameAndPassword(username, password string) (model.Admin, bool)

	//获取管理员总数
	GetAdminCount() (int64, error)
}

type adminService struct {
	engine *xorm.Engine
}

func NewAdminService(db *xorm.Engine) AdminService {
	return &adminService{
		engine: db,
	}
}

func (ac *adminService) GetAdminCount() (int64, error) {
	count, err := ac.engine.Count(new(model.Admin))

	if err != nil {
		panic(err.Error())
		return 0, err
	}
	return count, nil
}

func (ac adminService) GetByAdminNameAndPassword(username, password string) (model.Admin, bool) {
	var admin model.Admin
	ac.engine.Where("admin_name = ? and pwd = ? ", username, password).Get(&admin)
	return admin, admin.AdminId != 0
}
