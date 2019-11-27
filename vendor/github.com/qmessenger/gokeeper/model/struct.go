package model

import (
	"encoding/gob"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/qmessenger/gokeeper/model/parser"
)

func init() {
	gob.Register([]StructData{})
}

type StructData struct {
	Name    string              `json:"name"`
	Version int                 `json:"version"`
	Data    map[string]ConfData `json:"data"`
}

func NewStructData(name string, version int, data map[string]ConfData) StructData {
	return StructData{Name: name, Version: version, Data: data}
}

func (s *StructData) SetVersion(version int) {
	s.Version = version
}

type ConfData struct {
	Type      string      `json:"type"`
	RawKey    string      `json:"raw_key"`
	RawValue  string      `json:"raw_value"`
	Key       string      `json:"key"`
	Value     interface{} `json:"-"`
	StructKey string      `json:"-"`
}

func NewConfData(rawKey, rawValue string) (*ConfData, error) {
	typ, key, value, err := parser.TypeParser(rawKey, rawValue)
	if err != nil {
		return nil, err
	}
	return &ConfData{Type: typ, Key: key, RawKey: rawKey, RawValue: rawValue, Value: value, StructKey: ToCamlCase(key)}, nil
}

// ToCamlCase convert snake_case to CamlCase
func ToCamlCase(key string) string {
	ks := strings.Split(key, "_")
	for k, v := range ks {
		ks[k] = ToUpperFirst(v)
	}
	return strings.Join(ks, "")
}

// ToUpperFirst return first letter to upper
func ToUpperFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

func GetStructName(fname string) string {
	fname = filepath.Base(fname)
	f := strings.Split(fname, ".")
	return ToUpperFirst(f[0])
}
