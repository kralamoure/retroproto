// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountMiniClipInfo struct{}

func NewAccountMiniClipInfo(extra string) (AccountMiniClipInfo, error) {
	var m AccountMiniClipInfo

	if err := m.Deserialize(extra); err != nil {
		return AccountMiniClipInfo{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountMiniClipInfo) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountMiniClipInfo
}

func (m AccountMiniClipInfo) MessageName() string {
	return "AccountMiniClipInfo"
}

func (m AccountMiniClipInfo) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountMiniClipInfo) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
