package types

const (
	// ModuleName defines the module name
	ModuleName = "chat"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_chat"
)

var (
	ParamsKey = []byte("p_chat")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
