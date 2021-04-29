// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type InfosCompass struct{}

func (m InfosCompass) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.InfosCompass
}

func (m InfosCompass) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *InfosCompass) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
