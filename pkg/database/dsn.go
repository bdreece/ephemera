package database

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

const (
	pragma     string = "_pragma"
	timeFormat string = "_time_format"
	txLock     string = "_tx_lock"
)

var (
	ErrParseDSN    = errors.New("database: failed to parse DSN")
	ErrParseTxLock = errors.New("database: failed to parse TX lock")
)

//go:generate go tool stringer -type JournalingMode -linecomment
type JournalingMode int

const (
	ModeDelete   JournalingMode = iota // DELETE
	ModeTruncate                       // TRUNCATE
	ModePersist                        // PERSIST
	ModeMemory                         // MEMORY
	ModeWAL                            // WAL
	ModeOff                            // OFF
)

type Pragma struct {
	name  string
	value any
}

// String implements [fmt.Stringer].
func (p Pragma) String() string { return fmt.Sprintf("%s(%v)", p.name, p.value) }

func ApplicationID(id int32) Pragma {
	const name string = "application_id"
	return Pragma{name, id}
}

func JournalMode(mode string) Pragma {
	const name string = "journal_mode"
	return Pragma{name, mode}
}

func Synchronous(enable bool) Pragma {
	const name string = "synchronous"

	var val string
	if enable {
		val = "on"
	} else {
		val = "off"
	}

	return Pragma{name, val}
}

// A TxLock defines the transaction-locking behavior for a SQLite connection.
//
//go:generate go tool stringer -type TxLock -linecomment
type TxLock int

const (
	// Transaction deferred until initial DB access (default).
	LockDeferred TxLock = iota // deferred
	// Transaction begins immediately with shared lock.
	LockImmediate // immediate
	// Transaction begins immediately with exclusive lock
	LockExclusive // exclusive
)

func (lock *TxLock) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	case LockDeferred.String():
		*lock = LockDeferred
	case LockImmediate.String():
		*lock = LockImmediate
	case LockExclusive.String():
		*lock = LockExclusive
	default:
		return ErrParseTxLock
	}

	return nil
}

type SQLiteParams struct {
	url.Values
}

func (p SQLiteParams) Pragmas() []string { return p.Values[pragma] }

func (p SQLiteParams) AddPragma(pragma string, arg any) {
	var ok bool
	if arg, ok = arg.(string); ok {
		arg = fmt.Sprintf("'%s'", arg)
	}

	p.Add(pragma, fmt.Sprintf("%s(%v)", pragma, arg))
}

func (p SQLiteParams) TimeFormat() string { return p.Get(timeFormat) }

func (p SQLiteParams) TxLock() TxLock {
	var lock TxLock
	if err := lock.UnmarshalText([]byte(p.Get(txLock))); err != nil {
		return LockDeferred
	}

	return lock
}

func (p SQLiteParams) SetTxLock(lock TxLock) {
	p.Set(txLock, lock.String())
}

// A DSN (data source name) provides connection details for a database.
type DSN struct {
	// The path to the SQLite database file
	Path string

	// The connection parameters
	Params SQLiteParams
}

// Returns the DSN formatted as a file URL.
func (dsn DSN) URL() *url.URL {
	return &url.URL{
		Scheme:   "file",
		OmitHost: true,
		RawPath:  dsn.Path,
		RawQuery: dsn.Params.Encode(),
	}
}

// String implements [fmt.Stringer].
func (dsn DSN) String() string {
	return dsn.URL().String()
}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (dsn *DSN) UnmarshalText(text []byte) error {
	url, err := url.Parse(string(text))
	if err != nil {
		return errors.Join(ErrParseDSN, err)
	}

	dsn.Path = url.Path
	dsn.Params = SQLiteParams{
		Values: url.Query(),
	}

	return nil
}

// Set implements [flag.Value].
func (dsn *DSN) Set(s string) error {
	return dsn.UnmarshalText([]byte(s))
}

func ParseDSN(s string) (*DSN, error) {
	var dsn DSN
	if err := dsn.UnmarshalText([]byte(s)); err != nil {
		return nil, err
	}

	return &dsn, nil
}
