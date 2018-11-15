package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int
	UserName string `orm:"unique"`
	Pwd string
	Articles []*Article `orm:"rel(m2m)"`
}

type Article struct {
	Id int `orm:"pk;auto"`
	Title string `orm:"size(100)"`
	Content string `orm:"size(500)"`
	Time time.Time `orm:"type(datetime);auto_now"`
	ReadCount int  `orm:"default(0)"`
	Image string `orm:"null"`
	ArticleType *ArticleType `orm:"rel(fk);on_delete(set_null);null"`
	Users []*User `orm:"reverse(many)"`
}

type ArticleType struct {
	Id int
	TypeName string `orm:"size(100)"`
	Articles []*Article `orm:"reverse(many)"`
}

func init(){
	//生成表的三步骤
	//注册数据库
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/newsweb?charset=utf8")

	//注册表
	orm.RegisterModel(new(User),new(Article),new(ArticleType))

	//让项目跑起来
	orm.RunSyncdb("default",false,true)
}