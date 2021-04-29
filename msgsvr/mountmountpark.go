// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type MountMountPark struct{}

func (m MountMountPark) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.MountMountPark
}

func (m MountMountPark) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *MountMountPark) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
