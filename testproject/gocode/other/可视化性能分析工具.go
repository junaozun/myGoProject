package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
)

func main() {
	profiler.Start(profiler.Config{
		ApplicationName: "sxf",

		// replace this with the address of pyroscope server
		ServerAddress: "http://localhost:4040",

		// by default all profilers are enabled,
		//   but you can select the ones you want to use:
		ProfileTypes: []profiler.ProfileType{
			profiler.ProfileCPU,
			profiler.ProfileAllocObjects,
			profiler.ProfileAllocSpace,
			profiler.ProfileInuseObjects,
			profiler.ProfileInuseSpace,
		},
	})

	// your code goes here
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		longTimeTest()
		c.JSON(200, gin.H{"message": "ok"})
	})
	r.Run(":8080")
}

func longTimeTest() {
	a := 0
	for i := 0; i < 10000000000; i++ {
		a += i
	}
}

/*
启动工具服务
docker run -it -p 4040:4040 pyroscope/pyroscope:latest server
启动httpserver
go run XXXX.go ---- > localhost:8080访问网站

loclhost:4040 查看可视化页面


*/
