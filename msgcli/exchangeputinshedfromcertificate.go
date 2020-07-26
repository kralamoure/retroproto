// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangePutInShedFromCertificate struct{}

func (m ExchangePutInShedFromCertificate) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangePutInShedFromCertificate
}

func (m ExchangePutInShedFromCertificate) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangePutInShedFromCertificate) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
