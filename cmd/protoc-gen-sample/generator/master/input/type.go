package input

import (
	"github.com/scylladb/go-set/i32set"
)

type AccessorType int32

const (
	// AccessorType_All デフォルト値がALL
	AccessorType_All AccessorType = iota
	AccessorType_OnlyAdmin
	AccessorType_OnlyServer
	AccessorType_OnlyClient
	AccessorType_AdminAndServer
	AccessorType_AdminAndClient
	AccessorType_ServerAndClient
)

type AccessorSet struct {
	*i32set.Set
}

func (a AccessorSet) Contains(accessorType AccessorType) bool {
	return a.Has(int32(accessorType))
}

func NewAccessorSet(accessorTypes ...AccessorType) *AccessorSet {
	s := i32set.NewWithSize(len(accessorTypes))
	for _, accessorType := range accessorTypes {
		s.Add(int32(accessorType))
	}

	return &AccessorSet{s}
}

type TypeKind int32

const (
	TypeKind_Bool TypeKind = iota + 1
	TypeKind_Int32
	TypeKind_Int64
	TypeKind_Float32
	TypeKind_String
	TypeKind_Enum
	// TypeKind_Message DBに乗らないものにしか使われない
	TypeKind_Message
)

type FieldType = string

const (
	FieldType_Bool    = "bool"
	FieldType_Int32   = "int32"
	FieldType_Int64   = "int64"
	FieldType_Float32 = "float32"
	FieldType_String  = "string"
)

type FieldOptionDDLFK struct {
	TableSnakeName         string
	ColumnSnakeName        string
	ParentColumnSnakeNames []string
}

type FieldOptionDDL struct {
	PK bool
	// nilチェックが必要
	FK       *FieldOptionDDLFK
	Size     uint32
	Nullable bool
}

type FieldOptionValidate struct {
	Key   string
	Value string
}

type FieldOption struct {
	AccessorType AccessorType
	DDL          *FieldOptionDDL
	Validates    []*FieldOptionValidate
}

type Field struct {
	SnakeName string
	Comment   string
	Number    int32
	// TypeKind_Enumの場合はEnum名が入る
	Type     FieldType
	TypeKind TypeKind
	IsList   bool
	Option   *FieldOption
}

type Index struct {
	SnakeNameKeys []string
}

type MessageOptionDDL struct {
	Indexes []*Index
}

type MessageOption struct {
	AccessorType AccessorType
	DDL          *MessageOptionDDL
}

type Message struct {
	Messages  []*Message
	SnakeName string
	Comment   string
	Fields    []*Field
	Option    *MessageOption
}
