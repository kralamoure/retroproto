// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AksPing struct{}

func NewAksPing(extra string) (AksPing, error) {
	var m AksPing

	if err := m.Deserialize(extra); err != nil {
		return AksPing{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AksPing) MessageId() retroproto.MsgCliId {
	return retroproto.AksPing
}

func (m AksPing) MessageName() string {
	return "AksPing"
}

func (m AksPing) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AksPing) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
