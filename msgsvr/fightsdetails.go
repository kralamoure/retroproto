// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type FightsDetails struct{}

func (m FightsDetails) ProtocolId() d1proto.MsgSvrId {
	return d1proto.FightsDetails
}

func (m FightsDetails) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *FightsDetails) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
