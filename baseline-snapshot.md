# Pre-Scaffold Baseline Snapshot

## Current Proto Files
- file_entry.proto
- genesis.proto  
- hosting_contract.proto
- hosting_inquiry.proto
- hosting_offer.proto
- module/module.proto
- params.proto
- payment.proto (contains our custom EscrowRecord, ProviderStake)
- query.proto (contains our custom query methods)
- tx.proto

## Current Generated Types Count
- 9 .pb.go files in x/filespacechain/types/

## Current Custom Types in payment.proto
```proto
message EscrowRecord {
  uint64 inquiry_id = 1;
  cosmos.base.v1beta1.Coin amount = 2 [(gogoproto.nullable) = false];
  string creator = 3;
}

message ProviderStake {
  string provider = 1;
  cosmos.base.v1beta1.Coin amount = 2 [(gogoproto.nullable) = false];
  uint64 height = 3;
}
```

## Current Custom Query Methods in query.proto
- PaymentHistory (contract_id) -> PaymentHistoryResponse
- PaymentHistoryAll (pagination) -> PaymentHistoryAllResponse
- EscrowRecord (inquiry_id) -> EscrowRecordResponse  
- EscrowRecordAll (pagination) -> EscrowRecordAllResponse
- ProviderStake (provider) -> ProviderStakeResponse
- ProviderStakeAll (pagination) -> ProviderStakeAllResponse

## Current Query Handler Files
- query_business_logic.go (custom queries)
- query_escrow_stake.go (custom escrow/stake queries)
- query_file_entry.go (scaffold-generated)
- query_hosting_contract.go (scaffold-generated)
- query_hosting_inquiry.go (scaffold-generated)
- query_hosting_offer.go (scaffold-generated)
- query_list_hosting_contract_from.go (custom query)
- query_params.go (scaffold-generated)
- query_payment_analytics.go (custom queries)
- query_payment_history.go (custom queries)

## Current gRPC Handler Files
- grpc_query_extended.go (our custom handlers for payment, escrow, stake)

## Current Keeper Files Count