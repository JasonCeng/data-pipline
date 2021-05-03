package deployer

import (
	log "github.com/sirupsen/logrus"
)

//deployer模块的主要作用：
//1、读取properties:datapipline.yaml
//2、独立启instance
//3、监听instance的配置的变化，动态停止、启动或新增
//4、启动server，监听客户端请求
func main() {
	//1、读取properties

	//2、独立启动instance

	//3、监听instance的配置的变化，动态停止、启动或新增

	//4、启动server，监听客户端请求
}

func initLog() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
