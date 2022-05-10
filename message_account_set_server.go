package retroproto

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retropb/gen/go/retropb"
)

type AccountSetServer struct {
	*retropb.AccountSetServer
}

func NewAccountSetServer(extra string) (AccountSetServer, error) {
	var m retropb.AccountSetServer

	id, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return AccountSetServer{}, fmt.Errorf("could not parse int: %w", err)
	}
	m.Id = int32(id)

	return AccountSetServer{AccountSetServer: &m}, nil
}

func (m AccountSetServer) MessageId() string {
	return "AX"
}

func (m AccountSetServer) String() string {
	return fmt.Sprintf("%d", m.GetId())
}
