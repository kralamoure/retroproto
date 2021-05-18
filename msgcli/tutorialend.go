// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type TutorialEnd struct{}

func (m TutorialEnd) ProtocolId() retroproto.MsgCliId {
	return retroproto.TutorialEnd
}

func (m TutorialEnd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *TutorialEnd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
