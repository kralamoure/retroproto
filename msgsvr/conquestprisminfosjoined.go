// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ConquestPrismInfosJoined struct{}

func (m ConquestPrismInfosJoined) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismInfosJoined
}

func (m ConquestPrismInfosJoined) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismInfosJoined) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
