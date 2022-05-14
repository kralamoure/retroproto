// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type SubwayPrismUse struct{}

func (m SubwayPrismUse) MessageId() retroproto.MsgCliId {
	return retroproto.SubwayPrismUse
}

func (m SubwayPrismUse) MessageName() string {
	return "SubwayPrismUse"
}

func (m SubwayPrismUse) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayPrismUse) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
