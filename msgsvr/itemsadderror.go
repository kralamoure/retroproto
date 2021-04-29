package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ItemsAddError struct {
	Reason rune
}

func (m ItemsAddError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsAddError
}

func (m ItemsAddError) Serialized() (string, error) {
	return string(m.Reason), nil
}

func (m *ItemsAddError) Deserialize(extra string) error {
	if len(extra) == 0 {
		return d1proto.ErrInvalidMsg
	}

	for _, v := range extra {
		m.Reason = v
		break
	}

	return nil
}
