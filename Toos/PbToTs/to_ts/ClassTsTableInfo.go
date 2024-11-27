package to_ts

import (
	"fmt"
	"github.com/emicklei/proto"
)

type ClassTsTableInfo struct {
	m            *proto.Message
	file_content string
}

func NewClassTsTableInfo() *ClassTsTableInfo {
	t := &ClassTsTableInfo{}
	return t
}

func (t *ClassTsTableInfo) DoData(m *proto.Message) {
	t.WLine("export class %s{", m.Name)
	for _, e := range m.Elements {
		if field, ok := e.(*proto.NormalField); ok {
			fmt.Printf("  Field: %s:%s \n", field.Name, field.Type)
			t.WLine(" public %s:%s;", field.Name, t.GetOTypeByPbType(field.Type))
		}
	}
	t.WLine("}")
}

// 通过pb类型转ts语言类型
func (t *ClassTsTableInfo) GetOTypeByPbType(pbType string) string {
	otype := ""
	switch pbType {
	case "string":
		return "string"
	case "int32":
		return "number"
	case "int64":
		return "number"
	case "bool":
		return "boolean"
	}

	return otype
}

func (t *ClassTsTableInfo) WLine(format string, a ...any) {
	aline := fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"
}
func (t *ClassTsTableInfo) GetDataContent() string {

	return t.file_content
}
