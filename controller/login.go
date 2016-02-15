package controller
import "github.com/astaxie/beego"

type Login struct {
	beego.Controller
}

func (this *Login)Get(){
	this.ServeJSON()
}
