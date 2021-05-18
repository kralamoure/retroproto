package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type ExchangePutInShedFromCertificate struct {
	CertificateId int
}

func (m ExchangePutInShedFromCertificate) ProtocolId() retroproto.MsgCliId {
	return retroproto.ExchangePutInShedFromCertificate
}

func (m ExchangePutInShedFromCertificate) Serialized() (string, error) {
	return fmt.Sprint(m.CertificateId), nil
}

func (m *ExchangePutInShedFromCertificate) Deserialize(extra string) error {
	certificateId, err := strconv.Atoi(extra)
	if err != nil {
		return err
	}
	m.CertificateId = certificateId

	return nil
}
