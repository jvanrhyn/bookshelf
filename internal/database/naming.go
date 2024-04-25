package database

import (
	"github.com/jinzhu/inflection"
	"gorm.io/gorm/schema"
	"strings"
)

type PostgresNamingStrategy struct{}

func (p PostgresNamingStrategy) TableName(table string) string {
	return strings.ToLower(inflection.Singular(table))
}

func (p PostgresNamingStrategy) SchemaName(table string) string {
	return strings.ToLower(table)
}

func (p PostgresNamingStrategy) ColumnName(table, column string) string {
	return strings.ToLower(column)
}

func (p PostgresNamingStrategy) JoinTableName(joinTable string) string {
	return strings.ToLower(inflection.Singular(joinTable))
}

func (p PostgresNamingStrategy) RelationshipFKName(relationship schema.Relationship) string {
	return strings.ToLower(relationship.Name)
}

func (p PostgresNamingStrategy) CheckerName(table, column string) string {
	return strings.ToLower("chk_" + table + "_" + column)
}

func (p PostgresNamingStrategy) IndexName(table, column string) string {
	return strings.ToLower("idx_" + table + "_" + column)
}

func (p PostgresNamingStrategy) UniqueName(table, column string) string {
	return strings.ToLower("uq_" + table + "_" + column)
}
