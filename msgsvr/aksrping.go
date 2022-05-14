package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AksRPing struct {
	Value string
}

func NewAksRPing(extra string) (AksRPing, error) {
	var m AksRPing

	if err := m.Deserialize(extra); err != nil {
		return AksRPing{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AksRPing) MessageId() retroproto.MsgSvrId {
	return retroproto.AksRPing
}

func (m AksRPing) MessageName() string {
	return "AksRPing"
}

func (m AksRPing) Serialized() (string, error) {
	return m.Value, nil
}

func (m *AksRPing) Deserialize(extra string) error {
	m.Value = extra

	return nil
}
