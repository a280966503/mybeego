package controllers

import (
	"github.com/astaxie/beego"
)

type CartController struct {
	beego.Controller

}

func (this *CartController) HandleAddCart()  {
	////获取数据
	//skuid,err1 := this.GetInt("skuid")
	//count,err2 := this.GetInt("count")
	//resp := make(map[string]interface{})
	//defer this.ServeJSON()
	//
	////校验数据
	//if err1 != nil || err2 != nil {
	//	resp["code"]=1
	//	resp["msg"]="传递的数据不正确"
	//	this.Data["json"] = resp
	//	return
	//}
	//
	//userName := this.GetSession("userName")
	//if userName == nil {
	//	resp["code"]=2
	//	resp["msg"]="当前用户未登陆"
	//	this.Data["json"] = resp
	//	return
	//}
	//o := orm.NewOrm()
	//var user models.User
	//user.Name = userName.(string)
	//o.Read(&user,"Name")
	//
	////处理数据
	////购物车数据存在redis中,用hash
	//conn,err := redis.Dial("tcp","192.168.32.152:6379")
	//if err != nil {
	//	beego.Info("redis数据库链接错误")
	//	return
	//}
	//conn.Do()



}