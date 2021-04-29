package msgsvr

import (
	"github.com/kralamoure/d1encoding"
	"github.com/kralamoure/d1encoding/typ"
)

type AccountRestrictions struct {
	Restrictions typ.CommonRestrictions
}

func (m AccountRestrictions) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountRestrictions
}

func (m AccountRestrictions) Serialized() (string, error) {
	return m.Restrictions.Serialized()
}

func (m *AccountRestrictions) Deserialize(extra string) error {
	return m.Restrictions.Deserialize(extra)
}
