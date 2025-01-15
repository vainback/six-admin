package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"six-go/app/admin/entity"
	"six-go/app/admin/middleware"
	"six-go/database/dao"
	"six-go/database/models"
	"six-go/utils/response"
)

type OperateLogsController struct {
	dao.Query[models.OperateLog]
}

var OperateLogs = new(OperateLogsController)

func (ts *OperateLogsController) List(c *gin.Context) {
	var query = new(dao.Query[entity.LogJoinUser])
	if err := c.ShouldBindBodyWithJSON(query); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}
	query.M.User.TenantId = middleware.SessionUser(c).TenantId

	var total int64
	var list []entity.LogJoinUser
	if err := query.DB().Query().JoinList(&list, &total, query.M.LeftJoin(), query.M.Select()); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.JsonPage(c, response.OkDbSelect, list, total)
	}
}

func (ts *OperateLogsController) Select(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var list []models.OperateLog
	if err := ts.DB().Query().Find(&list); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.Json(c, response.OkDbSelect, list)
	}
}

func (ts *OperateLogsController) Get(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var row models.OperateLog
	if err := ts.DB().Query().First(&row); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.JsonError(c, response.ErrorDbNotFound, err)
			return
		}
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.Json(c, response.OkDbSelect, row)
	}
}
