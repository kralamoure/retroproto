package msgsvr

import (
	"fmt"
	"strings"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/typ"
)

type AccountHosts struct {
	Value []typ.AccountHostsHost
}

func NewAccountHosts(extra string) (AccountHosts, error) {
	var m AccountHosts

	if err := m.Deserialize(extra); err != nil {
		return AccountHosts{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountHosts) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountHosts
}

func (m AccountHosts) MessageName() string {
	return "AccountHosts"
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
