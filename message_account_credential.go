package retroproto

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retropb/gen/go/retropb"
)

type AccountCredential struct {
	*retropb.AccountCredential
}

func NewAccountCredential(extra string) (AccountCredential, error) {
	var m retropb.AccountCredential

	sli := strings.Split(extra, "\n")

	if len(sli) < 2 {
		return AccountCredential{}, errInvalidMessage
	}

	m.Username = sli[0]

	if len(sli[1]) < 3 {
		return AccountCredential{}, errInvalidMessage
	}

	if sli[1][0] != '#' {
		return AccountCredential{}, errInvalidMessage
	}

	cryptoMethod := sli[1][1]
	cryptoMethodN, err := strconv.ParseInt(string(cryptoMethod), 10, 32)
	if err != nil {
		return AccountCredential{}, fmt.Errorf("could not parse int: %w", err)
	}
	m.CryptoMethod = int32(cryptoMethodN)

	m.Hash = sli[1][2:]

	return AccountCredential{AccountCredential: &m}, nil
}

func (m AccountCredential) MessageId() string {
	return ""
}

func (m AccountCredential) String() string {
	return fmt.Sprintf("%s\n#%d%s", m.GetUsername(), m.GetCryptoMethod(), m.GetHash())
}
