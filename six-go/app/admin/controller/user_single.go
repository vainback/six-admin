package controller

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
	authruleService "six-go/app/admin/service/authrule"
	"six-go/extra/xredis"
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
	"six-go/app/admin/middleware"
	loginService "six-go/app/admin/service/login"
	"six-go/database/dao"
	"six-go/database/db"
	"six-go/database/models"
	"six-go/extra/uploader"
	"six-go/utils"
	"six-go/utils/arrays"
	"six-go/utils/response"
)

type UserSingleController struct {
	dao.Query[models.AuthUser]
}

var UserSingle = new(UserSingleController)

func (ts *UserSingleController) GetTenant(c *gin.Context) {
	var list []models.Tenant
	if err := db.DB().Find(&list).Error; err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
		return
	} else {
		response.Json(c, response.OkDbSelect, list)
		return
	}
}

// Login 登录接口
func (ts *UserSingleController) Login(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var tenant models.Tenant
	if err := db.DB().First(&tenant, "id = ?", ts.M.TenantId).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			response.JsonError(c, response.ErrorDbSelect, err)
			return
		}
		tenant.Status = 0
	}

	if tenant.Status != 1 {
		response.JsonWarn(c, "租户未启用")
		return
	}

	var row models.AuthUser
	if err := ts.DB().Query().First(&row); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.JsonWarn(c, response.ErrorUsernameOrPassword)
			return
		}
		response.JsonError(c, response.ErrorLogin, err)
		return
	}
	if !ts.M.Password.Verify(row.Password) {
		response.JsonWarn(c, response.ErrorUsernameOrPassword)
		return
	}
	if row.Status != 1 {
		response.JsonWarn(c, "账号状态异常！")
		return
	}

	token := utils.EncodeToken(row.Username + time.Now().String() + c.ClientIP())

	if err := loginService.New(token).Set(&row); err != nil {
		response.JsonError(c, response.ErrorLogin, err)
		return
	}

	response.Json(c, response.OkLogin, gin.H{
		"token":    token,
		"userinfo": row,
		"roles":    arrays.Number2String(row.RoleIds),
	})
}

func (ts *UserSingleController) Logout(c *gin.Context) {
	userinfo := middleware.SessionUser(c)
	token := middleware.GetHeaderToken(c)
	if userinfo != nil && token != "" {
		loginService.New(token).Clear()
	}

	response.Json(c, "ok", nil)
}

func (ts *UserSingleController) Update(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(&ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}
	sessionUser := middleware.SessionUser(c)
	if ts.M.Id > 0 && ts.M.Id != sessionUser.Id {
		response.JsonNoAuth(c)
		return
	}

	ts.M.Id = sessionUser.Id
	if err := ts.M.Valid(); err != nil {
		response.JsonWarn(c, err.Error())
		return
	}

	// 空值不会被update方法更新
	ts.M.Password = ""
	ts.M.Status = 0
	ts.M.RoleIds = nil

	if err := ts.DB().Query().Update(); err != nil {
		response.JsonError(c, response.ErrorDbUpdate, err)
		return
	} else {
		response.Json(c, response.OkDbUpdate, nil)
		return
	}
}

func (ts *UserSingleController) ResetPassword(c *gin.Context) {
	type Input struct {
		OldPassword utils.Password `json:"old_password"`
		NewPassword utils.Password `json:"new_password"`
	}
	var input Input
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	if err := validation.Required.Validate(input.OldPassword); err != nil {
		response.JsonError(c, "原密码必须填写", nil)
		return
	}

	if err := validation.Required.Validate(input.NewPassword); err != nil {
		response.JsonError(c, "新密码必须填写", nil)
		return
	}

	if err := validation.Length(6, 16).Validate(input.NewPassword); err != nil {
		response.JsonError(c, "密码长度必须在6-16位之间", nil)
		return
	}

	sessionUser := middleware.SessionUser(c)
	if err := db.DB().First(&sessionUser, "id = ?", sessionUser.Id).Error; err != nil {
		response.JsonError(c, response.ErrorDbUpdate, err)
		return
	}

	if !input.OldPassword.Verify(sessionUser.Password) {
		response.JsonWarn(c, "原密码输入有误")
		return
	}

	if err := db.DB().Table(models.TableAuthUser.TableName()).Where("id = ?", sessionUser.Id).Updates(map[string]any{
		"password":    input.NewPassword.Hash(),
		"update_time": time.Now(),
	}).Error; err != nil {
		response.JsonError(c, response.ErrorDbUpdate, err)
		return
	}

	// 以下方式会被GORM Struct Tag 阻止修改 root 用户的密码
	//sessionUser.Password = input.NewPassword.Hash()
	//
	//if err := db.DB().Save(sessionUser).Error; err != nil {
	//	response.JsonError(c, response.ErrorDbUpdate, err)
	//	return
	//}

	response.Json(c, response.OkDbUpdate, nil)
	loginService.ClearAll(sessionUser.Username)
	return
}

