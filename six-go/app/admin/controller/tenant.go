package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"six-go/database/dao"
	"six-go/database/models"
	"six-go/utils/response"
)

type TenantController struct {
	dao.Query[models.Tenant]
}

var Tenant = new(TenantController)

func (ts *TenantController) List(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var total int64
	var list []models.Tenant
	if err := ts.DB().Query().List(&list, &total); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.JsonPage(c, response.OkDbSelect, list, total)
	}
}

func (ts *TenantController) Select(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var list []models.Tenant
	if err := ts.DB().Query().Find(&list); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.Json(c, response.OkDbSelect, list)
	}
}

func (ts *TenantController) Get(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var row models.Tenant
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

func (ts *TenantController) Add(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	if err := ts.M.Valid(); err != nil {
		response.JsonError(c, err.Error(), err)
		return
	}

	if err := ts.DB().Create(); err != nil {
		response.JsonError(c, response.ErrorDbCreate, err)
	} else {
		response.Json(c, response.OkDbCreate, nil)
	}
}

func (ts *TenantController) Update(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	if err := ts.M.Valid(true); err != nil {
		response.JsonError(c, err.Error(), err)
		return
	}

	if err := ts.DB().Query().Update(); err != nil {
		response.JsonError(c, response.ErrorDbUpdate, err)
	} else {
		response.Json(c, response.OkDbUpdate, nil)
	}
}

func (ts *TenantController) Save(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}
	if err := ts.M.Valid(true); err != nil {
		response.JsonError(c, err.Error(), err)
		return
	}
	if err := ts.DB().Save(); err != nil {
		response.JsonError(c, response.ErrorDbUpdate, err)
	} else {
		response.Json(c, response.OkDbUpdate, nil)
	}
}

func (ts *TenantController) Delete(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	if err := ts.DB().Query().Delete(); err != nil {
		response.JsonError(c, response.ErrorDbDelete, err)
	} else {
		response.Json(c, response.OkDbDelete, nil)
	}
}
