// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type HousesBuyError struct{}

func (m HousesBuyError) MessageId() retroproto.MsgSvrId {
	return retroproto.HousesBuyError
}

func (m HousesBuyError) MessageName() string {
	return "HousesBuyError"
}

func (m HousesBuyError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesBuyError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
