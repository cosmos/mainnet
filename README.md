# Genesis Validator Ceremony

Welcome to the Cosmos Hub Genesis Validator Ceremony.

The recommended initial validator set for the Genesis State of the Cosmos Network is computed from the set of
signed `gentx` transactions with non-zero ATOMs submitted during this genesis ceremony.

Genesis transactions will be collected on Github in this repository and checked for validity by an automated script.

This repository contains a work-in-progress recommendation for the genesis file
- [`penultimate_genesis.json`](./penultimate_genesis.json)
It **IS NOT** the final recommended genesis file. 
If you find an error in this genesis file, please contact us
immediately at "genesis at interchain dot io".
A final recommendation will be available shortly, including a justification for
all components of the genesis file and scripts to recompute it.

Anyone with an ATOM allocation in the [`penultimate_genesis.json`](./penultimate_genesis.json) who intends to participate in the genesis ceremony must submit a pull request
containing a valid `gen-tx` to this repository in the `/gentx` folder with a file name like `<moniker>.json`.

Please keep the following things in mind.

1. This process is intended for technically inclined people who have participated in our testnets and Game of Stakes. If you aren't already familiar with this process, we advise against participating due to the risks involved.
2. ATOMs staked during genesis will be at risk of 5% slashing if your validator double signs. If you accidentally misconfigure your validator setup, this can easily happen, and we do not expect ATOMs so slashed to be recoverable by any means. 
3. ATOMs staked during genesis or after will be locked up as part of the defense against long range attacks for 3 weeks. They can be re-delegated or undelegated, but will not be transferrable until a hard-fork enables transfers.

Generally the steps to create a validator are as follows.

[Install Gaiad and Gaiacli](https://github.com/cosmos/cosmos-sdk/blob/master/docs/gaia/installation.md)

[Setup your fundraiser keys](https://github.com/cosmos/cosmos-sdk/blob/master/docs/gaia/delegator-guide-cli.md#restoring-an-account-from-the-fundraiser)

Download the [genesis file](https://raw.githubusercontent.com/cosmos/launch/master/penultimate_genesis.json) to `~/.gaiad/config/genesis.json`

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

To participate in the genesis ceremony, copy this file to the `gentx` folder in this repo and submit a pull request:

```
cp ~/.gaiad/config/gentx/gentx-<node_id>.json ./gentx/<moniker>.json
```

We will only accept self delegation transactions up to 100,000 atoms for genesis. We expect 1-5% of the ATOM allocation to 
be staked via genesis transactions.

On initialization, the Cosmos Hub Bonded Proof-of-Stake system will kick in to 
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
