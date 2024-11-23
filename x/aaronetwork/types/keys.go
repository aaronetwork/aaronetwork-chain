package types

const (
	// ModuleName defines the module name
	ModuleName = "aaronetwork"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_aaronetwork"
)

var (
	ParamsKey = []byte("p_aaronetwork")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
