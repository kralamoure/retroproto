package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type ExchangePutInShedFromCertificate struct {
	CertificateId int
}

func (m ExchangePutInShedFromCertificate) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangePutInShedFromCertificate
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
