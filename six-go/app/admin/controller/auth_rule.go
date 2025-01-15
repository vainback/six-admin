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

type AuthRuleController struct {
	dao.Query[models.AuthRule]
}

var AuthRule = new(AuthRuleController)

func (ts *AuthRuleController) List(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var total int64
	var list []models.AuthRule
	if err := ts.DB().Query().List(&list, &total); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.JsonPage(c, response.OkDbSelect, list, total)
	}
}

func (ts *AuthRuleController) TreeSelect(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var list []models.AuthRule
	if err := ts.DB().Query().Find(&list); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		var parentId int64 = 0
		if len(list) > 0 {
			root := slices.MinFunc[[]models.AuthRule](list, func(a, b models.AuthRule) int {
				return cmp.Compare(a.ParentId, b.ParentId)
			})
			parentId = root.ParentId
		}
		response.Json(c, response.OkDbSelect, utils.NewTree(list).Parser(parentId))
	}
}

func (ts *AuthRuleController) Select(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var list []models.AuthRule
	if err := ts.DB().Query().Find(&list); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		var parentId int64 = 0
		if len(list) > 0 {
			root := slices.MinFunc[[]models.AuthRule](list, func(a, b models.AuthRule) int {
				return cmp.Compare(a.ParentId, b.ParentId)
			})
			parentId = root.ParentId
		}
		response.Json(c, response.OkDbSelect, utils.NewTree(list).Parser(parentId))
	}
}

func (ts *AuthRuleController) Get(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var row models.AuthRule
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

func (ts *AuthRuleController) Add(c *gin.Context) {
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

func (ts *AuthRuleController) Update(c *gin.Context) {
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

func (ts *AuthRuleController) Save(c *gin.Context) {
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

func (ts *AuthRuleController) Delete(c *gin.Context) {
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
