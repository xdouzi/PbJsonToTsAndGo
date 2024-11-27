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
			t.WLine(" public %s:%s;", field.Name, field.Type)
		}
	}
	t.WLine("}")
}
func (t *ClassTsTableInfo) WLine(format string, a ...any) {
	aline := fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"
}
func (t *ClassTsTableInfo) GetDataContent() string {

	return t.file_content
}
