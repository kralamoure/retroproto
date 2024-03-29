package msgsvr

import (
	"fmt"
	"strings"
	"time"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/typ"
)

type AccountServersListSuccess struct {
	Subscription      time.Time
	ServersCharacters []typ.AccountServersListServerCharacters
}

func NewAccountServersListSuccess(extra string) (AccountServersListSuccess, error) {
	var m AccountServersListSuccess

	if err := m.Deserialize(extra); err != nil {
		return AccountServersListSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountServersListSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountServersListSuccess
}

func (m AccountServersListSuccess) MessageName() string {
	return "AccountServersListSuccess"
}

func (m AccountServersListSuccess) Serialized() (string, error) {
	sb := &strings.Builder{}

	serverCharactersSli := make([]string, len(m.ServersCharacters))
	for i, v := range m.ServersCharacters {
		serverCharacterStr, err := v.Serialized()
		if err != nil {
			return "", err
		}
		serverCharactersSli[i] = serverCharacterStr
	}

	serversCharacters := strings.Join(serverCharactersSli, "|")

	subscription := m.Subscription.Sub(time.Now()).Milliseconds()
	if subscription < 0 {
		subscription = 0
	}
	sb.WriteString(fmt.Sprintf("%d", subscription))

	if len(serverCharactersSli) > 0 {
		sb.WriteString(fmt.Sprintf("|%s", serversCharacters))
	}

	return sb.String(), nil
}

func (m *AccountServersListSuccess) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")

	if sli[0] != "" {
		subscriptionDur, err := time.ParseDuration(fmt.Sprintf("%sms", sli[0]))
		if err != nil {
			return err
		}
		m.Subscription = time.Now().Add(subscriptionDur)
	}

	if len(sli) < 2 {
		return nil
	}

	characters := sli[1:]

	m.ServersCharacters = make([]typ.AccountServersListServerCharacters, len(characters))
	for i, v := range characters {
		character := &typ.AccountServersListServerCharacters{}
		err := character.Deserialize(v)
		if err != nil {
			return err
		}
		m.ServersCharacters[i] = *character
	}

	return nil
}
