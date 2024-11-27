package to_ts

import (
	"fmt"
	"github.com/emicklei/proto"
	"io/ioutil"
	"os"
	"strings"
)

type PbToGo struct {
	m             *proto.Message
	output        string //输出生成文件路径 ./Cx_output
	file_content  string
	fileName      string
	Name          string
	EnumMap       []*EnumTsTableInfo
	SheetTableMap []*ClassTsTableInfo
	packageName   string
}

func NewPbToGo() *PbToGo {
	return &PbToGo{}
}

// 打开协议文件
func (t *PbToGo) OpenProtoFile(fileName string, path string, output string) {
	t.fileName = fileName
	t.output = output
	t.Name = strings.Replace(fileName, ".proto", "", -1)

	t.file_content = ""

	// 打开 .proto 文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening .proto file:", err)
		return
	}
	defer file.Close()

	// 解析 .proto 文件
	parser := proto.NewParser(file)
	definition, err := parser.Parse()
	if err != nil {
		fmt.Println("Error parsing .proto file:", err)
		return
	}

	//------------------------------------------------------------
	t.DoSheetTable(definition)
	//------------------------------------------------------------
	//整合处理
	t.DoAllIntegrate()
	t.SaveFile()
}
func (t *PbToGo) DoSheetTable(definition *proto.Proto) {
	// 遍历消息和字段
	proto.Walk(definition,
		proto.WithPackage(func(p *proto.Package) {
			t.packageName = p.Name
		}),
		proto.WithEnum(func(p *proto.Enum) {
			//处理美剧
			e := NewEnumTsTableInfo(p)
			t.EnumMap = append(t.EnumMap, e)
		}),
		proto.WithMessage(func(m *proto.Message) {
			t.m = m
			q := NewClassTsTableInfo()
			q.DoData(m)
			t.SheetTableMap = append(t.SheetTableMap, q)

			/*
				//fmt.Println("Message:", m.Name)
				for _, e := range m.Elements {
					if field, ok := e.(*proto.NormalField); ok {
						fmt.Printf("  Field: %s:%s \n", field.Name, field.Type)
						t.WLine(" public %s:%s;", field.Name, field.Type)

						t.SheetTableMap = append(t.SheetTableMap, q)
					}
				}
			*/

		}),
	)
}

// 所有整合
func (t *PbToGo) DoAllIntegrate() {
	t.WLine("/**")
	// 获取当前时间
	//currentTime := time.Now()
	// 格式化为 "年:月:日 00:00" 的格式
	//formattedTime := currentTime.Format("2006.01.02 15:04")
	t.WLine("由 %s.xlsx %s excel文件生成 ...", t.Name)
	t.WLine("author:yh ")
	t.WLine("*/")

	t.WLine("export namespace %s {", t.packageName)

	for _, table := range t.EnumMap {
		t.WLine(table.GetDataContent())
	}

	for _, table := range t.SheetTableMap {
		t.WLine(table.GetDataContent())
	}

	t.WLine("}")
}

func (t *PbToGo) SaveFile() {
	//dirPath := "./Cx_output"
	dirPath := t.output
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

	outputFile := fmt.Sprintf("%s/%s.ts", t.output, t.Name)
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
func (t *PbToGo) WLine(format string, a ...any) {
	aline := fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"

}
