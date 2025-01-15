package models

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"six-go/database/db"
)

type Son struct {
	root []int64
	list []map[string]any
}

func NewSon(id ...int64) *Son {
	return &Son{root: id, list: nil}
}

func (s *Son) GetIds(model db.ModelIof, where ...func(db *gorm.DB) *gorm.DB) []int64 {
	if len(s.root) == 0 {
		return nil
	}

	var list []map[string]any
	tx := db.DB().Table(model.TableName())
	if len(where) > 0 {
		tx = tx.Scopes(where...)
	}
	if err := tx.Scan(&list).Error; err != nil {
		return nil
	}
	s.list = list
	var sons []int64
	for _, v := range s.root {
		sons = append(sons, s.parserSons(v)...)
	}
	return sons
}

func (s *Son) parserSons(id int64) []int64 {
	if len(s.list) == 0 {
		return nil
	}
	var sons []int64
	for _, v := range s.list {
		vid := cast.ToInt64(v["id"])
		if cast.ToInt64(v["parent_id"]) == id {
			if childs := s.parserSons(vid); len(childs) > 0 {
				sons = append(sons, childs...)
			}
			sons = append(sons, vid)
		}
	}
	return sons
}
