# Genesis instructions for validators

This README contains instructions for validators to prepare for and participate in the Columbus network genesis. Those of you that participated in the genesis drill should be familiar with the requisite steps. 

## Overview of steps

Generally the steps to create a validator are as follows:

1. [Install terrad and terracli version v0.1.1](https://docs.terra.money/guide/installation)

2. [Setup your genesis keys](https://docs.terra.money/guide/users)

3. Download the [genesis template file](https://raw.githubusercontent.com/terra-project/launch/master/params/genesis_template.json) to `~/.terrad/config/genesis.json`. Note that the final genesis file is currently being formed. You will be asked to replace the temporary genesis file with the genesis.json file later. 

4. Sign a genesis transaction:

```bash
terrad gentx \
  --amount <amount_of_delegation_uluna> \
  --commission-rate <commission_rate> \
  --commission-max-rate <commission_max_rate> \
  --commission-max-change-rate <commission_max_change_rate> \
  --pubkey <consensus_pubkey> \
  --name <key_name>
  --min-self-delegation "1uluna"
```

This will produce a file in the ~/.terrad/config/gentx/ folder that has a name with the format `gentx-<node_id>.json`. The content of the file should have a structure as follows:

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
          "delegator_address": "terra1msz843gguwhqx804cdc97n22c4lllfkk39qlnc",
          "validator_address": "terravaloper1msz843gguwhqx804cdc97n22c4lllfkk5352lt",
          "pubkey": "<consensus_pubkey>",
          "value": {
            "denom": "uluna",
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

__**NOTE**__: If you would like to override the memo field use the `--ip` and `--node-id` flags for the `terrad gentx` command above.

Finally, copy this file to the `gentx` folder in this repo and submit a pull request:

```
cp ~/.terrad/config/gentx/gentx-<node_id>.json ./gentx/<moniker>.json
```

## Timelines

1. ~~**22 April 2019 07:00 PST**: Validators from the Columbus genesis drill must submit their addresses via pull request to [the validator address directory](https://github.com/terra-project/accounts/validators/address.json) to reclaim their drill rewards.~~

2. ~~**22 April 2019 09:00 PST**: Terraform Labs will publish its `penultimate_genesis.json` file to be used by validators to assemble genesis transactions.~~ 

3. **23 April 2019 02:00 PST**: By this time, all genesis validators MUST submit a gentx PR to this repository to be included in the genesis block. We will not be accepting gentxs past this time.

4. **23 April 2019 05:00 PST**: At this point, Terraform Labs will publish the `genesis.json` file containing the gentx submissions from validators. Validators should replace their `penultimate_genesis.json` file with this file. 

5. **23 April 2019 23:00 PST**: The Genesis Block for Columbus will be mined.

6. **24 April 2019 02:00 PST**: Terraform Labs' ecosystem tools (finder, station) will go live for the mainnet. 


## Important Notices

1. The atomic units on Terra Core has been switched from m-asset (for example, `mluna`) to u-asset (for example, `uluna`) to be consistent with SI unit conventions. https://github.com/terra-project/core/pull/113 Make sure you are using the right nomenclature for fees and transactions. 

2. You must re-install `terrad` and `terracli`, as the software that is being used has changed from the drill. The software version that is being used has changed from `v0.1.0rc0` to `v0.1.1`. 

3. Please stay tuned to the Validator Discord channel for updates leading up to the launch. Important notices will happen in real time close to genesis, and you should stay updated to not miss out on important updates. 


## Asks

The strength of the Terra ecosystem ultimately depends on the enthusiasm and engagement of its community. If you have bandwidth & technical capability to deploy tutorials, ecosystem tools and offer dev work for the Columbus network, we ask that you do so. We will be offering additional delegation rewards to validators that actively add value to the community. 


## A Note about your Validator Signing Key

Your validator signing private key lives at `~/.terrad/config/priv_validator_key.json`. If this key is stolen, an attacker would be able to make
your validator double sign, causing a slash of 1% of your luna and the [tombstoning](https://github.com/cosmos/cosmos-sdk/blob/master/docs/spec/slashing/07_tombstone.md) of your validator. If you are interested in how to better protect this key please see the [`tendermint/kms`](https://github.com/tendermint/kms) (_*use at your own risk*_) repo. We will have a complete guide for how to secure this file soon after launch.

## Next Steps

Wait for Terraform Labs to publish a final penultimat_genesis.json file and be ready to come online at the recommended
time.

Terraform Labs will recommend a particular genesis file and software version, but there
is no guarantee a network will ever start from it - nodes and validators may
never come online, the community may disregard the recommendation and choose
different genesis files, and/or they may modify the software in arbitrary ways. Such
outcomes and many more are outside Terraform Labs's control and completely in the hands
of the community.

On initialization of the software, the Columbus Bonded Proof-of-Stake system will kick in to
determine the initial validator set (max 100 validators) from the set of `gentx` transactions.
More than 2/3 of the voting power of this set must be online and participating in consensus
in order to create the first block and start Columbus.

