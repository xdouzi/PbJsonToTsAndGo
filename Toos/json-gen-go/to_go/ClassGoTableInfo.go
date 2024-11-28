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

	// 遍历消息的所有字段
	for _, e := range m.Elements {
		switch field := e.(type) {
		case *proto.NormalField:
			// 普通字段
			fmt.Printf("Normal Field: %s:%s\n", field.Name, field.Type)

			if field.Repeated == true {
				t.WLine(" %s   []%s", field.Name, t.GetOTypeByPbType(field.Type))
			} else {
				//普通
				t.WLine(" %s   %s", field.Name, t.GetOTypeByPbType(field.Type))
			}

		case *proto.MapField:
			t.WLine(" %s map[%s]%s ", field.Name, t.GetOTypeByPbType(field.KeyType), t.GetOTypeByPbType(field.Type))
		case *proto.Comment: //RepeatedField
			//注视的不处理
			//t.WLine(" public %s:%s;", field.Message(), t.GetOTypeByPbType(field.Message()))
		// Repeated 类型字段
		//fmt.Printf("Repeated Field: %s (Type: %s)\n", field.Name, field.Type)
		default:
			// 其他类型
			fmt.Printf("----未知类型 Other Field Type: %T\n", field)
		}
	}
	/*
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
	*/
	t.WLine("}")
}

// 通过pb类型转ts语言类型
func (t *ClassGoTableInfo) GetOTypeByPbType(pbType string) string {
	otype := pbType
	switch pbType {
	case "string":
		return "string"
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	case "bool":
		return "bool"
	}

	return "*" + otype
}

func (t *ClassGoTableInfo) WLine(format string, a ...any) {
	aline := " " + fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"
}
func (t *ClassGoTableInfo) GetDataContent() string {

	return t.file_content
}
