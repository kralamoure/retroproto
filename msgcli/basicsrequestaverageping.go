// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type BasicsRequestAveragePing struct{}

func (m BasicsRequestAveragePing) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsRequestAveragePing
}

func (m BasicsRequestAveragePing) MessageName() string {
	return "BasicsRequestAveragePing"
}

func (m BasicsRequestAveragePing) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsRequestAveragePing) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
