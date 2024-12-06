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

	// 遍历消息的所有字段
	for _, e := range m.Elements {
		switch field := e.(type) {
		case *proto.NormalField:
			// 普通字段
			fmt.Printf("Normal Field: %s:%s\n", field.Name, field.Type)
			if field.Repeated == true {
				//数组
				t.WLine(" public %s:%s[];", field.Name, t.GetOTypeByPbType(field.Type))
			} else {
				//普通
				t.WLine(" public %s:%s;", field.Name, t.GetOTypeByPbType(field.Type))
			}

		case *proto.MapField:
			t.WLine(" public %s: Map<%s, %s>;", field.Name, t.GetOTypeByPbType(field.KeyType), t.GetOTypeByPbType(field.Type))
		case *proto.Comment: //RepeatedField
			//注视了的不处理
			//t.WLine(" public %s:%s;", field.Message(), t.GetOTypeByPbType(field.Message()))
		// Repeated 类型字段
		//fmt.Printf("Repeated Field: %s (Type: %s)\n", field.Name, field.Type)
		default:
			// 其他类型
			fmt.Printf("Other Field Type: %T\n", field)
		}
	}
	/*
		for _, e := range m.Elements {
			log.Printf("---xxxxx--%s", e)
			if field, ok := e.(*proto.NormalField); ok {
				fmt.Printf("  Field: %s:%s \n", field.Name, field.Type)
				t.WLine(" public %s:%s;", field.Name, t.GetOTypeByPbType(field.Type))
			}
		}
	*/

	t.WLine("}")
}

// 通过pb类型转ts语言类型
func (t *ClassTsTableInfo) GetOTypeByPbType(pbType string) string {
	otype := pbType
	switch pbType {
	case "string":
		return "string"
	case "int32":
		return "number"
	case "int64":
		return "number"
	case "bool":
		return "boolean"
	case "double":
		return "number"
	case "bytes":
		return "Uint8Array"

	}

	return otype
}

func (t *ClassTsTableInfo) WLine(format string, a ...any) {
	aline := " " + fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"
}
func (t *ClassTsTableInfo) GetDataContent() string {

	return t.file_content
}
