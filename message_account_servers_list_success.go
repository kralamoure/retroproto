package retroproto

import (
	"fmt"
	"strings"
	"time"

	"github.com/kralamoure/retropb/gen/go/retropb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountServersListSuccess struct {
	*retropb.AccountServersListSuccess
}

func NewAccountServersListSuccess(extra string) (AccountServersListSuccess, error) {
	var m retropb.AccountServersListSuccess

	sli := strings.Split(extra, "|")

	if sli[0] != "" {
		subscriptionDur, err := time.ParseDuration(fmt.Sprintf("%sms", sli[0]))
		if err != nil {
			return AccountServersListSuccess{}, err
		}
		m.Subscription = timestamppb.New(time.Now().Add(subscriptionDur))
	}

	if len(sli) < 2 {
		return AccountServersListSuccess{}, errInvalidMessage
	}

	charactersSli := sli[1:]

	serverCharacters := make([]*retropb.AccountServersListServerCharacters, len(charactersSli))
	for i, v := range charactersSli {
		characters, err := NewAccountServersListServerCharacters(v)
		if err != nil {
			return AccountServersListSuccess{}, fmt.Errorf("could not create AccountServersListServerCharacters: %w", err)
		}
		serverCharacters[i] = characters.AccountServersListServerCharacters
	}
	m.ServerCharacters = serverCharacters

	return AccountServersListSuccess{AccountServersListSuccess: &m}, nil
}

func (m AccountServersListSuccess) MessageId() string {
	return "AxK"
}

func (m AccountServersListSuccess) String() string {
	sb := &strings.Builder{}

	serverCharactersSli := make([]string, len(m.GetServerCharacters()))

	for i := 0; i < len(m.GetServerCharacters()); i++ {
		serverCharactersSli[i] = m.GetServerCharacters()[i].String()
	}

	serversCharacters := strings.Join(serverCharactersSli, "|")

	subscription := m.GetSubscription().AsTime().Sub(time.Now()).Milliseconds()
	if subscription < 0 {
		subscription = 0
	}
	sb.WriteString(fmt.Sprintf("%d", subscription))

	if len(serverCharactersSli) > 0 {
		sb.WriteString(fmt.Sprintf("|%s", serversCharacters))
	}

	return sb.String()
}
