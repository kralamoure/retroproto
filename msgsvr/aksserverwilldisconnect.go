package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AksServerWillDisconnect struct{}

func NewAksServerWillDisconnect(extra string) (AksServerWillDisconnect, error) {
	var m AksServerWillDisconnect

	if err := m.Deserialize(extra); err != nil {
		return AksServerWillDisconnect{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AksServerWillDisconnect) MessageId() retroproto.MsgSvrId {
	return retroproto.AksServerWillDisconnect
}

func (m AksServerWillDisconnect) MessageName() string {
	return "AksServerWillDisconnect"
}

func (m AksServerWillDisconnect) Serialized() (string, error) {
	return "", nil
}

func (m *AksServerWillDisconnect) Deserialize(extra string) error {
	return nil
}
