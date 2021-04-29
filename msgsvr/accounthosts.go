package msgsvr

import (
	"strings"

	"github.com/kralamoure/d1encoding"
	"github.com/kralamoure/d1encoding/typ"
)

type AccountHosts struct {
	Value []typ.AccountHostsHost
}

func (m AccountHosts) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountHosts
}

func (m AccountHosts) Serialized() (string, error) {
	hosts := make([]string, len(m.Value))
	for i, v := range m.Value {
		hostStr, err := v.Serialized()
		if err != nil {
			return "", err
		}
		hosts[i] = hostStr
	}

	return strings.Join(hosts, "|"), nil
}

func (m *AccountHosts) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")

	m.Value = make([]typ.AccountHostsHost, len(sli))
	for i, v := range sli {
		host := &typ.AccountHostsHost{}
		err := host.Deserialize(v)
		if err != nil {
			return err
		}
		m.Value[i] = *host
	}

	return nil
}
