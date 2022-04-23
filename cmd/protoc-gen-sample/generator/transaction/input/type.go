package input

import (
	"github.com/scylladb/go-set/i32set"
)

type MessageAccessorType int32

const (
	MessageAccessorType_OnlyServer MessageAccessorType = iota
	MessageAccessorType_OnlyClient
	MessageAccessorType_OnlyClientWithCommonResponse
	MessageAccessorType_ServerAndClient
	MessageAccessorType_ServerAndClientWithCommonResponse
)

type MessageAccessorSet struct {
	*i32set.Set
}

func (a MessageAccessorSet) Contains(accessorType MessageAccessorType) bool {
	return a.Has(int32(accessorType))
}

func NewMessageAccessorSet(accessorTypes ...MessageAccessorType) *MessageAccessorSet {
	s := i32set.NewWithSize(len(accessorTypes))
	for _, accessorType := range accessorTypes {
		s.Add(int32(accessorType))
	}

	return &MessageAccessorSet{s}
}

type FieldAccessorType int32

const (
	// FieldAccessorType_All デフォルト値がALL
	FieldAccessorType_All FieldAccessorType = iota
	FieldAccessorType_OnlyServer
	FieldAccessorType_OnlyClient
)

type FieldAccessorSet struct {
	*i32set.Set
}

func (a FieldAccessorSet) Contains(accessorType FieldAccessorType) bool {
	return a.Has(int32(accessorType))
}

func NewFieldAccessorSet(accessorTypes ...FieldAccessorType) *FieldAccessorSet {
	s := i32set.NewWithSize(len(accessorTypes))
	for _, accessorType := range accessorTypes {
		s.Add(int32(accessorType))
	}

	return &FieldAccessorSet{s}
}

type TypeKind int32

const (
	TypeKind_Bool TypeKind = iota + 1
	TypeKind_Int32
	TypeKind_Int64
	TypeKind_String
	TypeKind_Enum
	// TypeKind_Message クライアント向けのレスポンスにしか使われない
	TypeKind_Message
)

type FieldType = string

const (
	FieldType_Bool   = "bool"
	FieldType_Int32  = "int32"
	FieldType_Int64  = "int64"
	FieldType_String = "string"
)

type InterleaveAnnotationType int32

type Interleave struct {
	TableSnakeName string
}

type MasterRef struct {
	TableSnakeName  string
	ColumnSnakeName string
}

type FieldOptionDDL struct {
	PK bool
	// nilチェックが必要
	MasterRef *MasterRef
}

type FieldOption struct {
	AccessorType FieldAccessorType
	DDL          *FieldOptionDDL
}

type Field struct {
	SnakeName string
	Comment   string
	// TypeKind_Enumの場合はEnum名が入る
	Type     FieldType
	TypeKind TypeKind
	IsList   bool
	Number   int32
	Option   *FieldOption
}

type IndexKey struct {
	SnakeName string
	Desc      bool
}

type Index struct {
	Keys         []*IndexKey
	Unique       bool
	SnakeStoring []string
}

type MessageOptionDDL struct {
	Indexes []*Index
	// nilチェックが必要
	Interleave *Interleave
}

type MessageOption struct {
	AccessorType MessageAccessorType
	DDL          *MessageOptionDDL
}

type Message struct {
	Messages  []*Message
	SnakeName string
	Comment   string
	Fields    []*Field
	Option    *MessageOption
}
