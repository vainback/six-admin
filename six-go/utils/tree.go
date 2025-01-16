package utils

import (
	"github.com/spf13/cast"
	"github.com/vainback/six-util/v3"
)

const (
	TreeIDKey    = "id"
	TreePIDKey   = "parent_id"
	TreeChildKey = "children"
)

type Tree struct {
	list []map[string]any
}

func NewTree[T any](list []T) *Tree {
	return &Tree{six.Structs2SliceMap(list)}
}

func NewTreeWithMapAny(list []map[string]any) *Tree {
	return &Tree{list}
}

func (t *Tree) Parser(pid int64) []map[string]any {
	var tree []map[string]any
	for _, v := range t.list {
		tPid := cast.ToInt64(v[TreePIDKey])
		if tPid == pid {
			child := t.Parser(cast.ToInt64(v[TreeIDKey]))
			if len(child) > 0 {
				v[TreeChildKey] = child
			}
			tree = append(tree, v)
		}
	}
	return tree
}
