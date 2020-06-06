// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ConquestPrismInfosClosing struct{}

func (m ConquestPrismInfosClosing) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ConquestPrismInfosClosing
}

func (m ConquestPrismInfosClosing) Serialized() (string, error) {
	return "", nil
}

func (m *ConquestPrismInfosClosing) Deserialize(extra string) error {
	return nil
}
