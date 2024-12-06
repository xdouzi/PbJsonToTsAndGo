package to_ts

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type PbToTsMain struct {
	TsItemMap    []*PbToTsItem
	file_content string
}

func NewPbToTsMain() *PbToTsMain {
	return &PbToTsMain{}
}
func (t *PbToTsMain) CheckToTime() bool {
	return false
}

func (t *PbToTsMain) Hang() {

}

func (t *PbToTsMain) DeleteFiles(output string, s string) {

}

func (t *PbToTsMain) Start(dirPath string) {
	// 检查文件夹是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 如果不存在则创建文件夹
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("Failed to create directory:", err)
			return
		}
		fmt.Println("Directory created successfully!")
	}

	files, err := os.Open(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer files.Close()

	fileNames, err := files.Readdirnames(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	//go build excel_to_cx.go
	/*
		fmt.Printf("----------*.xlsx生成规则--------------------\n")
		fmt.Printf("--第一行 字段属性为 默认空:客户端和服务端使用;c:只有客户端用;s:只有服务端使用;\n")
		fmt.Printf("--第二行 字段属性名字 如果为空字符串这个列将不生成配置\n")
		fmt.Printf("--第三行 字段类型 如果为空 将默认为string,其它类型有 string,float\n")
		fmt.Printf("--第四行 字段介绍名称\n")
		fmt.Printf("--第五行 字段属性使用介绍\n")
		fmt.Printf("--第六行 配置第一行数据开始\n")
	*/
	fmt.Printf("added by yh @ 2024/11/27 17:35 408309839@qq.com \n")
	fmt.Printf("\n")
	if t.CheckToTime() {
		fmt.Printf("--少年有报错联系管理员...\n")
		t.Hang()
		return
	}
	go_output := "./Protocal" //"./bin/cfg_go"
	//json_output := "./cfg_json" //"./bin/cfg_json"
	t.DeleteFiles(go_output, ".ts")
	//m.DeleteFiles(json_output, ".json")

	for _, fileName := range fileNames {
		if filepath.Ext(fileName) == ".proto" {
			filePath := filepath.Join(dirPath, fileName)
			item := NewPbToTsItem()
			item.OpenProtoFile(fileName, filePath, go_output)
			t.TsItemMap = append(t.TsItemMap, item)

			//js := to_json.NewExcelToJson()
			//js.OpenExcelFile(fileName, filePath, json_output)
		}
	}

	t.WLine("export namespace %s {", "pb")

	for _, table := range t.TsItemMap {
		t.WLine(table.GetItemContent())
	}
	t.WLine("}")

	t.SaveFile(go_output)
	fmt.Printf("\n")
	fmt.Printf("----------生成所有Ts配置完成---------------------\n")
}
func (t *PbToTsMain) WLine(format string, a ...any) {
	aline := fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"

}
func (t *PbToTsMain) SaveFile(output string) {
	//dirPath := "./Cx_output"
	dirPath := output
	// 检查文件夹是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 如果不存在则创建文件夹
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("Failed to create directory:", err)
			return
		}
		fmt.Println("Directory created successfully!")
	}

	outputFile := fmt.Sprintf("%s/%s.ts", output, "pb")
	// 将配置文件写入文件中
	err := ioutil.WriteFile(outputFile, []byte(t.file_content), 0644)
	if err != nil {
		fmt.Println(fmt.Sprintf("失败 配置文件%s err=%s ", outputFile, err))
		return
	} else {
		fmt.Println("配置文件已生成", outputFile)
	}
	t.file_content = ""
}
