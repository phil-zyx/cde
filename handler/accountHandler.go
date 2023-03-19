package handler

import (
	"github.com/cde/database"
	"github.com/cde/model"
	"github.com/cde/repository"
	"github.com/cde/util"
	"github.com/cde/util/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @BasePath /api/v1

// 注册
// @Summary 注册
// @Schemes
// @Description 注册
// @Tags admin
// @Accept json
// @Produce json
// @Param user body model.Admin true "参数"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /account/register [post]
func Register(ctx *gin.Context) {
	db := database.GetDB()
	data := &model.Admin{}
	if err := ctx.ShouldBindJSON(data); err != nil {
		logger.Errorf("bind json err %v", err)
		util.Fail(ctx, "参数错误", nil)
		return
	}
	// 获取参数
	// params, _ := util.DataMapByRequest(ctx)
	// username := params["username"].(string)
	// password := params["password"].(string)
	// nickname := params["nickname"].(string)
	// phone := params["phone"].(string)
	password := data.Password
	username := data.Username
	nickname := data.UserNick
	phone := data.Phone

	if len(password) < 6 {
		util.Fail(ctx, "密码不能小于6位！", nil)
		return
	}

	if len(username) == 0 {
		username = util.RandomString(10)
	}

	if len(nickname) == 0 {
		nickname = util.RandomString(10)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		util.Fail(ctx, "加密错误", nil)
		return
	}

	user := model.Admin{
		Username: username,
		UserNick: nickname,
		Password: string(hashPassword),
		Phone:    phone,
	}
	err = repository.AdminRepository.Create(db, &user)
	if err != nil {
		util.Fail(ctx, "create db data err %v", err)
		return
	}
	util.Success(ctx, "注册成功", nil)
}

// Login 登录
// @Summary 登录
// @Schemes
// @Description 登录
// @Tags admin
// @Accept json
// @Produce json
// @Param data body model.Admin true "用户名"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /account/login [post]
func Login(ctx *gin.Context) {
	db := database.GetDB()
	data := &model.Admin{}
	if err := ctx.ShouldBind(data); err != nil {
		logger.Errorf("bind json err %v", err)
		util.Fail(ctx, "参数错误", nil)
		return
	}

	// 获取参数
	// params, _ := util.DataMapByRequest(ctx)
	// username := params["username"].(string)
	// password := params["password"].(string)
	//captchaId := params["captchaId"].(string)
	//captchaCode := params["captchaCode"].(string)

	username := data.Username
	password := data.Password
	// 表单验证
	if len(username) == 0 {
		util.Fail(ctx, "用户名不能为空", nil)
		return
	}
	if len(password) == 0 {
		util.Fail(ctx, "密码不能为空", nil)
		return
	}
	//if len(captchaCode) == 0 {
	//	util.Fail(ctx, "验证码不能为空", nil)
	//	return
	//}
	// 验证用户密码
	user := repository.AdminRepository.Get(db, username)
	if user.ID == 0 {
		util.Fail(ctx, "该用户未注册", nil)
		return
	}
	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		util.Fail(ctx, "密码错误", nil)
		return
	}
	//// 验证验证码
	//if !captcha.VerfiyCaptcha(captchaId, captchaCode) {
	//	util.Fail(ctx, "验证码不正确", nil)
	//	return
	//}
	// 生成token
	token, err := util.ReleaseToken(user.Username)
	if err != nil {
		util.Fail(ctx, "生成token失败", nil)
		return
	}
	// 获取 本机真实IP
	ip, _ := util.ExternalIp()
	// 更新user
	var localTime util.LocalTime
	updateData := make(map[string]interface{})
	updateData["login_ip"] = ip.String()
	updateData["login_at"] = localTime.FormatDateString(localTime.String())
	if err := repository.AdminRepository.Updates(db, int64(user.ID), updateData); err != nil {
		logger.Debugf("update data err %v", err)
	}
	// 返回值
	util.Success(ctx, "登录成功", gin.H{
		"token": token,
	})
}

func Logout(ctx *gin.Context) {
	util.Success(ctx, "退出成功", nil)
}
