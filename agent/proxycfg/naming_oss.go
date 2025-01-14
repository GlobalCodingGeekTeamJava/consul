//go:build !consulent
// +build !consulent

package proxycfg

import (
	"github.com/hashicorp/consul/agent/structs"
)

func UpstreamIDString(typ, dc, name string, _ *structs.EnterpriseMeta) string {
	ret := name

	if dc != "" {
		ret += "?dc=" + dc
	}

	if typ == "" || typ == structs.UpstreamDestTypeService {
		return ret
	}

	return typ + ":" + ret
}

func parseInnerUpstreamIDString(input string) (string, *structs.EnterpriseMeta) {
	return input, structs.DefaultEnterpriseMetaInDefaultPartition()
}

func (u UpstreamID) enterpriseIdentifierPrefix() string {
	return ""
}
