package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		FileEntryList:       []FileEntry{},
		HostingInquiryList:  []HostingInquiry{},
		HostingContractList: []HostingContract{},
		HostingOfferList:    []HostingOffer{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in fileEntry
	fileEntryIdMap := make(map[uint64]bool)
	fileEntryCount := gs.GetFileEntryCount()
	for _, elem := range gs.FileEntryList {
		if _, ok := fileEntryIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for fileEntry")
		}
		if elem.Id >= fileEntryCount {
			return fmt.Errorf("fileEntry id should be lower or equal than the last id")
		}
		fileEntryIdMap[elem.Id] = true
	}
	// Check for duplicated ID in hostingInquiry
	hostingInquiryIdMap := make(map[uint64]bool)
	hostingInquiryCount := gs.GetHostingInquiryCount()
	for _, elem := range gs.HostingInquiryList {
		if _, ok := hostingInquiryIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for hostingInquiry")
		}
		if elem.Id >= hostingInquiryCount {
			return fmt.Errorf("hostingInquiry id should be lower or equal than the last id")
		}
		hostingInquiryIdMap[elem.Id] = true
	}
	// Check for duplicated ID in hostingContract
	hostingContractIdMap := make(map[uint64]bool)
	hostingContractCount := gs.GetHostingContractCount()
	for _, elem := range gs.HostingContractList {
		if _, ok := hostingContractIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for hostingContract")
		}
		if elem.Id >= hostingContractCount {
			return fmt.Errorf("hostingContract id should be lower or equal than the last id")
		}
		hostingContractIdMap[elem.Id] = true
	}
	// Check for duplicated ID in hostingOffer
	hostingOfferIdMap := make(map[uint64]bool)
	hostingOfferCount := gs.GetHostingOfferCount()
	for _, elem := range gs.HostingOfferList {
		if _, ok := hostingOfferIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for hostingOffer")
		}
		if elem.Id >= hostingOfferCount {
			return fmt.Errorf("hostingOffer id should be lower or equal than the last id")
		}
		hostingOfferIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
