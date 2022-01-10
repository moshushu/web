package main

import (
	"fmt"
	"web/dao/mysql"
	"web/logger"
	"web/routes"
	"web/settings"

	"go.uber.org/zap"
)

//搭建比较通用的web脚手架模板

func main() {
	// 1：加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	// 2：初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync() //注册关闭全局logger
	zap.L().Debug("init logger success ....")
	// 3：初始化mysql连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	// 4：初始化Redis连接
	// 5：注册路由
	r := routes.Routes(settings.Conf.Mode)
	// 6：启动服务
	r.Run()
}
