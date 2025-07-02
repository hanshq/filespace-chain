# Discovery Answers - Hosting Provider Reward System

## Q1: Should rewards be paid periodically (e.g., per block/epoch) or only upon contract completion?
**Answer:** A percentage should be paid periodically, the rest when the provider successfully completes the hosting until end of contract. For now, 50/50 split.

**Implications:**
- 50% of rewards distributed periodically during active storage
- 50% held as completion bonus to incentivize full contract duration
- Need to track both periodic payments and completion status

## Q2: Will the system require proof-of-storage verification before releasing rewards?
**Answer:** Yes, but that is out of scope for now.

**Implications:**
- Design should accommodate future proof-of-storage integration
- For MVP, assume providers are honest (trust-based system)
- Payment architecture should support adding verification later

## Q3: Should the escrow amount from inquiries be locked immediately when creating a hosting inquiry?
**Answer:** Yes

**Implications:**
- Integrate with bank module to lock funds on inquiry creation
- Need escrow account or module account to hold locked funds
- Implement refund mechanism for unused escrow
- Prevent inquiry creation if insufficient funds

## Q4: Will there be penalties/slashing for providers who fail to maintain storage availability?
**Answer:** Yes, they will be slashed a percentage of their staking amount.

**Implications:**
- Providers must stake tokens before offering storage
- Need staking mechanism for providers
- Implement slashing logic for failures
- Define failure conditions and slashing percentages
- Staking is separate from escrow (provider stakes, inquirer escrows)

## Q5: Should rewards vary based on factors like storage duration, file size, or replication requirements?
**Answer:** Yes, there should be a platform-wide price for cost per byte size and block. Everything else is calculated by that (replication: 3 = 3 workers receive payments = 3x price).

**Implications:**
- Global parameter for base price per byte per block
- Total cost = file_size × duration_in_blocks × replication_factor × base_price
- Each replica provider receives equal payment
- Need governance mechanism to adjust platform-wide pricing
- Simplifies pricing model while maintaining flexibility