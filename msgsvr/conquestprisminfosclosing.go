// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestPrismInfosClosing struct{}

func (m ConquestPrismInfosClosing) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ConquestPrismInfosClosing
}

func (m ConquestPrismInfosClosing) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestPrismInfosClosing) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
