package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/typ"
)

type AccountRestrictions struct {
	Restrictions typ.CommonRestrictions
}

func NewAccountRestrictions(extra string) (AccountRestrictions, error) {
	var m AccountRestrictions

	if err := m.Deserialize(extra); err != nil {
		return AccountRestrictions{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountRestrictions) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountRestrictions
}

func (m AccountRestrictions) MessageName() string {
	return "AccountRestrictions"
}

func (m AccountRestrictions) Serialized() (string, error) {
	return m.Restrictions.Serialized()
}

func (m *AccountRestrictions) Deserialize(extra string) error {
	return m.Restrictions.Deserialize(extra)
}
