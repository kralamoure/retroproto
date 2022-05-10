package retroproto

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retropb/gen/go/retropb"
)

type AccountServersListServerCharacters struct {
	*retropb.AccountServersListServerCharacters
}

func NewAccountServersListServerCharacters(extra string) (AccountServersListServerCharacters, error) {
	var m retropb.AccountServersListServerCharacters

	sli := strings.Split(extra, ",")
	if len(sli) < 2 {
		return AccountServersListServerCharacters{}, errInvalidMessage
	}

	if sli[0] != "" {
		id, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return AccountServersListServerCharacters{}, fmt.Errorf("could not parse int: %w", err)
		}
		m.Id = int32(id)
	}

	if sli[1] != "" {
		qty, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return AccountServersListServerCharacters{}, fmt.Errorf("could not parse int: %w", err)
		}
		m.Quantity = int32(qty)
	}

	return AccountServersListServerCharacters{AccountServersListServerCharacters: &m}, nil
}

func (m AccountServersListServerCharacters) MessageId() string {
	return ""
}

func (m AccountServersListServerCharacters) String() string {
	return fmt.Sprintf("%d,%d", m.GetId(), m.GetQuantity())
}
