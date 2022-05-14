// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type SubwayUse struct{}

func (m SubwayUse) MessageId() retroproto.MsgCliId {
	return retroproto.SubwayUse
}

func (m SubwayUse) MessageName() string {
	return "SubwayUse"
}

func (m SubwayUse) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayUse) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
