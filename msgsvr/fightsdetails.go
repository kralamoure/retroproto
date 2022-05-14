// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type FightsDetails struct{}

func (m FightsDetails) MessageId() retroproto.MsgSvrId {
	return retroproto.FightsDetails
}

func (m FightsDetails) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsDetails) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
