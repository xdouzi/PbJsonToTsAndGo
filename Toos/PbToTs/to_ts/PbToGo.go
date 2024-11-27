package to_ts

import (
	"fmt"
	"github.com/emicklei/proto"
	"os"
)

type PbToGo struct {
	m *proto.Message
}

func NewPbToGo() *PbToGo {
	return &PbToGo{}
}

// 打开协议文件
func (t *PbToGo) OpenProtoFile(name string, path string, output string) {
	// 打开 .proto 文件
	file, err := os.Open("Net_Login.proto")
	if err != nil {
		fmt.Println("Error opening .proto file:", err)
		return
	}
	defer file.Close()

	// 解析 .proto 文件
	parser := proto.NewParser(file)
	definition, err := parser.Parse()
	if err != nil {
		fmt.Println("Error parsing .proto file:", err)
		return
	}

	// 遍历消息和字段
	proto.Walk(definition,
		proto.WithMessage(func(m *proto.Message) {
			fmt.Println("Message:", m.Name)
			t.m = m
			//t.Name = m.Name
			for _, e := range m.Elements {
				if field, ok := e.(*proto.NormalField); ok {
					fmt.Printf("  Field: %s (%s)\n", field.Name, field.Type)
				}
			}
		}),
	)
}
