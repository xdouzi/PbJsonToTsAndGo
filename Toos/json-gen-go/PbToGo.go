/*
*ExcelToGo
added by yh @ 2023/6/25 17:35
注意: go build excel_to_cx.go
*/
package main

import "PbJsonToTsAndGo/Toos/json-gen-go/to_go"

func main() {
	m := to_go.NewPbToGoMain()
	dirPath := "./Protol"
	m.Start(dirPath)
}
