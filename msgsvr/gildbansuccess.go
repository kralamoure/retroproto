// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GildBanSuccess struct{}

func (m GildBanSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GildBanSuccess
}

func (m GildBanSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *GildBanSuccess) Deserialize(extra string) error {
	return nil
}
