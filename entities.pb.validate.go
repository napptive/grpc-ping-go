// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ping/entities.proto

package grpc_ping_go

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _entities_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on PingRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PingRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RequestNumber

	// no validation rules for Data

	return nil
}

// PingRequestValidationError is the validation error returned by
// PingRequest.Validate if the designated constraints aren't met.
type PingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PingRequestValidationError) ErrorName() string { return "PingRequestValidationError" }

// Error satisfies the builtin error interface
func (e PingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PingRequestValidationError{}

// Validate checks the field values on PingResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PingResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RequestNumber

	// no validation rules for Data

	return nil
}

// PingResponseValidationError is the validation error returned by
// PingResponse.Validate if the designated constraints aren't met.
type PingResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PingResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PingResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PingResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PingResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PingResponseValidationError) ErrorName() string { return "PingResponseValidationError" }

// Error satisfies the builtin error interface
func (e PingResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPingResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PingResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PingResponseValidationError{}
