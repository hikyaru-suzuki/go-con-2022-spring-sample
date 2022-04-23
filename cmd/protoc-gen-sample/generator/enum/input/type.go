package input

type AccessorType int32

const (
	AccessorType_OnlyServer AccessorType = iota + 1
	AccessorType_ServerAndClient
)

type Element struct {
	// Protoに定義してある名前
	RawName string
	Value   int32
	Comment string
}

type Enum struct {
	AccessorType AccessorType
	SnakeName    string
	Comment      string
	Elements     []*Element
}
