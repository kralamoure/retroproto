// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type FightsDetails struct{}

func (m FightsDetails) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.FightsDetails
}

func (m FightsDetails) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FightsDetails) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
