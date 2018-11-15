package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"newsWeb/models"
	"encoding/base64"
)

type UserController struct {
	beego.Controller
}

//展示注册页面
func (this*UserController)ShowRegister(){
	this.TplName = "register.html"
}

//处理注册业务
func(this*UserController)HandleReg(){
	//接受数据
	userName := this.GetString("userName")
	pwd := this.GetString("password")
	//校验数据
	if userName == "" || pwd ==""{
		this.Data["errmsg"] = "用户名或密码不能为空"
		this.TplName = "register.html"
		return
	}

	//处理数据
	//插入数据
	o := orm.NewOrm()
	//插入对象
	var user models.User
	//给插入对象赋值
	user.UserName = userName
	user.Pwd = pwd
	//插入
	_,err := o.Insert(&user)
	if err != nil{
		this.Data["errmsg"] = "注册失败，请重新注册！"
		this.TplName = "register.html"
		return
	}

	//返回数据
	//this.Ctx.WriteString("注册成功")
	this.Redirect("/login",302)
	//this.TplName = "login.html"
}

//展示登录页面
func(this*UserController)ShowLogin(){
	dec := this.Ctx.GetCookie("userName")
	userName,_:=base64.StdEncoding.DecodeString(dec)
	if string(userName) != ""{
		this.Data["userName"] = string(userName)
		this.Data["checked"] = "checked"
	}else{
		this.Data["userName"] = ""
		this.Data["checked"] = ""
	}
	this.TplName = "login.html"
}

//处理登录业务
func(this*UserController)HandleLogin(){
	//接受数据
	userName := this.GetString("userName")
	pwd := this.GetString("password")
	//校验数据
	if userName == "" || pwd ==""{
		this.Data["errmsg"] = "用户名或者密码不能为空"
		this.TplName = "login.html"
		return
	}

	//处理数据
	//查询业务
	//获取orm对象
	o := orm.NewOrm()
	//获取查询对象
	var user models.User
	//给查询条件赋值
	user.UserName = userName
	//查询
	err := o.Read(&user,"UserName")
	if err != nil{
		this.Data["errmsg"] = "用户名不存在"
		this.TplName = "login.html"
		return
	}
	//判断密码是否正确
	if user.Pwd != pwd{
		this.Data["errmsg"] = "密码错误，请重新输入！"
		this.TplName = "login.html"
		return
	}

	//获取是否记住用户名
	remember := this.GetString("remember")
	if remember == "on"{
		enc := base64.StdEncoding.EncodeToString([]byte(userName))
		this.Ctx.SetCookie("userName",enc,3600 * 1)
	}else {
		this.Ctx.SetCookie("userName",userName,-1)
	}

	//返回数据
	//this.Ctx.WriteString("登录成功！")
	this.SetSession("userName",userName)
	this.Redirect("/article/articleList",302)
}

//退出登录
func(this*UserController)Logout(){
	//删除session
	this.DelSession("userName")
	//返回页面
	this.Redirect("/login",302)
}