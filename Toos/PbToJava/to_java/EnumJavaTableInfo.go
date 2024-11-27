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

	if p.Comment != nil && p.Comment.Lines != nil {
		t.WLine("//%s", p.Comment.Lines)
	}
	t.WLine("enum %s {", p.Name)

	//fmt.Println("Enum Name:", p.Name) // 获取枚举名称
	for _, value := range p.Elements {
		if e, ok := value.(*proto.EnumField); ok {
			//fmt.Printf("Enum Value: Name=%s, Value=%d\n", e.Name, e.Integer)
			//CmdModule_IdleModule CmdModule = 0
			t.WLine(" %s(%d),", e.Name, e.Integer)

		}
	}
	t.WLine("	;")

	t.WLine("")
	t.WLine("private final int code;")
	t.WLine("// 构造器")
	t.WLine("%s(int code) {", p.Name)
	t.WLine("	this.code = code;")
	t.WLine("}")

	t.WLine("// 获取枚举值的代码")
	t.WLine("public int getCode() {")
	t.WLine("	return code;")
	t.WLine("}")

	t.WLine(" }")
}

func (t *EnumJavaTableInfo) WLine(format string, a ...any) {
	aline := " " + fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"
}
func (t *EnumJavaTableInfo) GetDataContent() string {

	return t.file_content
}
