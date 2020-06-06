// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type InfosQuantity struct{}

func (m InfosQuantity) ProtocolId() d1proto.MsgSvrId {
	return d1proto.InfosQuantity
}

func (m InfosQuantity) Serialized() (string, error) {
	return "", nil
}

func (m *InfosQuantity) Deserialize(extra string) error {
	return nil
}
