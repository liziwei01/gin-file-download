/*
 * @Author: liziwei01
 * @Date: 2022-03-03 15:33:30
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-04-06 16:29:01
 * @Description: main
 */
package main

import (
	"log"

	"github.com/liziwei01/gin-file-download/httpapi"
	"github.com/liziwei01/gin-lib/bootstrap"
)

func main() {
	app, err := bootstrap.Setup()
	if err != nil {
		log.Fatalln(err)
	}
	// 注册接口路由
	httpapi.InitRouters(app.Handler)

	app.Start()
}
