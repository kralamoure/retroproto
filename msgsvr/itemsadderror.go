package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ItemsAddError struct {
	Reason rune
}

func (m ItemsAddError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ItemsAddError
}

func (m ItemsAddError) Serialized() (string, error) {
	return string(m.Reason), nil
}

func (m *ItemsAddError) Deserialize(extra string) error {
	if len(extra) == 0 {
		return retroproto.ErrInvalidMsg
	}

	for _, v := range extra {
		m.Reason = v
		break
	}

	return nil
}
