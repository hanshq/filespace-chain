# Discovery Questions for Hosting Provider Reward System

These questions will help understand the problem space and requirements for implementing a reward system for hosting providers in Filespace Chain.

## Q1: Should rewards be paid periodically (e.g., per block/epoch) or only upon contract completion?
**Default if unknown:** Yes, periodic payments (this provides continuous incentive for providers to maintain service)

## Q2: Will the system require proof-of-storage verification before releasing rewards?
**Default if unknown:** Yes (essential for preventing fraud and ensuring providers actually store the files)

## Q3: Should the escrow amount from inquiries be locked immediately when creating a hosting inquiry?
**Default if unknown:** Yes (ensures funds are available for payment and prevents double-spending)

## Q4: Will there be penalties/slashing for providers who fail to maintain storage availability?
**Default if unknown:** Yes (necessary to ensure service quality and discourage bad actors)

## Q5: Should rewards vary based on factors like storage duration, file size, or replication requirements?
**Default if unknown:** Yes (different storage requirements should have different compensation levels)