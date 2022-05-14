package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type ExchangePutInShedFromCertificate struct {
	CertificateId int
}

func NewExchangePutInShedFromCertificate(extra string) (ExchangePutInShedFromCertificate, error) {
	var m ExchangePutInShedFromCertificate

	if err := m.Deserialize(extra); err != nil {
		return ExchangePutInShedFromCertificate{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangePutInShedFromCertificate) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangePutInShedFromCertificate
}

func (m ExchangePutInShedFromCertificate) MessageName() string {
	return "ExchangePutInShedFromCertificate"
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
