package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:AccessController"] = append(beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:AccessController"],
		beego.ControllerComments{
			Method: "CheckToken",
			Router: `/checkToken`,
			AllowHTTPMethods: []string{"post","get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:AccessController"] = append(beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:AccessController"],
		beego.ControllerComments{
			Method: "GetToken",
			Router: `/getToken`,
			AllowHTTPMethods: []string{"post","get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:CaptchaController"] = append(beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:CaptchaController"],
		beego.ControllerComments{
			Method: "CaptchaVerifyHandle",
			Router: `/captchaVerify`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:CaptchaController"] = append(beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:CaptchaController"],
		beego.ControllerComments{
			Method: "CheckMobileCode",
			Router: `/checkMobileCode`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:CaptchaController"] = append(beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:CaptchaController"],
		beego.ControllerComments{
			Method: "GenerateCaptchaHandler",
			Router: `/generateCaptcha`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:CaptchaController"] = append(beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:CaptchaController"],
		beego.ControllerComments{
			Method: "MobileCode",
			Router: `/getMobileCode`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:ResourceController"] = append(beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:ResourceController"],
		beego.ControllerComments{
			Method: "GetTreeGrid",
			Router: `/getTreeGrid`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "DoLogin",
			Router: `/doLogin`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/gtck520/kcapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUserInfo",
			Router: `/getUserInfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
