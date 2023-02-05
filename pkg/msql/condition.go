package msql

// Eq 创建一个等值比较条件
func Eq(column string) string {
	return column + " = ?"
}

// Gt 创建一个大于条件
func Gt(column string) string {
	return column + " > ?"
}

// In 创建一个 IN 条件
func In(column string) string {
	return column + " IN ?"
}
