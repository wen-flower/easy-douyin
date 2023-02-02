package msql

// Eq 创建一个等值比较条件
func Eq(column string) string {
	return column + " = ?"
}
