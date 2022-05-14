package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AksServerMessage struct {
	// TODO
	Value string
}

func NewAksServerMessage(extra string) (AksServerMessage, error) {
	var m AksServerMessage

	if err := m.Deserialize(extra); err != nil {
		return AksServerMessage{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AksServerMessage) MessageId() retroproto.MsgSvrId {
	return retroproto.AksServerMessage
}

func (m AksServerMessage) MessageName() string {
	return "AksServerMessage"
}

func (m AksServerMessage) Serialized() (string, error) {
	return m.Value, nil
}

func (m *AksServerMessage) Deserialize(extra string) error {
	m.Value = extra

	return nil
}
