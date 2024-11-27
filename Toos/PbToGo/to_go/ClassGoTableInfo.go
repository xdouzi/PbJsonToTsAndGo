package to_go

import (
	"fmt"
	"github.com/emicklei/proto"
)

type ClassGoTableInfo struct {
	m            *proto.Message
	file_content string
}

func NewClassGoTableInfo() *ClassGoTableInfo {
	t := &ClassGoTableInfo{}
	return t
}

func (t *ClassGoTableInfo) DoData(m *proto.Message) {
	t.WLine("type %s struct {", m.Name)

	for _, e := range m.Elements {
		if field, ok := e.(*proto.NormalField); ok {
			fmt.Printf("  Field: %s:%s \n", field.Name, field.Type)
			otype, isp := t.GetOTypeByPbType(field.Type)
			if isp == false {
				t.WLine(" %s   %s", field.Name, otype)
			} else {
				t.WLine(" %s   *%s", field.Name, otype)
			}

		}
	}
	t.WLine("}")
}

// 通过pb类型转ts语言类型
func (t *ClassGoTableInfo) GetOTypeByPbType(pbType string) (string, bool) {
	otype := pbType
	switch pbType {
	case "string":
		return "string", false
	case "int32":
		return "int32", false
	case "int64":
		return "int64", false
	case "bool":
		return "bool", false
	}

	return otype, true
}

func (t *ClassGoTableInfo) WLine(format string, a ...any) {
	aline := " " + fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"
}
func (t *ClassGoTableInfo) GetDataContent() string {

	return t.file_content
}
