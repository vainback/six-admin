package controller

import (
	"errors"
	"six-go/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"six-go/database/dao"
	"six-go/database/models"
	"six-go/utils/response"
)

type JobController struct {
	dao.Query[models.Job]
}

var Job = new(JobController)

func (ts *JobController) List(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var total int64
	var list []models.Job
	if err := ts.DB().Query().List(&list, &total); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.JsonPage(c, response.OkDbSelect, list, total)
	}
}

func (ts *JobController) Select(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var list []models.Job
	if err := ts.DB().Query().Find(&list); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.Json(c, response.OkDbSelect, list)
	}
}

func (ts *JobController) TreeSelect(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var list []models.Job
	if err := ts.DB().Query().Find(&list); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.Json(c, response.OkDbSelect, utils.NewTree(list).Parser(0))
	}
}

func (ts *JobController) Get(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var row models.Job
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

func (ts *JobController) Add(c *gin.Context) {
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

func (ts *JobController) Update(c *gin.Context) {
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

func (ts *JobController) Save(c *gin.Context) {
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

func (ts *JobController) Delete(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	if nums, err := ts.M.HasChildren(); err != nil {
		response.JsonError(c, response.ErrorDbDelete, err)
		return
	} else if nums > 0 {
		response.JsonWarn(c, "有子节点存在，请先删除子节点")
		return
	}

	if err := ts.DB().Query().Delete(); err != nil {
		response.JsonError(c, response.ErrorDbDelete, err)
	} else {
		response.Json(c, response.OkDbDelete, nil)
	}
}
