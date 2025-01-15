package controller

import (
	"errors"
	"six-go/extra/uploader"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"six-go/database/dao"
	"six-go/database/models"
	"six-go/utils/response"
)

type FilesController struct {
	dao.Query[models.Files]
}

var Files = new(FilesController)

func (ts *FilesController) List(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var total int64
	var list []models.Files
	if err := ts.DB().Query().List(&list, &total); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.JsonPage(c, response.OkDbSelect, list, total)
	}
}

func (ts *FilesController) Select(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var list []models.Files
	if err := ts.DB().Query().Find(&list); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.Json(c, response.OkDbSelect, list)
	}
}

func (ts *FilesController) Get(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var row models.Files
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

func (ts *FilesController) Update(c *gin.Context) {
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

func (ts *FilesController) Save(c *gin.Context) {
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

func (ts *FilesController) Delete(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	if err := ts.DB().Query().Delete(); err != nil {
		response.JsonError(c, response.ErrorDbDelete, err)
	} else {
		// 真实删除 移除文件
		if ts.IsDelete {
			_ = uploader.New().Remove(ts.M.Url)
		}

		response.Json(c, response.OkDbDelete, nil)
	}
}
