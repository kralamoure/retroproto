// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountLeave struct{}

func NewMountLeave(extra string) (MountLeave, error) {
	var m MountLeave

	if err := m.Deserialize(extra); err != nil {
		return MountLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.MountLeave
}

func (m MountLeave) MessageName() string {
	return "MountLeave"
}

func (m MountLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *MountLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
