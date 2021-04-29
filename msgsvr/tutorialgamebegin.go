// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type TutorialGameBegin struct{}

func (m TutorialGameBegin) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.TutorialGameBegin
}

func (m TutorialGameBegin) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *TutorialGameBegin) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
