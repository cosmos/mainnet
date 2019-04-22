# Genesis Backgrounder

This document outlines the genesis parameters (including genesis luna allocations) and explains the project history & future roadmap that led to such allocations. 

- [Project History](#Project-History)
- [Token Allocation & Fundraiser](#Token-Allocation-and-Fundraiser)
- [Validator Drill](#Validator-Drill)
- [Looking Ahead](#Looking-Ahead)


## Project History

Terra Money was created in January 2018 with the singular vision of facilitating the mass adoption of cryptocurrencies by creating digitially native assets that are price-stable against the world's major fiat currencies. Keeping in mind that previous innovations in the technology of money was bootstrapped by large payment networks (Alipay with Taobao, Paypal with eBay, Visa with banks), Terra was born with the support of the Terra Alliance, 15 large e-commerce companies in Asia that collectively process 25 billion USD in annualized transaction volume and 45 million users. The vision of the project is that with the adoption and user engagement of a massive payment network, we will be able to, for the first time, bootstrap a blockchain payment network to the scale it deserves. 
 

##  Token Allocation and Fundraiser

Terra Core facilitates the creation of many fiat-pegged currencies, facilitated by Luna the staking token. The genesis issuance of Luna is 1 billion, with initial allocations broken down in accordance with the following:  

- **Terraform Labs (10%)**: Used to facilitate the research & development of the Terra Project. can be found in the `accounts/terraform-labs` directory. 
- **Employees & Contributor Pool (20%)**: Used to compensate employees and contributors of the project. Some part of the coins have already been granted, and can be found in the `accounts/employees` directory. Currently 4.7% of this pool has been granted. 
- **Terra Alliance (20%)**: [The Terra Alliance](https://medium.com/terra-money/state-of-the-terra-alliance-d7f3ff8f6411?fbclid=IwAR2xyZ2sRi_gTHeNPH8tL_VoXpvmDq3sdWMwXaSQCAbHhQGhIEx-yHxWRio) is key to driving early adoption and usage for Terra. We will be using this pool to set incentives, mainly marketing discount programs (such as coupons for users) and volume incentives for alliance partners. Terraform Labs will be playing custodian for this pool, taking input from the community to best allocate resources from this pool. 
- **Stability Reserves (20%)**: Bootstrapping stablecoins is no easy feat, with threats to the peg coming from every adversarial angle. Stability reserves will be to manage the network's early stability close to genesis. 
- **Backers (26%)**:

In order to finance the development of the Terra Project, Terraform Labs held three token sales: 

1. Pre-seed sale: 10 cents per Luna, sold ~100M tokens. Lockup 12 months, with 30% early liquidity. 
2. Seed-sale: 23 cents per Luna, sold ~100M tokens. Lockup 10 - 18 months, with 30% early liquidity. 
3. Private-sale: 80 cents per Luna, sold ~60M tokens. Lockup 3 months, with 6 months linear vest thereafter. 

Several backers have failed to meet the guidelines and/or deadlines to reclaim their tokens, leading significant allocations to not be included in the genesis account generation. 

- **Genesis liquidity (4%)**: 4% of Luna will be made available to the market close to genesis to allow everyday users to use and interact with it. 

## Validator Drill

On April 19th, 2019, [30 validators](https://medium.com/terra-money/countdown-to-terra-mainnet-launch-f8c0b998c12a) gathered together to conduct a drill for the Columbus network genesis. The goal for the drill was simple: "exploit the broken oracle." Terraform Labs would use its voting stake to submit faulty price votes to the blockchain, which would allow arbitragers to benefit by sending in swap trades that would capture the arbitrage opportunity. 

::**Note**:For more information in how swaps & stability work in Terra, please read up on our [docs](https://docs.terra.money/features/terra-stability). 

The results from the drill are [here](drill_results.json). We had a very high level of engagement with the drill participants, with several of the validators getting outrageous returns. The winners of the drill are: 

- **The Staking Fund** (#1 returns) Also created terra.stake.id, which proved to be instrumental in monitoring the state of the drill. 
- **Cosmostation** (#2 returns, even managed to perform a consensus attack by capturing returns in luna and staking rapidly)
- **ATEAM** (#3 returns)
- **Castlenode** (#4 returns)
- **Chorus-one** (#5)

Outside of the winning list, we recognize the following validators: 
- **Dokia Capital**: managed to find a [bug](https://github.com/terra-project/core/pull/112) while minting a non-whitelisted currency.
- **B-Harvest**: actively educated other validators during the course of the drill, and showed the highest level of understanding / activity.
- **Certus One**: Helped to design and run the drill, as well as the setup of our own validators. 

Terraform Labs will delegate at least 1 million Luna to each validator mentioned above shortly after genesis. Each validator must charge a commission rate of less than 20% and meet our key management guidelines, which will be communicated separately. 

All validators that took part in the genesis drill will be granted 100 Luna to boostrap their validators. We will be working closely with the initial validator set close to genesis to offer delegation rewards for community engagement and ecosystem tool development. 

The full retrospective from the drill is [here](https://docs.google.com/document/d/1cNWCK8GyIfb1CWfn6IaoO8T80mxNsfv6rA3S06hb5As). 


## Looking Ahead

A few weeks after network launch, Columbus will [power payment services](http://fortune.com/2018/08/29/cryptocurrency-exchanges-back-32-million-stable-coin-project/) across the Terra Alliance. It will also be used by a federation including the Mongolian government to [pay for taxes and utilities in the country](https://www.forbes.com/sites/yoavvilner/2019/01/11/mongolia-starts-off-2019-with-its-eyes-on-crypto-payment-adoption/?fbclid=IwAR0GqylqQev_7JhnEGdi7lJwDgfaMZRMODqVXZbos30z1eqqlLO1lSJ9_Nk). 

Though the protocol has been robustly been tested over the course of several [testnets](https://github.com/terra-project/networks) and [simulations](https://github.com/terra-project/research), having millions of real users interact with the system across a diverse range of use cases may break or test our assumptions. The Terra community may be obliged to update the Terra Core software to better accomodate the needs of its users.