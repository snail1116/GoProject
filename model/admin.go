package model

import "time"

//todo:定义管理员结构体
//如果field名称为Id，且为int64，且没有定义tag，则会被xorm视为主键，拥有自增属性
type Admin struct {
	AdminId    int64     `xorm:"pk autoincr" json:"id"`
	AdminName  string    `xorm:"varchar(32)" json:"admin_name"`
	CreateTime time.Time `xorm:"DateTiem" json:"create_time"`
	Status     int64     `xorm:"default 0" json:"status"`
	Avatar     string    `xorm:"varchar(255)" json:"avatar"`
	Pwd        string    `xorm:"varchar(255)" json:"pwd"`
	CityName   string    `xorm:"varchar(12)" json:"city_name"`
	CityId     int64     `xorm:"index" json:"city_id"`
	City       *City     `xorm:"- <- ->"`
}

func (this *Admin) AdminToRespDesc() interface{} {
	respDesc := map[string]interface{}{
		"user_name":   this.AdminName,
		"id":          this.AdminId,
		"create_time": this.CreateTime,
		"status":      this.Status,
		"avatar":      this.Avatar,
		"city":        this.CityName,
		"admin": "管理员",
	}
	return respDesc
}
