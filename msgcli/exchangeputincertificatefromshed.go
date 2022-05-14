package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type ExchangePutInCertificateFromShed struct {
	MountId int
}

func (m ExchangePutInCertificateFromShed) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangePutInCertificateFromShed
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
