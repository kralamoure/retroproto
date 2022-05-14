package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AksPong struct{}

func NewAksPong(extra string) (AksPong, error) {
	var m AksPong

	if err := m.Deserialize(extra); err != nil {
		return AksPong{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AksPong) MessageId() retroproto.MsgSvrId {
	return retroproto.AksPong
}

func (m AksPong) MessageName() string {
	return "AksPong"
}

func (m AksPong) Serialized() (string, error) {
	return "", nil
}

func (m *AksPong) Deserialize(extra string) error {
	return nil
}
