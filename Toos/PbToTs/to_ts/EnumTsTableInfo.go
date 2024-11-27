package to_ts

import (
	"fmt"
	"github.com/emicklei/proto"
)

type EnumTsTableInfo struct {
	p            *proto.Enum
	file_content string
}

func NewEnumTsTableInfo(p *proto.Enum) *EnumTsTableInfo {
	t := &EnumTsTableInfo{}
	t.Init(p)
	return t
}
func (t *EnumTsTableInfo) Init(p *proto.Enum) {
	t.p = p
	t.WLine("export enum  %s {", p.Name)
	//fmt.Println("Enum Name:", p.Name) // 获取枚举名称
	for _, value := range p.Elements {
		if e, ok := value.(*proto.EnumField); ok {
			fmt.Printf("Enum Value: Name=%s, Value=%d\n", e.Name, e.Integer)
			t.WLine(" %s = %s,", e.Name, e.Integer)
		}
	}

	t.WLine("}")
}

func (t *EnumTsTableInfo) WLine(format string, a ...any) {
	aline := " " + fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"
}
func (t *EnumTsTableInfo) GetDataContent() string {

	return t.file_content
}
