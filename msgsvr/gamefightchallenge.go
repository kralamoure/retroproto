// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameFightChallenge struct{}

func (m GameFightChallenge) MessageId() retroproto.MsgSvrId {
	return retroproto.GameFightChallenge
}

func (m GameFightChallenge) MessageName() string {
	return "GameFightChallenge"
}

func (m GameFightChallenge) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFightChallenge) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
