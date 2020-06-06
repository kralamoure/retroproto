// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangePutInMountParkFromShed struct{}

func (m ExchangePutInMountParkFromShed) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangePutInMountParkFromShed
}

func (m ExchangePutInMountParkFromShed) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangePutInMountParkFromShed) Deserialize(extra string) error {
	return nil
}
