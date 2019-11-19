package main

import (
	"fmt"
	"raspberry-go/utils"

	"github.com/astaxie/beego"
)

// AES 盐值
const salt = "B1827B657FFF9232"

// ======================================================================================
// Router
// ======================================================================================

func RegisterRouter() {

	// fish 请求
	beego.Router("/fish", &FishController{})

	beego.Router("/*", &MainController{}, "*:Do")
}

// ======================================================================================
// Controller
// ======================================================================================

type FishController struct {
	beego.Controller
}

func (c *FishController) Get() {

	// 捕获异常
	defer func() {
		if r := recover(); r != nil {
			// "%v" prints the value of expression
			// for strings, it is the string, for errors .Error() method, for Stringer the .String() etc
			// Errorf returns an error instead of a string
			c.Ctx.WriteString(fmt.Errorf("%v", r).Error())
		}
	}()

	// 入参中获取加密后的food
	encryptFood := c.GetString("food")
	key := salt

	// 如果可以成功解密标识这确实是服务器发送的请求
	food, _ := utils.AesDecryptWithBase64(encryptFood, key)
	if food != encryptFood {
		// TODO 这里执行真正的喂鱼操作
		c.Ctx.WriteString(food)
		return
	}

	panic("[" + encryptFood + "] fail")
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Do() {
	c.Ctx.WriteString("Not Support Request")
}

// ======================================================================================
// main
// ======================================================================================

func main() {

	RegisterRouter()

	beego.BConfig.Listen.HTTPPort = 11011
	beego.Run()
}
