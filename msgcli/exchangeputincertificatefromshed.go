package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type ExchangePutInCertificateFromShed struct {
	MountId int
}

func (m ExchangePutInCertificateFromShed) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangePutInCertificateFromShed
}

func (m ExchangePutInCertificateFromShed) Serialized() (string, error) {
	return fmt.Sprint(m.MountId), nil
}

func (m *ExchangePutInCertificateFromShed) Deserialize(extra string) error {
	mountId, err := strconv.Atoi(extra)
	if err != nil {
		return err
	}
	m.MountId = mountId

	return nil
}
