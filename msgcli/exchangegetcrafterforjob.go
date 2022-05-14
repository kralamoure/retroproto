// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeGetCrafterForJob struct{}

func NewExchangeGetCrafterForJob(extra string) (ExchangeGetCrafterForJob, error) {
	var m ExchangeGetCrafterForJob

	if err := m.Deserialize(extra); err != nil {
		return ExchangeGetCrafterForJob{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeGetCrafterForJob) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeGetCrafterForJob
}

func (m ExchangeGetCrafterForJob) MessageName() string {
	return "ExchangeGetCrafterForJob"
}

func (m ExchangeGetCrafterForJob) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeGetCrafterForJob) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
