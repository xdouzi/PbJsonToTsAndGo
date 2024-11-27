package to_java

import (
	"fmt"
	"github.com/emicklei/proto"
)

type EnumJavaTableInfo struct {
	p            *proto.Enum
	file_content string
}

func NewEnumJavaTableInfo(p *proto.Enum) *EnumJavaTableInfo {
	t := &EnumJavaTableInfo{}
	t.Init(p)
	return t
}
func (t *EnumJavaTableInfo) Init(p *proto.Enum) {
	t.p = p
	// 消息模块
	type CmdModule int32

	if p.Comment != nil && p.Comment.Lines != nil {
		t.WLine("//%s", p.Comment.Lines)
	}
	t.WLine("type %s int32", p.Name)

	t.WLine("const (")
	//fmt.Println("Enum Name:", p.Name) // 获取枚举名称
	for _, value := range p.Elements {
		if e, ok := value.(*proto.EnumField); ok {
			//fmt.Printf("Enum Value: Name=%s, Value=%d\n", e.Name, e.Integer)
			//CmdModule_IdleModule CmdModule = 0
			t.WLine(" %s %s = %d", e.Name, p.Name, e.Integer)

		}
	}

	t.WLine(" )")
}

func (t *EnumJavaTableInfo) WLine(format string, a ...any) {
	aline := " " + fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"
}
func (t *EnumJavaTableInfo) GetDataContent() string {

	return t.file_content
}
