// Package retroproto is a library that implements the network protocol of Dofus Retro,
// for communication between server and client. It defines the id and structure
// of the protocol messages, and it provides functionality for their
// serialization and deserialization.
package retroproto

import "errors"

var ErrNotFound = errors.New("not found")
var ErrNotImplemented = errors.New("not implemented")
var ErrInvalidMsg = errors.New("invalid message")
