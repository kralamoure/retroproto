// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestPrismFightAddPlayerAdd struct{}

func NewConquestPrismFightAddPlayerAdd(extra string) (ConquestPrismFightAddPlayerAdd, error) {
	var m ConquestPrismFightAddPlayerAdd

	if err := m.Deserialize(extra); err != nil {
		return ConquestPrismFightAddPlayerAdd{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestPrismFightAddPlayerAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismFightAddPlayerAdd
}

func (m ConquestPrismFightAddPlayerAdd) MessageName() string {
	return "ConquestPrismFightAddPlayerAdd"
}

func (m ConquestPrismFightAddPlayerAdd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismFightAddPlayerAdd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
