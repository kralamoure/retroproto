package retroproto

import (
	"github.com/kralamoure/retropb/gen/go/retropb"
)

type AksHelloConnect struct {
	*retropb.AksHelloConnect
}

func NewAksHelloConnect(extra string) (AksHelloConnect, error) {
	var m retropb.AksHelloConnect

	m.Salt = extra

	return AksHelloConnect{AksHelloConnect: &m}, nil
}

func (m AksHelloConnect) MessageId() string {
	return "HC"
}

func (m AksHelloConnect) String() string {
	return m.GetSalt()
}
