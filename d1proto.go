// Package d1proto is a library for the network protocol of Dofus 1. It can serialize and deserialize messages for server and client communication.
package d1proto

import "errors"

var ErrNotFound = errors.New("not found")
var ErrNotImplemented = errors.New("not implemented")
var ErrInvalidMsg = errors.New("invalid message")
