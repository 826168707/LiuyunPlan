package routers

import (
	"LedgerProject/controller"
	"LedgerProject/logic"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)


func SetupRouter() *gin.Engine {

	r := gin.Default()



	r.Use(controller.Cors())

	//使用zap日志
	r.Use(ginzap.Ginzap(zap.L(),time.RFC3339,true))
	r.Use(ginzap.RecoveryWithZap(zap.L(),true))

	//注册登录相关路由组
	v1Group := r.Group("sign")
	{
		//注册
		v1Group.PUT("/up",controller.UserRegistered)
		//发送验证码
		v1Group.POST("/up",controller.SendEmail)
		//登录
		v1Group.POST("",controller.UserLogin)
	}

	//主页路由组
	v2Group := r.Group("home")
	{
		//登录主页后页面获取信息
		v2Group.GET("",logic.JWTAuthMiddleware(),controller.GetHome)

		//设置金额 截止日期  日常固定支出
		v2Group.PUT("",logic.JWTAuthMiddleware(),controller.SetHome)

		//退出登录
		v2Group.POST("/out",logic.JWTAuthMiddleware(),controller.UserSignOut)

	}

	//支出 收录 路由组
	v3Group := r.Group("set")
	{
		//想要添加特殊支出
		v3Group.POST("/cost",logic.JWTAuthMiddleware(),controller.WantCost)
		//确认支出
		v3Group.PUT("/cost",logic.JWTAuthMiddleware(),controller.AddCost)
		//添加收入
		v3Group.PUT("/income",logic.JWTAuthMiddleware(),controller.AddIncome)
	}

	//历史记录路由组
	v4Group := r.Group("history")
	{
		//支出历史记录
		v4Group.GET("/cost",logic.JWTAuthMiddleware(),controller.CostHistory)
		//修改历史记录
		v4Group.PUT("",logic.JWTAuthMiddleware(),controller.UpdateHistory)
		//收入历史记录
		v4Group.GET("/income",logic.JWTAuthMiddleware(),controller.IncomeHistory)
		//删除历史记录
		v4Group.DELETE("",logic.JWTAuthMiddleware(),controller.DeleteHistory)
	}

	//推荐路由
	r.GET("/recommend",logic.JWTAuthMiddleware(),controller.Recommend)

	//上传图片
	r.POST("/upload",logic.JWTAuthMiddleware(),controller.UploadFile)

	//获取图片
	r.GET("/show_img",controller.ShowImg)

	return r
}



