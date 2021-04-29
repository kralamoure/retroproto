// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountMiniClipInfo struct{}

func (m AccountMiniClipInfo) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountMiniClipInfo
}

func (m AccountMiniClipInfo) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountMiniClipInfo) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
