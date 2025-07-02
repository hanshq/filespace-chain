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

const (
	FileEntryKey      = "FileEntry/value/"
	FileEntryCountKey = "FileEntry/count/"
)

const (
	HostingInquiryKey      = "HostingInquiry/value/"
	HostingInquiryCountKey = "HostingInquiry/count/"
)

const (
	HostingContractKey      = "HostingContract/value/"
	HostingContractCountKey = "HostingContract/count/"
)

const (
	HostingOfferKey      = "HostingOffer/value/"
	HostingOfferCountKey = "HostingOffer/count/"
)

const (
	PaymentHistoryKey = "PaymentHistory/value/"
)
