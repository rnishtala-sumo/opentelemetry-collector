// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configcompression // import "go.opentelemetry.io/collector/config/configcompression"

import "fmt"

// Type represents a compression method
type Type string

type Level int

type CompressionParams struct {
	Level Level `mapstructure:"level"`
}

const (
	TypeGzip    Type = "gzip"
	TypeZlib    Type = "zlib"
	TypeDeflate Type = "deflate"
	TypeSnappy  Type = "snappy"
	TypeZstd    Type = "zstd"
	TypeLz4     Type = "lz4"
	TypeNone    Type = "none"
	TypeEmpty   Type = ""
)

// IsCompressed returns false if CompressionType is nil, none, or empty.
// Otherwise, returns true.
func (ct *Type) IsCompressed() bool {
	return *ct != TypeEmpty && *ct != TypeNone
}

func (ct *Type) UnmarshalText(in []byte) error {
	typ := Type(in)
	if typ == TypeGzip ||
		typ == TypeZlib ||
		typ == TypeDeflate ||
		typ == TypeSnappy ||
		typ == TypeZstd ||
		typ == TypeLz4 ||
		typ == TypeNone ||
		typ == TypeEmpty {
		*ct = typ
		return nil
	}
	return fmt.Errorf("unsupported compression type %q", typ)
}
