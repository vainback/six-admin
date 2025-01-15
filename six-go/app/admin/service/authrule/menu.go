package authruleService

import "six-go/database/models"

type MenuMeta struct {
	Title        string   `json:"locale,omitempty"`
	Icon         string   `json:"icon,omitempty"`
	Roles        []string `json:"roles"`
	KeepAlive    *bool    `json:"ignoreCache,omitempty"`
	Order        int      `json:"order"`
	RequiresAuth bool     `json:"requiresAuth"`
}

type Menu struct {
	Path      string   `json:"path,omitempty"`
	Name      string   `json:"name,omitempty"`
	Redirect  string   `json:"redirect,omitempty"`
	Component string   `json:"component,omitempty"`
	Meta      MenuMeta `json:"meta"`
	Children  []Menu   `json:"children,omitempty"`
}

func ParserMenu(list []models.AuthRuleForMenu, pid int64) []Menu {
	var tree []Menu
	for _, v := range list {
		val := Menu{
			Path:      v.Path,
			Name:      v.Name,
			Component: v.Component,
			Meta: MenuMeta{
				Title:        v.Locale,
				Icon:         v.Icon,
				Roles:        v.Roles,
				Order:        v.Order,
				RequiresAuth: true,
			},
		}

		if v.ParentId == pid {
			child := ParserMenu(list, v.Id)
			if len(child) > 0 {
				val.Children = child
			}
			tree = append(tree, val)
		}
	}
	return tree
}
