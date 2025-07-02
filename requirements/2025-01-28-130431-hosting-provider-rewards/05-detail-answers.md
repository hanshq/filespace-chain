# Expert Requirements Answers - Hosting Provider Reward System

## Q6: Should the BankKeeper be added to the keeper struct in x/filespacechain/keeper/keeper.go alongside the existing fields?
**Answer:** Yes (need to store BankKeeper reference to call methods like SendCoinsFromAccountToModule for escrow)

**Implications:**
- Update keeper struct to include bankKeeper field
- Modify NewKeeper constructor to accept BankKeeper parameter
- Add methods for escrow fund management using bank operations

## Q7: Should we create a dedicated module account for escrow funds separate from hosting provider staking pools?
**Answer:** Yes

**Implications:**
- Create separate module accounts for escrow vs staking to maintain clean separation of concerns
- Escrow account holds inquiry funds, staking account holds provider bonds
- Different permission models and access patterns for each account type

## Q8: Should the periodic payment processing happen in BeginBlock or EndBlock hooks in x/filespacechain/module/module.go?
**Answer:** BeginBlock

**Implications:**
- Implement periodic payment logic in BeginBlock hook
- Follows Cosmos SDK pattern where rewards are processed early in block
- Process 50% periodic payments before other block operations

## Q9: Should the platform-wide pricing parameter be added to the existing Params struct in x/filespacechain/types/params.go?
**Answer:** Yes, should be possible to change later by governance

**Implications:**
- Add base price per byte per block parameter to existing Params
- Integrate with governance module for parameter updates
- Enable community-driven pricing adjustments over time
- Maintain parameter validation and default values

## Q10: Should hosting providers be required to stake before creating offers, or can staking happen when accepting a contract?
**Answer:** Before creating offers

**Implications:**
- Providers must stake tokens before creating any hosting offers
- Prevents spam offers and ensures provider commitment upfront
- Validates provider stake amount during offer creation transaction
- Reject offer creation if insufficient stake