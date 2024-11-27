package to_ts

import "github.com/emicklei/proto"

type EnumTsTableInfo struct {
}

func NewEnumTsTableInfo(p *proto.Enum) *EnumTsTableInfo {
	t := &EnumTsTableInfo{}
	return t
}
