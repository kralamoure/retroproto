// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestPrismSurvived struct{}

func NewConquestPrismSurvived(extra string) (ConquestPrismSurvived, error) {
	var m ConquestPrismSurvived

	if err := m.Deserialize(extra); err != nil {
		return ConquestPrismSurvived{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestPrismSurvived) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismSurvived
}

func (m ConquestPrismSurvived) MessageName() string {
	return "ConquestPrismSurvived"
}

func (m ConquestPrismSurvived) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismSurvived) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
