// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type SubwayPrismUse struct{}

func (m SubwayPrismUse) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.SubwayPrismUse
}

func (m SubwayPrismUse) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *SubwayPrismUse) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
