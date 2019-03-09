# Genesis Validator Ceremony

Welcome to the Cosmos Hub Genesis Validator Ceremony!

## What is it?

The recommended initial validator set for the Genesis State of the Cosmos Network is computed from the set of
signed `gentx` transactions with non-zero ATOMs submitted during this genesis ceremony.

Genesis transactions will be collected on Github in this repository and checked for validity by an automated script.

By participating in this ceremony and submitting a gen-tx, you are making a commitment to your fellow Cosmonauts
that you will be around to bring your validator online by the recommended genesis time of 13 March 2019 23:00 GMT to launch the network. Note that you can start `gaiad` before that time and, assuming you configure it successfully, it will automatically start the peer-to-peer and consensus processes once the genesis timestamp is reached.

Please keep the following things in mind.

1. This process is intended for technically inclined people who have participated in our testnets and Game of Stakes. If you aren't already familiar with this process, we advise against participating due to the risks involved.
2. ATOMs staked during genesis will be at risk of 5% slashing if your validator double signs. If you accidentally misconfigure your validator setup, this can easily happen, and we do not expect ATOMs so slashed to be recoverable by any means.
3. ATOMs staked during genesis or after will be locked up as part of the defense against long range attacks for 3 weeks. They can be re-delegated or undelegated, but will not be transferrable until a hard-fork enables transfers.

## Genesis File

**WARNING: THIS IS NOT THE FINAL RECOMMENDATION FOR THE GENESIS FILE**

This repository contains a work-in-progress recommendation for the genesis file called [`penultimate_genesis.json`](./penultimate_genesis.json).
It **IS NOT** the final recommended genesis file.
If you find an error in this genesis file, please contact us
immediately at "genesis at interchain dot io".

A final recommendation will be available shortly, including a justification for
all components of the genesis file and scripts to recompute it.

Anyone with an ATOM allocation in the [`penultimate_genesis.json`](./penultimate_genesis.json) who intends to participate in the genesis ceremony must submit a pull request
containing a valid `gen-tx` to this repository in the `/gentx` folder with a file name like `<moniker>.json`.

## Genesis Parameters

