// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsAveragePing struct{}

func NewBasicsAveragePing(extra string) (BasicsAveragePing, error) {
	var m BasicsAveragePing

	if err := m.Deserialize(extra); err != nil {
		return BasicsAveragePing{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsAveragePing) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsAveragePing
}

func (m BasicsAveragePing) MessageName() string {
	return "BasicsAveragePing"
}

func (m BasicsAveragePing) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAveragePing) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
