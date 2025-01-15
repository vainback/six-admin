package controller

import (
	"fmt"
	authService "six-go/app/admin/service/auth"
	loginService "six-go/app/admin/service/login"
	"slices"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"six-go/app/admin/middleware"
	"six-go/database/dao"
	"six-go/database/db"
	"six-go/database/models"
	"six-go/utils/response"
)

type AuthRelationController struct {
	dao.Query[models.AuthRelation]
}

var AuthRelation = new(AuthRelationController)

func (ts *AuthRelationController) SelectGroupRole(c *gin.Context) {
	user := middleware.SessionUser(c)
	ts.M.TenantId = user.TenantId
	var list []models.AuthRelation
	if err := ts.DB().Find(&list); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
		return
	}

	var rules []models.AuthRule
	if err := db.DB().Find(&rules, "status = ?", 1).Error; err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
		return
	}

	var rulesMap = make(map[int64]models.AuthRule)
	for _, v := range rules {
		rulesMap[v.Id] = v
	}

	var result = make(map[int64][]models.AuthRule)
	for _, v := range list {
		if _, ok := result[v.RoleId]; !ok {
			result[v.RoleId] = make([]models.AuthRule, 0)
		}

		if value, ok := rulesMap[v.RuleId]; ok {
			result[v.RoleId] = append(result[v.RoleId], value)
		}
	}

	response.Json(c, response.OkDbSelect, result)
}

func (ts *AuthRelationController) SelectGroupRule(c *gin.Context) {
	user := middleware.SessionUser(c)
	ts.M.TenantId = user.TenantId
	var list []models.AuthRelation
	if err := ts.DB().Find(&list); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
		return
	}

	var roles []models.AuthRole
	if err := db.DB().Find(&roles, "status = ?", 1).Error; err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
		return
	}

	var rolesMap = make(map[int64]models.AuthRole)
	for _, v := range roles {
		rolesMap[v.Id] = v
	}

	var result = make(map[int64][]models.AuthRole)
	for _, v := range list {
		if _, ok := result[v.RuleId]; !ok {
			result[v.RuleId] = make([]models.AuthRole, 0)
		}

		if value, ok := rolesMap[v.RoleId]; ok {
			result[v.RuleId] = append(result[v.RuleId], value)
		}
	}

	response.Json(c, response.OkDbSelect, result)
}

func (ts *AuthRelationController) SelectRuleIdsWithRoleId(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	if ts.M.RoleId < 1 {
		response.JsonError(c, "缺少参数", nil)
		return
	}

	var ids []int64
	if err := ts.DB().Query().GORM().Table(ts.M.TableName()).Select("rule_id").Scan(&ids).Error; err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.Json(c, response.OkDbSelect, ids)
	}
}

func (ts *AuthRelationController) Set(c *gin.Context) {
	type Input struct {
		RoleId  int64   `json:"role_id"`
		RuleIds []int64 `json:"rule_ids"`
	}
	var input Input
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var row models.AuthRole
	if err := db.DB().First(&row, "id = ?", input.RoleId).Error; err != nil {
		response.JsonError(c, response.ErrorSetting, err)
		return
	}

	if row.TenantId != middleware.SessionUser(c).TenantId {
		response.JsonNoAuth(c)
		return
	}

	if err := db.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.AuthRelation{}, "role_id = ? and rule_id not in ?", input.RoleId, input.RuleIds).Error; err != nil {
			return err
		}
		var exists []models.AuthRelation
		if err := tx.Find(&exists, "role_id = ?", input.RoleId).Error; err != nil {
			return err
		}

		if len(input.RuleIds) > 0 {
			var inserts []models.AuthRelation
			for _, ruleId := range input.RuleIds {
				// 未存在数据库的rule_id
				isExist := slices.IndexFunc(exists, func(relation models.AuthRelation) bool {
					return relation.RuleId == ruleId
				})
				if isExist == -1 {
					inserts = append(inserts, models.AuthRelation{
						RoleId:         input.RoleId,
						RuleId:         ruleId,
						TenantIdentify: middleware.SessionUser(c).TenantIdentify,
					})
				}
			}

			if err := tx.Create(&inserts).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		response.JsonError(c, response.ErrorSetting, err)
		return
	}

	response.Json(c, response.OkSetting, nil)
	go authService.LoadAllRelation()
	go func() {
		sql := fmt.Sprintf("SELECT username FROM %s WHERE JSON_CONTAINS(role_ids, '%d')", models.TableAuthUser.TableName(), input.RoleId)
		var users []string
		if err := db.DB().Raw(sql).Scan(&users).Error; err != nil {
			return
		}
		if len(users) > 0 {
			for _, user := range users {
				loginService.ClearAll(user)
			}
		}
	}()
}
