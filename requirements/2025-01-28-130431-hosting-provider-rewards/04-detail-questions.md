# Expert Requirements Questions - Hosting Provider Reward System

Based on deep analysis of the Filespace Chain codebase, these detailed questions will clarify expected system behavior for the reward implementation.

## Q6: Should the BankKeeper be added to the keeper struct in x/filespacechain/keeper/keeper.go alongside the existing fields?
**Default if unknown:** Yes (the BankKeeper is already injected in module.go but not stored in the keeper struct)

## Q7: Should we create a dedicated module account for escrow funds separate from hosting provider staking pools?
**Default if unknown:** Yes (separate concerns - escrow for inquiry funds vs staking for provider bonds)

## Q8: Should the periodic payment processing happen in BeginBlock or EndBlock hooks in x/filespacechain/module/module.go?
**Default if unknown:** BeginBlock (follows Cosmos SDK pattern where staking rewards are processed in BeginBlock)

## Q9: Should the platform-wide pricing parameter be added to the existing Params struct in x/filespacechain/types/params.go?
**Default if unknown:** Yes (maintains consistency with existing parameter management patterns)

## Q10: Should hosting providers be required to stake before creating offers, or can staking happen when accepting a contract?
**Default if unknown:** Before creating offers (prevents spam offers and ensures provider commitment before matching)