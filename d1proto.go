// Package d1proto is a library that implements the network protocol of Dofus 1,
// for communication between server and client. It defines the Id and structure
// of the protocol messages, and it provides functionality for their
// serialization and deserialization.
package d1proto

import "errors"

var ErrNotFound = errors.New("not found")
var ErrNotImplemented = errors.New("not implemented")
var ErrInvalidMsg = errors.New("invalid message")
