// Package d1proto is a library for communication between Dofus 1 client and server.
package d1proto

import "errors"

var ErrNotFound = errors.New("not found")
var ErrNotImplemented = errors.New("not implemented")
var ErrInvalidMsg = errors.New("invalid message")