func (ts *UserSingleController) Menu(c *gin.Context) {
	errMessage := "菜单查询失败"

	sessionUser := middleware.SessionUser(c)

	key := fmt.Sprintf("%s:admin-menu:%s", xredis.KeyPrefix, middleware.GetHeaderToken(c))
	value := xredis.GetString(key)
	var treeMenu []authruleService.Menu
	if err := jsoniter.UnmarshalFromString(value, &treeMenu); err == nil {
		response.Json(c, "ok", treeMenu)
		return
	}

	// 构造查询条件
	var query = &dao.Query[models.AuthRule]{
		M: models.AuthRule{
			Type:   models.RuleTypesMenu,
			Status: 1,
		},
		QueryOrderBys: dao.QueryOrderBys{
			OrderBy: []dao.QueryOrderBy{
				{Field: "`order`", IsDesc: false},
			},
		},
	}

	var rules []models.AuthRule
	if err := query.DB().Query().Find(&rules); err != nil {
		response.JsonError(c, errMessage, err)
		return
	}

	var relations []models.AuthRelation
	if err := db.DB().Scopes(sessionUser.TenantIdentify.SqlBuilder()).Find(&relations).Error; err != nil {
		response.JsonError(c, errMessage, err)
		return
	}

	var relationMapSlices = make(map[int64][]int64)
	for _, v := range relations {
		if _, ok := relationMapSlices[v.RuleId]; !ok {
			relationMapSlices[v.RuleId] = make([]int64, 0)
		}
		relationMapSlices[v.RuleId] = append(relationMapSlices[v.RuleId], v.RoleId)
	}

	var results = make([]models.AuthRuleForMenu, len(rules))
	for i, rule := range rules {
		var ids = make([]int64, 0)
		if relation, ok := relationMapSlices[rule.Id]; ok {
			ids = relation
		}
		ids = append(ids, models.RootRoleId)
		results[i] = models.AuthRuleForMenu{
			AuthRule: rule,
			Roles:    arrays.Number2String(ids),
		}
	}
	treeMenu = authruleService.ParserMenu(results, 0)

	marshal, _ := jsoniter.Marshal(treeMenu)
	_ = xredis.SetExp(key, string(marshal), time.Hour*1)

	response.Json(c, "ok", treeMenu)
}

func (ts *UserSingleController) Btns(c *gin.Context) {
	errMessage := "菜单查询失败"

	sessionUser := middleware.SessionUser(c)

	key := fmt.Sprintf("%s:admin-buttons:%s", xredis.KeyPrefix, middleware.GetHeaderToken(c))
	value := xredis.GetString(key)
	var buttons []string
	if err := jsoniter.UnmarshalFromString(value, &buttons); err == nil {
		response.Json(c, "ok", buttons)
		return
	}

	// 构造查询条件
	var query = &dao.Query[models.AuthRule]{
		M: models.AuthRule{
			Type:   models.RuleTypesButton,
			Status: 1,
		},
	}

	tx := query.DB().Query().GORM().Model(&models.AuthRule{})

	if !sessionUser.IsRoot {
		var ids []int64
		if err := db.DB().Model(&models.AuthRelation{}).Where("role_id in ?", sessionUser.RoleIds).Select("rule_id").Scan(&ids).Error; err != nil {
			response.JsonError(c, errMessage, err)
			return
		}
		tx = tx.Where("id in ?", ids)
	}

	var auths []string
	if err := tx.Select("auth").Scan(&auths).Error; err != nil {
		log.Println(err)
	}

	marshal, _ := jsoniter.Marshal(auths)
	_ = xredis.SetExp(key, string(marshal), time.Hour*1)

	response.Json(c, "ok", auths)
}

func (ts *UserSingleController) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.JsonError(c, "获取上传的文件失败", err)
		return
	}

	res := uploader.New().Upload(file)
	if res.Error != nil {
		response.JsonError(c, "上传失败", res.Error)
		return
	}

	var createFile = models.Files{
		TenantIdentify: models.TenantIdentify{
			TenantId: middleware.SessionUser(c).Id,
		},
		Mime: file.Header.Get("Content-Type"),
		Type: "remote",
		Url:  res.Url,
	}

	if res.IsLocal {
		if err = c.SaveUploadedFile(file, res.Url); err != nil {
			response.JsonError(c, "上传失败", err)
			return
		}
		createFile.Type = "local"
	}

	if err = db.DB().Create(&createFile).Error; err != nil {
		res.Error = fmt.Errorf("文件上传成功，保存到附件列表失败：%s", err.Error())
		log.Println(res.Error)
	}

	response.Json(c, "ok", res)
}

func (ts *UserSingleController) Userinfo(c *gin.Context) {
	response.Json(c, "ok", *middleware.SessionUser(c))
}
