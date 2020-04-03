package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"iris_study/CmsProject/model"
)

//todo 数据库引擎

//todo 实例化数据库引擎方法

func NewMysqlEngine() *xorm.Engine {
	//todo 注意，一定记得在import中引入 mysql依赖库
	engine, err := xorm.NewEngine("mysql", "root:123456@/ginger?charset=utf8")
	//todo Sync2进行如下操作
	/**
	 * 自动检测和创建表，这个检测是根据表的名字
	 * 自动检测和新增表中的字段，这个检测是根据字段名，同时对表中多余的字段给出警告信息
	 * 自动检测，创建和删除索引和唯一索引，这个检测是根据索引的一个或多个字段名，而不根据索引名称。因此这里需要注意，如果在一个有大量数据的表中引入新的索引，数据库可能需要一定的时间来建立索引。
	 * 自动转换varchar字段类型到text字段类型，自动警告其它字段类型在模型和数据库之间不一致的情况。
	 * 自动警告字段的默认值，是否为空信息在模型和数据库之间不匹配的情况
	 */
	err = engine.Sync2(
		new(model.City),
		new(model.Permission),
		new(model.Admin),
	//new(model.AdminPermission),
	//new(model.User),
	//new(model.UserOrder),
	)
	if err != nil {
		panic(err.Error())
	}
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)
	return engine
}
