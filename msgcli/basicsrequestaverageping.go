// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsRequestAveragePing struct{}

func NewBasicsRequestAveragePing(extra string) (BasicsRequestAveragePing, error) {
	var m BasicsRequestAveragePing

	if err := m.Deserialize(extra); err != nil {
		return BasicsRequestAveragePing{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsRequestAveragePing) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsRequestAveragePing
}

func (m BasicsRequestAveragePing) MessageName() string {
	return "BasicsRequestAveragePing"
}

func (m BasicsRequestAveragePing) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsRequestAveragePing) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
