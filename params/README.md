## Genesis Parameters

Many genesis fields are self-evident, null, or uncontroversial (e.g. gas prices, which are chosen for spam prevention).
Here the more subjective parameter choices are documented with the reasons behind their recommendation.
Note that all durations are specified in nanoseconds.

Parameters that are also in the Cosmos Hub but different in Columbus are highlighted. 

### AUTH Module

- `"tx_sig_limit": "100"`.

### Staking Module

- `"unbonding_time": "1814400000000000"`. The unbonding time determines the duration for which bonded stake is held accountable for any discovered equivocations, specified in nanoseconds. 3 weeks was chosen to balance the concerns of a sufficient unbonding period for lite client safety and a modicum of staking token liquidity.
- `"max_validators": "100"`. The maximum validator count is the total number of validators which can be bonded and voting in consensus for any given block - which validators are in this set is dynamically determined to be the top hundred validator candidates sorted by delegated stake. The value of 100 was specified in the Cosmos whitepaper. It is expected to grow over time, but automatic increases aren't yet implemented.

### Distribution Module

- `community_tax: "0.0"`. Ignore cosmos community tax.
- `base_proposer_reward: "0.01"`. 1% of inflation and fees (flat) will be allocated to the block proposer. This provides an incentive for validators to be good proposers by being available when it's their turn to propose, including lots of transactions in their proposed block, and gossiping the proposed block quickly.
- `bonus_proposer_reward": "0.04"`. 4% of inflation and fees (varying according to the fraction of precommits included) will be allocated to the block proposer to incentivize them to include as many precommits from other validators as possible.
- `"withdraw_addr_enabled": true`. Changing reward withdrawal addresses will be initially enabled. 

### Market Module

- `"daily_swap_limit": "0.01"`. We impose a 1% swap cap on Luna inflation for market trades. In practice, this means that the daily inflation of Luna is at most 1%. This prevents a malicious swapper exploiting a broken oracle to vastly inflate luna supply and threaten network consensus. 

### Treasury Module

- `"tax_policy": { "rate_min": 0.0005,  "rate_max": 0.01, "change_max": 0.00025, "cap": { denom":usdr, "amount":1000000 }} `. Tax rate is between 0.1% and 1%. Max change rate at a given 4 week period is 0.025%, and the absolute value cap on any transaction is 1TerraSDR. 
- `"reward_policy": { "rate_min": 0.05,  "rate_max": 0.2, "change_max": 0.025}`. Seigniorage reward rate for miners is between 5% and 20%. Max change rate at a given 4 week period is 2.5% 
- `"seigniorage_burden_target": "0.67"` We target the burden of miner rewards to come 67% from seigniorage and 33% from tx taxes (not the same as gas fees).  
- `"mining_increment": "1.07"`, We increment mining rewards by 1.07 every 4 weeks. Implies a target return rate of ~50% barring mining token value volatility
- `"window_short": "4"`, Internal period metric, see `policy.go`
- `"window_long": "52"`, Internal period metric, see `policy.go`
- `"window_probation": "12"`, Number of weeks where updates to tx fees and reward weights are barred 
- `"oracle_share": "0.10"`, ratio of seigniorage rewards that goes to oracle votes (1-mining reward) - oracle_share
- `"budget_share": "0.9"`, ratio of seigniorage rewards that goes to budget programs (1-mining reward) - budget_share

### Budget Module

- `"active_threshold": "0.1"`. The ratio of approving votes (yes vote stake - no vote stake) that the budget vote needs to rise over until the budget program enters the active set
- `"legacy_threshold": "0.0"`.   The ratio of approving votes (yes vote stake - no vote stake) that the budget vote needs to fall under until the budget program is taken out of the active set 
- `"vote_period": "518400"`. # of blocks the vote will be held over for a given budget program. 
- `"deposit": "{\"denom\": usdr, \"amount\": 100,000,000 }"`. Amount of deposit that has to be submitted with the budget program vote. The deposit is burned if the budget program fails to enter the active set. 

### Oracle Module

- `"vote_period": "180"`. Number of blocks (approx. 15 minutes) for a oracle vote to be tallied. 
- `"vote_threshold": "0.5"`.  ratio of bonded tokens that have to submit votes before the vote can be tallied
- `"drop_threshold": "100"`. # of vote cycles that oracle tallies have to be skipped sequentially before the Terra denom is dropped from the oracle
- `"oracle_reward_band": "0.01"`. oracle rewardee range. You must submit a vote within the reward band to receive rewards for loyal oracle votes.

### Slashing Module

- `"max_evidence_age": "1814400000000000"`. The maximum age of evidence possibly considered valid is 3 weeks (it must be the same as the unbonding period).
- `"signed_blocks_window": "10000"`. The rolling window for uptime measurement is 10,000 blocks.
- `"min_signed_per_window": "0.05"`. A minimum of 5% of the blocks in the last window must have been signed or else a validator will be slashed for downtime. To nurture network launch, a lenient uptime requirement is recommended that can later be increased by governance.
- `"downtime_jail_duration": "600000000000"`. Validators slashed for downtime are jailed for ten minutes. This provides a disincentive for validator downtime.
- `"slash_fraction_double_sign": "0.01"`. Validators who equivocate (double-sign a block, and thereby compromise safety) and are caught are slashed by 1% of their bonded stake.
- `"slash_fraction_downtime": "0.0001"`. Validators who are slashed for downtime and thereby compromise the availability of the network are slashed by 0.01% of their bonded stake. This is to provide additional disincentive for validator downtime.

