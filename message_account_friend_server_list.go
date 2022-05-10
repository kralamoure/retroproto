package retroproto

import (
	"fmt"
	"strings"

	"github.com/kralamoure/retropb/gen/go/retropb"
)

type AccountFriendServerList struct {
	*retropb.AccountFriendServerList
}

func NewAccountFriendServerList(extra string) (AccountFriendServerList, error) {
	if extra == "" {
		return AccountFriendServerList{}, errInvalidMessage
	}

	var m retropb.AccountFriendServerList

	if extra == "null" {
		return AccountFriendServerList{}, nil
	}

	charactersSli := strings.Split(extra, ";")

	serverCharacters := make([]*retropb.AccountServersListServerCharacters, len(charactersSli))
	for i, v := range charactersSli {
		characters, err := NewAccountServersListServerCharacters(v)
		if err != nil {
			return AccountFriendServerList{}, fmt.Errorf("could not create AccountServersListServerCharacters: %w", err)
		}
		serverCharacters[i] = characters.AccountServersListServerCharacters
	}
	m.ServerCharacters = serverCharacters

	return AccountFriendServerList{AccountFriendServerList: &m}, nil
}

func (m AccountFriendServerList) MessageId() string {
	return "AF"
}

func (m AccountFriendServerList) String() string {
	if len(m.GetServerCharacters()) == 0 {
		return "null"
	}

	serverCharactersSli := make([]string, len(m.GetServerCharacters()))

	for i := 0; i < len(m.GetServerCharacters()); i++ {
		serverCharactersSli[i] = m.GetServerCharacters()[i].String()
	}

	return strings.Join(serverCharactersSli, ";")
}
