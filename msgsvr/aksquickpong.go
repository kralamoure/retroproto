package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AksQuickPong struct{}

func NewAksQuickPong(extra string) (AksQuickPong, error) {
	var m AksQuickPong

	if err := m.Deserialize(extra); err != nil {
		return AksQuickPong{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AksQuickPong) MessageId() retroproto.MsgSvrId {
	return retroproto.AksQuickPong
}

func (m AksQuickPong) MessageName() string {
	return "AksQuickPong"
}

func (m AksQuickPong) Serialized() (string, error) {
	return "", nil
}

func (m *AksQuickPong) Deserialize(extra string) error {
	return nil
}
