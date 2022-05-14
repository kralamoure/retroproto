package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AksHelloConnect struct {
	Salt string
}

func NewAksHelloConnect(extra string) (AksHelloConnect, error) {
	var m AksHelloConnect

	if err := m.Deserialize(extra); err != nil {
		return AksHelloConnect{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AksHelloConnect) MessageId() retroproto.MsgSvrId {
	return retroproto.AksHelloConnect
}

func (m AksHelloConnect) MessageName() string {
	return "AksHelloConnect"
}

func (m AksHelloConnect) Serialized() (string, error) {
	return m.Salt, nil
}

func (m *AksHelloConnect) Deserialize(extra string) error {
	m.Salt = extra

	return nil
}
