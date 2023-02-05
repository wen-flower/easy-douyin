package msql

func Desc(column string) string {
	return column + " DESC"
}

func Asc(column string) string {
	return column + " Asc"
}
