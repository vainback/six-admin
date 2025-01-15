package controller

import (
	"errors"
	loginService "six-go/app/admin/service/login"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"six-go/app/admin/entity"
	"six-go/database/dao"
	"six-go/database/models"
	"six-go/utils/response"
)

type AuthUserController struct {
	dao.Query[models.AuthUser]
}

var AuthUser = new(AuthUserController)

func (ts *AuthUserController) List(c *gin.Context) {
	var query = new(dao.Query[entity.AuthUserJoinJob])
	if err := c.ShouldBindBodyWithJSON(query); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}
	var total int64
	var list []entity.AuthUserJoinJob
	if err := query.DB().Query().JoinList(&list, &total, query.M.LeftJoin(), query.M.Select()); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.JsonPage(c, response.OkDbSelect, list, total)
	}
}

func (ts *AuthUserController) Select(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var list []models.AuthUser
	if err := ts.DB().Query().Find(&list); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.Json(c, response.OkDbSelect, list)
	}
}

func (ts *AuthUserController) Get(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var row models.AuthUser
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

func (ts *AuthUserController) Add(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	ts.M.Password = ts.M.Password.TrimSpace()
	if err := ts.M.Valid(); err != nil {
		response.JsonError(c, err.Error(), err)
		return
	}
	ts.M.IsRoot = false
	ts.M.Password = ts.M.Password.Hash()

	if err := ts.DB().Create(); err != nil {
		response.JsonError(c, response.ErrorDbCreate, err)
	} else {
		response.Json(c, response.OkDbCreate, nil)
	}
}

func (ts *AuthUserController) Update(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	ts.M.Password = ts.M.Password.TrimSpace()
	if err := ts.M.Valid(true); err != nil {
		response.JsonError(c, err.Error(), err)
		return
	}

	if ts.M.Password != "" {
		ts.M.Password = ts.M.Password.Hash()
	}

	if err := ts.DB().Query().Where("is_root = ?", false).Update(); err != nil {
		response.JsonError(c, response.ErrorDbUpdate, err)
	} else {
		response.Json(c, response.OkDbUpdate, nil)
		loginService.ClearAll(ts.M.Username)
	}
}

func (ts *AuthUserController) Save(c *gin.Context) {
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
		loginService.ClearAll(ts.M.Username)
	}
}

func (ts *AuthUserController) Delete(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	if err := ts.DB().Query().Where("is_root = ?", false).Delete(); err != nil {
		response.JsonError(c, response.ErrorDbDelete, err)
	} else {
		response.Json(c, response.OkDbDelete, nil)
		loginService.ClearAll(ts.M.Username)
	}
}
