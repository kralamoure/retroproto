// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestPrismInfosJoined struct{}

func (m ConquestPrismInfosJoined) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ConquestPrismInfosJoined
}

func (m ConquestPrismInfosJoined) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestPrismInfosJoined) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
