package controller

import (
	"cmp"
	"errors"
	authService "six-go/app/admin/service/auth"
	"six-go/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"six-go/database/dao"
	"six-go/database/models"
	"six-go/utils/response"
)

type <Tpl-Module>Controller struct {
	dao.Query[models.<Tpl-Module>]
}

var <Tpl-Module> = new(<Tpl-Module>Controller)

func (ts *<Tpl-Module>Controller) List(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var total int64
	var list []models.<Tpl-Module>
	if err := ts.DB().Query().List(&list, &total); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.JsonPage(c, response.OkDbSelect, list, total)
	}
}

func (ts *<Tpl-Module>Controller) TreeSelect(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var list []models.<Tpl-Module>
	if err := ts.DB().Query().Find(&list); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		var parentId int64 = 0
		if len(list) > 0 {
			root := slices.MinFunc[[]models.<Tpl-Module>](list, func(a, b models.<Tpl-Module>) int {
				return cmp.Compare(a.ParentId, b.ParentId)
			})
			parentId = root.ParentId
		}
		response.Json(c, response.OkDbSelect, utils.NewTree(list).Parser(parentId))
	}
}

func (ts *<Tpl-Module>Controller) Select(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var list []models.<Tpl-Module>
	if err := ts.DB().Query().Find(&list); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		var parentId int64 = 0
		if len(list) > 0 {
			root := slices.MinFunc[[]models.<Tpl-Module>](list, func(a, b models.<Tpl-Module>) int {
				return cmp.Compare(a.ParentId, b.ParentId)
			})
			parentId = root.ParentId
		}
		response.Json(c, response.OkDbSelect, utils.NewTree(list).Parser(parentId))
	}
}

func (ts *<Tpl-Module>Controller) Get(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var row models.<Tpl-Module>
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

func (ts *<Tpl-Module>Controller) Add(c *gin.Context) {
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
		go authService.GetRules()
	}
}

func (ts *<Tpl-Module>Controller) Update(c *gin.Context) {
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
		go authService.GetRules()
	}
}

func (ts *<Tpl-Module>Controller) Save(c *gin.Context) {
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
		go authService.GetRules()
	}
}

func (ts *<Tpl-Module>Controller) Delete(c *gin.Context) {
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
		go authService.GetRules()
	}
}