See [Pull Request](https://github.com/cosmos/launch/pull/2) for some discussion.

Many genesis fields are self-evident, null, or uncontroversial (e.g. gas prices, which are chosen for spam prevention).

Here we document the more subjective parameter choices and the reasons behind their recommendation.

Note that all durations are specified in nanoseconds.

### Staking Module

- `"unbonding_time": "1814400000000000"`. The unbonding time determines the duration for which bonded stake is
  held accountable for any discovered equivocations, specified in nanoseconds. 3 weeks was chosen to balance
  the concerns of a sufficient unbonding period for lite client safety and a modicum of staking token liquidity.
- `"max_validators": "100"`. The maximum validator count is the total number of validators which can be bonded
  and voting in consensus for any given block - which validators are in this set is dynamically determined
  to be the top hundred validator candidates sorted by delegated stake. The value of `100` was specified in the Cosmos whitepaper.
  It is expected to grow over time, but automatic increases aren't yet
  implemented.

### Minting Module

- `"inflation": "0.07"`. The initial annual inflation rate will be 7%, as specified in the Cosmos whitepaper.
- `"inflation_max": "0.2"`. The maximum annual inflation rate will be 20%, as specified in the Cosmos whitepaper.
- `"inflation_min": "0.07"`. The minimum annual inflation rate will be 7%, as specified in the Cosmos whitepaper.
- `"inflation_rate_change": "0.13"`. The rate at which the inflation rate changes (second derivative of inflation),
  per year squared, will be 13%, as specified in the Cosmos whitepaper.

### Distribution Module

- `"community_tax": "0.02"`. The tax on inflation and fees levied to fund the public goods pool will be 2%,
  as specified in the Cosmos whitepaper.
- `"base_proposer_reward": "0.01"`. 1% of inflation and fees (flat) will be allocated to the block proposer. This provides an incentive for validators to maintain uptime.
- `"bonus_proposer_reward": "0.04"`. 4% of inflation and fees (varying according to the fraction of precommits included)
  will be allocated to the block proposer.
- `"withdraw_addr_enabled": false`. Changing reward withdrawal addresses will be initially disabled. It may later be enabled via a hard fork.

### Governance Module

- `"min_deposit": 512atom`. The minimum deposit to bring a proposal up for a vote is 512 Atoms. Because the price of Atoms is uncertain at launch we tried to pick a value that was high enough to prevent spam proposals, while not being too expensive. As a note, the proposer doesn't have to provide the deposit. It can be crowd-funded. Proposals which pass refund all deposits.
- `"max_deposit_period": "1209600000000000"`. The duration in which a proposal can collect deposits is 14 days. We tried to choose this value to be long enough for a proposal to have time to gain support from the community.
- `"voting_period": "1209600000000000"`. The duration in which a proposal can be voted upon is 14 days. We wanted a voting period long enough that all staked Atom holders had time to participate.
- `"quorum": "0.4"`. A minimum quorum of 40% of bonded stake must vote on a proposal in order for it to be considered for passage. This is to ensure that proposals don't pass that have support from only a small segment of the community.
- `"threshold": "0.5"`. Over half the voting stake must vote in favor of a proposal in order for it to pass.
- `"veto": "0.334"`. 1/3 of voting stake vetoing a proposal prevents it from passing. This is necessitated by the 1/3 BFT safety bound,
  since 1/3 of stake could also elect to halt the chain or compromise safety.

### Slashing Module

- `"max_evidence_age": "1814400000000000"`. The maximum age of evidence possibly considered valid is 3 weeks
  (it must be the same as the unbonding period).
- `"signed_blocks_window": "10000"`. The rolling window for uptime measurement is 10,000 blocks.
- `"min_signed_per_window": "0.05"`. A minimum of 5% of the blocks in the last window must have been signed or
  else a validator will be slashed for downtime. During network launch we decided on a lenient uptime requirement that can later be increased by governance.
- `"downtime_jail_duration": "600000000000"`. Validators slashed for downtime are jailed for ten minutes. This provides a disincentive for validator downtime.
- `"slash_fraction_double_sign": "0.05"`. Validators who equivocate (double-sign a block, and thereby compromise safety)
  and are caught are slashed by 5% of their bonded stake.
- `"slash_fraction_downtime": "0.0001"`. Validators who are slashed for downtime and thereby compromise the availability
  of the network are slashed by 0.01% of their bonded stake. This is to provide additional disincentive for validator downtime.

## Instructions

Generally the steps to create a validator are as follows:

1. [Install Gaiad and Gaiacli](https://github.com/cosmos/cosmos-sdk/blob/master/docs/gaia/installation.md)

2. [Setup your fundraiser keys](https://github.com/cosmos/cosmos-sdk/blob/master/docs/gaia/delegator-guide-cli.md#restoring-an-account-from-the-fundraiser)

3. Download the [genesis file](https://raw.githubusercontent.com/cosmos/launch/master/penultimate_genesis.json) to `~/.gaiad/config/genesis.json`

4. Sign a genesis transaction:

```bash
gaiad gentx \
  --amount <amount_of_delegation> \
  --commission-rate <commission_rate> \
  --commission-max-rate <commission_max_rate> \
  --commission-max-change-rate <commission_max_change_rate> \
  --pubkey <consensus_pubkey> \
  --name <key_name>
```

This will produce a file in the ~/.gaiad/config/gentx/ folder that has a name with the format `gentx-<node_id>.json`. The content of the file should have a structure as follows:

```json
{
  "type": "auth/StdTx",
  "value": {
    "msg": [
      {
        "type": "cosmos-sdk/MsgCreateValidator",
        "value": {
          "description": {
            "moniker": "<moniker>",
            "identity": "",
            "website": "",
            "details": ""
          },
          "commission": {
            "rate": "<commission_rate>",
            "max_rate": "<commission_max_rate>",
            "max_change_rate": "<commission_max_change_rate>"
          },
          "min_self_delegation": "1",
          "delegator_address": "cosmos1msz843gguwhqx804cdc97n22c4lllfkk39qlnc",
          "validator_address": "cosmosvaloper1msz843gguwhqx804cdc97n22c4lllfkk5352lt",
          "pubkey": "<consensus_pubkey>",
          "value": {
            "denom": "uatom",
            "amount": "100000000000"
          }
        }
      }
    ],
    "fee": {
      "amount": null,
      "gas": "200000"
    },
    "signatures": [
      {
        "pub_key": {
          "type": "tendermint/PubKeySecp256k1",
          "value": "AlT62zuYGlZGUG3Yv0RtIFoPTzVY4N+WEFmBvz1syjws"
        },
        "signature": ""
      }
    ],
    "memo": ""
  }
}
```

Finally, to participate in this ceremony, Copy this file to the `gentx` folder in this repo
and submit a pull request:

```
cp ~/.gaiad/config/gentx/gentx-<node_id>.json ./gentx/<moniker>.json
```

We will only accept self delegation transactions up to 100,000 atoms for genesis. We expect 1-5% of the ATOM allocation to
be staked via genesis transactions.

## A Note about your Validator Signing Key

Your validator signing private key lives at `~/.gaiad/config/priv_validator_key.json`. If this key is stolen, an attacker would be able to make
your validator double sign, causing a slash of 5% of your atoms and the tombstoning of your validator. If you are interested in how to better protect this key please see the [`tendermint/kms`](https://github.com/tendermint/kms) repo. We will have a complete guide for how to secure this file soon after launch.

## Next Steps

Wait for the Interchain Foundation to publish a final recommendation for the
Genesis Block Release Software and be ready to come online at the recommended
time.

On initialization of the software, the Cosmos Hub Bonded Proof-of-Stake system will kick in to
determine the initial validator set (max 100 validators) from the set of `gentx` transactions.
More than 2/3 of the voting power of this set must be online and participating in consensus
in order to create the first block and start the Cosmos Hub.

We expect and hope that ATOM holders will exercise discretion in initial staking to ensure the network
does not ever become excessively centralized as we move steadily to the target of 66% ATOMs staked. This is
a first of its kind experiment in bootstrapping a decentralized network. Other proof of stake networks have
bootstrapped with the aid of a foundation or other administrator. We hope to bootstrap as a decentralized community, building on the shared experiences of many many testnets.


# Disclaimer

The Cosmos Hub is *highly* experimental software. In these early days, we can
expect to have issues, updates, and bugs. The existing tools require advanced
technical skills and involve risks which are outside of the control of the
Interchain Foundation and/or the Tendermint team (see also the risk section in
the Interchain Cosmos Contribution Terms). Any use of this open source Apache
2.0 licensed software is done at your *own risk and on a “AS IS” basis, without
warranties or conditions of any kind*, and any and all liability of the
Interchain Foundation and/or the Tendermint team for damages arising in
connection to the software is excluded. **Please exercise extreme caution!**
