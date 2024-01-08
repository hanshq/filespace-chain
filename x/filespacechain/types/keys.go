package types

const (
	// ModuleName defines the module name
	ModuleName = "filespacechain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_filespacechain"
)

var (
	ParamsKey = []byte("p_filespacechain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
