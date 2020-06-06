// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type InfosCompass struct{}

func (m InfosCompass) ProtocolId() d1proto.MsgSvrId {
	return d1proto.InfosCompass
}

func (m InfosCompass) Serialized() (string, error) {
	return "", nil
}

func (m *InfosCompass) Deserialize(extra string) error {
	return nil
}
