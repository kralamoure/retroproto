package msgsvr

import (
	"github.com/kralamoure/d1proto"
	"github.com/kralamoure/d1proto/typ"
)

type AccountRestrictions struct {
	Restrictions typ.CommonRestrictions
}

func (m AccountRestrictions) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountRestrictions
}

func (m AccountRestrictions) Serialized() (string, error) {
	return m.Restrictions.Serialized()
}

func (m *AccountRestrictions) Deserialize(extra string) error {
	return m.Restrictions.Deserialize(extra)
}
