package server

import (
	"errors"
	"sync"
)

var (
	ErrOffsetNotFound = errors.New("offset not found")
)

func NewLog() *Log {
	return &Log{}
}

type Log struct {
	mu      sync.Mutex
	records []Record
}

func (l *Log) Append(record Record) (uint64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	record.Offset = uint64(len(l.records))
	l.records = append(l.records, record)
	return record.Offset, nil
}

func (l *Log) Read(offset uint64) (Record, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if offset >= uint64(len(l.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return l.records[offset], nil
}

type Record struct {
	Value  []byte `json:"value,omitempty"`
	Offset uint64 `json:"offset,omitempty"`
}
