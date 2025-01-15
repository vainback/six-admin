package dao

import (
	"gorm.io/gorm"
	"six-go/database/db"
)

type Query[T db.ModelIof] struct {
	db              *gorm.DB // db
	M               T        `json:"model"`
	QueryPage                // 分页查询
	QueryTime                // 时间查询
	QueryKeyword             // 关键字查询
	QueryOrderBys            // 排序
	QuerySoftDelete          // 软删除 删除接口该值为true则执行物理删除 查询接口该值为true将查询被软删除的数据
}

func (q *Query[T]) DB() *Query[T] {
	q.db = db.DB()
	return q
}

func (q *Query[T]) Where(where string, args ...any) *Query[T] {
	q.db = q.db.Where(where, args...)
	return q
}

func (q *Query[T]) Query() *Query[T] {
	q.db = q.db.Scopes(
		q.QuerySoftDelete.SqlBuilder(),
		q.M.FilterSqlBuilder(),
		q.QueryKeyword.SqlBuilder(q.M.KeywordFields()...),
		q.QueryTime.SqlBuilder(q.M.TableName()),
	).Session(&gorm.Session{})
	return q
}

func (q *Query[T]) Scopes(scopes ...func(*gorm.DB) *gorm.DB) *Query[T] {
	q.db = q.db.Scopes(scopes...)
	return q
}

// First 单条查询
func (q *Query[T]) First(row any) error {
	return q.db.Scopes(q.QueryOrderBys.SqlBuilder()).First(row).Error
}

// Find 列表查询
func (q *Query[T]) Find(list any) error {
	return q.db.Scopes(q.QueryOrderBys.SqlBuilder()).Find(list).Error
}

// List 分页列表查询
func (q *Query[T]) List(list any, total *int64) error {
	if err := q.db.Model(q.M).Count(total).Error; err != nil {
		return err
	}

	return q.db.Scopes(q.QueryOrderBys.SqlBuilder(), q.QueryPage.SqlBuilder()).Find(list).Error
}

func (q *Query[T]) JoinFind(list any, joins, selectFields func(db *gorm.DB) *gorm.DB) error {
	return q.db.Scopes(joins, selectFields, q.QueryOrderBys.SqlBuilder()).Find(list).Error
}

func (q *Query[T]) JoinList(list any, total *int64, joins, selectFields func(db *gorm.DB) *gorm.DB) error {
	if err := q.db.Model(q.M).Scopes(joins).Count(total).Error; err != nil {
		return err
	}
	return q.db.Scopes(joins, selectFields, q.QueryOrderBys.SqlBuilder(), q.QueryPage.SqlBuilder()).Find(list).Error
}

func (q *Query[T]) GORM() *gorm.DB {
	return q.db
}

// Create 创建
func (q *Query[T]) Create() error {
	var value = q.M
	return q.db.Create(&value).Error
}

func (q *Query[T]) CreateGetId() (int64, error) {
	var value = q.M
	err := q.db.Create(&value).Error
	return value.GetId(), err
}

// Save 保存
func (q *Query[T]) Save() error {
	var value = q.M
	return q.db.Save(&value).Error
}

func (q *Query[T]) Update() error {
	var value = q.M
	return q.db.Updates(&value).Error
}

func (q *Query[T]) Delete() error {
	var value = q.M
	return q.db.Delete(&value).Error
}
