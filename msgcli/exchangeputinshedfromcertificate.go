package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type ExchangePutInShedFromCertificate struct {
	CertificateId int
}

func (m ExchangePutInShedFromCertificate) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangePutInShedFromCertificate
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
