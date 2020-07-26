// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangePutInCertificateFromShed struct{}

func (m ExchangePutInCertificateFromShed) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangePutInCertificateFromShed
}

func (m ExchangePutInCertificateFromShed) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangePutInCertificateFromShed) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
