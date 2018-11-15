package routers

import (
	"newsWeb/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//添加路由过滤器函数   正则匹配路由  过滤器位置    过滤器函数
    beego.InsertFilter("/article/*",beego.BeforeExec,funcFilter)
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandleReg")
    beego.Router("/login",&controllers.UserController{},"get:ShowLogin;post:HandleLogin")
    beego.Router("/article/articleList",&controllers.ArticleController{},"get:ShowArticleList")
    beego.Router("/article/addArticle",&controllers.ArticleController{},"get:ShowAddArticle;post:HandeAddArticle")
    beego.Router("/article/articleDetail",&controllers.ArticleController{},"get:ShowArticleDetail")
    beego.Router("/article/updateArticle",&controllers.ArticleController{},"get:ShowUpdateArticle;post:HandleUpdateArticle")
    beego.Router("/article/deleteArticle",&controllers.ArticleController{},"get:DeleteArticle")
    beego.Router("/article/addType",&controllers.ArticleController{},"get:ShowAddType;post:HandleAddType")
    beego.Router("/article/deleteType",&controllers.ArticleController{},"get:DeleteType")
    beego.Router("/article/logout",&controllers.UserController{},"get:Logout")
}

var funcFilter = func(ctx*context.Context) {
	//登录校验
	userName := ctx.Input.Session("userName")
	if userName == nil{
		ctx.Redirect(302,"/login")
	}
}
