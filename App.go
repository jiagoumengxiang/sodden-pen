package main
import (
	"github.com/astaxie/beego"
	"github.com/jiagoumengxiang/sodden-pen/controller"
)


func main(){
	beego.SetStaticPath("/app","Views")
	beego.Router("/api/login",&controller.Login{})
	beego.Run()
}
