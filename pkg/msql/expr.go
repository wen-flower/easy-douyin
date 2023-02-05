package msql

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Inc 创建一个自增的 SQL 表达式
func Inc(column string) clause.Expr {
	return IncByStep(column, 1)
}

// Inc 创建一个自增 SQL 表达式
func IncByStep(column string, step int) clause.Expr {
	return gorm.Expr(column+" + ?", step)
}

// Dec 创建一个自减的 SQL 表达式
func Dec(column string) clause.Expr {
	return DecByStep(column, 1)
}

// Inc 创建一个自增 SQL 表达式
func DecByStep(column string, step int) clause.Expr {
	return gorm.Expr(column+" - ?", step)
}
