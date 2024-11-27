package to_java

import (
	"fmt"
	"github.com/emicklei/proto"
)

type ClassJavaTableInfo struct {
	m            *proto.Message
	file_content string
}

func NewClassJavaTableInfo() *ClassJavaTableInfo {
	t := &ClassJavaTableInfo{}
	return t
}

func (t *ClassJavaTableInfo) DoData(m *proto.Message) {
	t.WLine("class %s {", m.Name)

	for _, e := range m.Elements {
		if field, ok := e.(*proto.NormalField); ok {
			fmt.Printf("  Field: %s:%s \n", field.Name, field.Type)
			otype, _ := t.GetOTypeByPbType(field.Type)
			t.WLine("  public %s %s;", otype, field.Name)

		}
	}
	t.WLine("}")
}

// 通过pb类型转ts语言类型
func (t *ClassJavaTableInfo) GetOTypeByPbType(pbType string) (string, bool) {
	otype := pbType
	switch pbType {
	case "string":
		return "String", false
	case "int32":
		return "int", false
	case "int64":
		return "long", false
	case "bool":
		return "Boolean", false
	}

	return otype, true
}

func (t *ClassJavaTableInfo) WLine(format string, a ...any) {
	aline := " " + fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"
}
func (t *ClassJavaTableInfo) GetDataContent() string {

	return t.file_content
}
