// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangePutInShedFromMountPark struct{}

func (m ExchangePutInShedFromMountPark) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangePutInShedFromMountPark
}

func (m ExchangePutInShedFromMountPark) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangePutInShedFromMountPark) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
