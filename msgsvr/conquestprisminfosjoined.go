// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ConquestPrismInfosJoined struct{}

func (m ConquestPrismInfosJoined) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ConquestPrismInfosJoined
}

func (m ConquestPrismInfosJoined) Serialized() (string, error) {
	return "", nil
}

func (m *ConquestPrismInfosJoined) Deserialize(extra string) error {
	return nil
}
