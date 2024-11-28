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

	// 遍历消息的所有字段
	for _, e := range m.Elements {
		switch field := e.(type) {
		case *proto.NormalField:
			// 普通字段
			fmt.Printf("Normal Field: %s:%s\n", field.Name, field.Type)

			if field.Repeated == true {
				t.WLine("  public %s []%s;", t.GetOTypeByPbType(field.Type, false), field.Name)
			} else {
				//普通
				t.WLine("  public %s %s;", t.GetOTypeByPbType(field.Type, false), field.Name)
			}

		case *proto.MapField:
			t.WLine("  public Map<%s,%s> %s; ",
				t.GetOTypeByPbType(field.KeyType, true), t.GetOTypeByPbType(field.Type, false), field.Name)
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
				otype, _ := t.GetOTypeByPbType(field.Type)
				t.WLine("  public %s %s;", otype, field.Name)

			}
		}
	*/
	t.WLine("}")
}

// 通过pb类型转ts语言类型
func (t *ClassJavaTableInfo) GetOTypeByPbType(pbType string, isJava bool) string {
	otype := pbType
	switch pbType {
	case "string":
		return "String"
	case "int32":
		if isJava {
			return "Integer"
		}
		return "int"
	case "int64":
		return "long"
	case "bool":
		return "Boolean"
	}

	return otype
}

func (t *ClassJavaTableInfo) WLine(format string, a ...any) {
	aline := " " + fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"
}
func (t *ClassJavaTableInfo) GetDataContent() string {

	return t.file_content
}
