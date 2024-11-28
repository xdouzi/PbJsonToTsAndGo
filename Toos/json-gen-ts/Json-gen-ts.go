/*
*ExcelToGo
added by yh @ 2023/6/25 17:35
注意: go build excel_to_cx.go
*/
package main

import (
	"PbJsonToTsAndGo/Toos/json-gen-ts/to_ts"
)

func main() {
	m := to_ts.NewPbToTsMain()
	dirPath := "./Protol"
	m.Start(dirPath)
}
