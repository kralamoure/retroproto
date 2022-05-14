// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type InfosCompass struct{}

func NewInfosCompass(extra string) (InfosCompass, error) {
	var m InfosCompass

	if err := m.Deserialize(extra); err != nil {
		return InfosCompass{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m InfosCompass) MessageId() retroproto.MsgSvrId {
	return retroproto.InfosCompass
}

func (m InfosCompass) MessageName() string {
	return "InfosCompass"
}

func (m InfosCompass) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *InfosCompass) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
