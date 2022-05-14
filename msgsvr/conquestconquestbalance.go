// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestConquestBalance struct{}

func NewConquestConquestBalance(extra string) (ConquestConquestBalance, error) {
	var m ConquestConquestBalance

	if err := m.Deserialize(extra); err != nil {
		return ConquestConquestBalance{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestConquestBalance) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestConquestBalance
}

func (m ConquestConquestBalance) MessageName() string {
	return "ConquestConquestBalance"
}

func (m ConquestConquestBalance) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestConquestBalance) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
