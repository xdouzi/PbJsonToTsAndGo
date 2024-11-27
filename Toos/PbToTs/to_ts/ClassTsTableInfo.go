package to_ts

import "github.com/emicklei/proto"

type ClassTsTableInfo struct {
	field *proto.NormalField
}

func (t *ClassTsTableInfo) GetDataContent() string {

}

func NewClassTsTableInfo(field *proto.NormalField) *ClassTsTableInfo {
	t := &ClassTsTableInfo{}
	t.field = field
	return t
}
