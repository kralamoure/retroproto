// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type BasicsAveragePing struct{}

func (m BasicsAveragePing) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsAveragePing
}

func (m BasicsAveragePing) MessageName() string {
	return "BasicsAveragePing"
}

func (m BasicsAveragePing) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAveragePing) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
