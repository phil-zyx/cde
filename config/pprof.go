package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http/pprof"
)

// RunProf 启动性能监控
func RunProf(r *gin.Engine) {
	// 将 pprof 与路由器实例关联
	r.GET("/debug/pprof/*name", gin.WrapH(http.DefaultServeMux))
}
