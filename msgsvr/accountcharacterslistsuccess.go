package msgsvr

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/typ"
)

type AccountCharactersListSuccess struct {
	Subscription    time.Time
	CharactersCount int
	Characters      []typ.AccountCharactersListCharacter
}

func (m AccountCharactersListSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharactersListSuccess
}

func (m AccountCharactersListSuccess) Serialized() (string, error) {
	sb := &strings.Builder{}

	charactersSli := make([]string, len(m.Characters))
	for i, v := range m.Characters {
		characterStr, err := v.Serialized()
		if err != nil {
			return "", err
		}
		charactersSli[i] = characterStr
	}

	characters := strings.Join(charactersSli, "|")

	subscription := m.Subscription.Sub(time.Now()).Milliseconds()
	if subscription < 0 {
		subscription = 0
	}
	sb.WriteString(fmt.Sprintf("%d|%d", subscription, m.CharactersCount))

	if len(charactersSli) > 0 {
		sb.WriteString(fmt.Sprintf("|%s", characters))
	}

	return sb.String(), nil
}

func (m *AccountCharactersListSuccess) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 2 {
		return retroproto.ErrInvalidMsg
	}

	if sli[0] != "" {
		subscriptionDur, err := time.ParseDuration(fmt.Sprintf("%sms", sli[0]))
		if err != nil {
			return err
		}
		m.Subscription = time.Now().Add(subscriptionDur)
	}

	if sli[1] != "" {
		charactersCount, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return err
		}
		m.CharactersCount = int(charactersCount)
	}

	if len(sli) < 3 {
		return nil
	}

	characters := sli[2:]

	m.Characters = make([]typ.AccountCharactersListCharacter, len(characters))
	for i, v := range characters {
		character := &typ.AccountCharactersListCharacter{}
		err := character.Deserialize(v)
		if err != nil {
			return err
		}
		m.Characters[i] = *character
	}

	return nil
}
