package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type ExchangePutInCertificateFromShed struct {
	MountId int
}

func (m ExchangePutInCertificateFromShed) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangePutInCertificateFromShed
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
