package main

import (

	"github.com/astaxie/beego"
	controller "./controller"
)

func main()  {
	beego.Router("/corona/qt",&controller.CoronaQT{})
	beego.Router("/corona/vn",&controller.CoronaVN{})
	beego.Run(":80")
}