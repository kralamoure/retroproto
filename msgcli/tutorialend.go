// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type TutorialEnd struct{}

func (m TutorialEnd) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.TutorialEnd
}

func (m TutorialEnd) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *TutorialEnd) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
