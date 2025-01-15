package controller

import (
	"cmp"
	"errors"
	authService "six-go/app/admin/service/auth"
	"six-go/utils"
	"slices"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"six-go/app/admin/middleware"
	"six-go/database/dao"
	"six-go/database/models"
	"six-go/utils/response"
)

type AuthRoleController struct {
	dao.Query[models.AuthRole]
}

var AuthRole = new(AuthRoleController)

func (ts *AuthRoleController) List(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}
	user := middleware.SessionUser(c)

	var total int64
	var list []models.AuthRole
	tx := ts.DB().Query()
	if !user.IsRoot {
		sons := models.NewSon(user.RoleIds...).GetIds(ts.M, func(db *gorm.DB) *gorm.DB {
			return db.Where("status = ?", true)
		})
		tx = tx.Where("id in ?", sons)
	}
	if err := tx.List(&list, &total); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.JsonPage(c, response.OkDbSelect, list, total)
	}
}

func (ts *AuthRoleController) Select(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}
	user := middleware.SessionUser(c)

	var list []models.AuthRole
	tx := ts.DB().Query()
	if !user.IsRoot {
		sons := models.NewSon(user.RoleIds...).GetIds(ts.M, func(db *gorm.DB) *gorm.DB {
			return db.Where("status = ?", true)
		})
		tx = tx.Where("id in ?", sons)
	}
	if err := tx.Find(&list); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		response.Json(c, response.OkDbSelect, list)
	}
}

func (ts *AuthRoleController) TreeSelect(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}
	user := middleware.SessionUser(c)

	var list []models.AuthRole
	tx := ts.DB().Query()
	if !user.IsRoot {
		sons := models.NewSon(user.RoleIds...).GetIds(ts.M, func(db *gorm.DB) *gorm.DB {
			return db.Where("status = ?", true)
		})
		tx = tx.Where("id in ?", sons)
	}
	if err := tx.Find(&list); err != nil {
		response.JsonError(c, response.ErrorDbSelect, err)
	} else {
		var parentId int64 = 0
		if len(list) > 0 {
			root := slices.MinFunc[[]models.AuthRole](list, func(a, b models.AuthRole) int {
				return cmp.Compare(a.ParentId, b.ParentId)
			})
			parentId = root.ParentId
		}
		response.Json(c, response.OkDbSelect, utils.NewTree(list).Parser(parentId))
	}
}

func (ts *AuthRoleController) Get(c *gin.Context) {
	if err := c.ShouldBindBodyWithJSON(ts); err != nil {
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	var row models.AuthRole
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

func (ts *AuthRoleController) Add(c *gin.Context) {
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
		go authService.GetRoles()
	}
}

func (ts *AuthRoleController) Update(c *gin.Context) {
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
		go authService.GetRoles()
	}
}

func (ts *AuthRoleController) Save(c *gin.Context) {
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
		go authService.GetRoles()
	}
}

func (ts *AuthRoleController) Delete(c *gin.Context) {
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
		go authService.GetRoles()
	}
}
