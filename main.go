/*
 * @Author: liziwei01
 * @Date: 2022-03-03 15:33:30
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-02 21:27:04
 * @Description: main
 */
package main

import (
	"log"

	"github.com/liziwei01/gin-file-download/bootstrap"
	"github.com/liziwei01/gin-file-download/httpapi"
	"github.com/liziwei01/gin-file-download/library/env"
)

func main() {
	// priKeyPath := getRSAKeyPath()
	// priKey, _ := ioutil.ReadFile(priKeyPath)
	// b := utils.Encrypt.Base64Encode([]byte("liziwei01"))
	// fmt.Println(gconv.String(b))
	// rsaDecodedPasswd, err := utils.Encrypt.RsaPublicEncrypt(b, priKey)
	// b = utils.Encrypt.Base64Encode(rsaDecodedPasswd)
	// fmt.Println(gconv.String(b))
	// fmt.Println(err)
	// pwd, _ := utils.Encrypt.PasswordHashString("liziwei01")
	// b := utils.Encrypt.Base64Encode([]byte(pwd))
	// fmt.Println(gconv.String(b))

	// fmt.Println(utils.Encrypt.PasswordVerifyString("liziwei01", pwd))
	app, err := bootstrap.Setup()
	if err != nil {
		log.Fatalln(err)
	}
	// 注册接口路由
	httpapi.InitRouters(app.Handler)

	app.Start()
}

func getRSAKeyPath() string {
	return env.ConfDir() + "/pem/rsa_public_key.pem"
}
