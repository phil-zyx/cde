package util

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 获取全部请求参数
func DataMapByRequest(ctx *gin.Context) (dataMap map[string]any, err error) {
	//必须先解析Form
	err = ctx.Request.ParseForm()
	dataMap = make(map[string]any)
	//说明:须post方法,加: 'Content-Type': 'application/x-www-form-urlencoded'
	for key := range ctx.Request.PostForm {
		dataMap[key] = ctx.PostForm(key)
	}
	// 获取Url上的请求参数
	for key := range ctx.Request.URL.Query() {
		dataMap[key] = ctx.Query(key)
	}
	return
}

func Response(ctx *gin.Context, httpStatus int, code int, msg string, data any) {
	ctx.JSON(httpStatus, gin.H{"code": code, "msg": msg, "data": data})
}

func Success(ctx *gin.Context, msg string, data any) {
	Response(ctx, http.StatusOK, 200, msg, data)
}

func Fail(ctx *gin.Context, msg string, data any) {
	Response(ctx, http.StatusOK, 400, msg, data)
}

// RandomString 生成指定长度的随机字符
func RandomString(n int) string {
	var letters = []byte("qwertyuioplkjhgfdsazxcvbnmMNBVCXZASDFGHJKLPOIUYTREWQ")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// ExternalIp 获取外网IP
func ExternalIp() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

// 格式化 IP
func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil
	}
	return ip
}

// ParamsFilter 过滤指定数组中的key
func ParamsFilter(isFilterStr string, params map[string]any) map[string]any {
	var data = make(map[string]any)
	for key, value := range params {
		if value != "" {
			find := strings.Contains(isFilterStr, key)
			if !find {
				data[key] = value
			}
		}
	}
	return data
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //文件或者目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// AnyToMap interface 转 map
func AnyToMap(v any) (map[string]any, error) {
	dataJson, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var MapData map[string]any
	err = json.Unmarshal(dataJson, &MapData)
	if err != nil {
		return nil, err
	}
	return MapData, nil
}
