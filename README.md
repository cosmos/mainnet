# Cosmos Hub Mainnet

The current Gaia Version of the Cosmos Hub mainnet is [v5.0.2](https://github.com/cosmos/gaia/releases/tag/v5.0.2). However if you want to build a node from scratch you need to first run [v4.2.6](https://github.com/cosmos/gaia/releases/tag/v4.2.6) until the node panics at block height [6910000](https://github.com/cosmos/gaia/blob/main/docs/migration/cosmoshub-4-delta-upgrade.md#Upgrade-will-take-place-July-12,-2021). The node should stop running after the panic, if it does not stop automatically, wait for 5-10 minutes and then kill it manually. Now install the latest version of gaia ([v5.0.2](https://github.com/cosmos/gaia/releases/tag/v5.0.2)) and then begin running the binary agian with the optional flag `--x-crisis-skip-assert-invariants`. This will begin syncing the node since the last upgrade until it is at the current height.

You can skip this process if you have access to a snapshot of the blockchain after height 6910000. You can find snapshots at [cosmos.quicksync.io](https://cosmos.quicksync.io/).

# Join the Cosmos Hub Mainnet

## Quickstart

**Instant Gratification Snippet**

```bash
git clone -b v4.2.0 https://github.com/cosmos/gaia
cd gaia
make install
gaiad init chooseanicehandle
wget https://github.com/cosmos/mainnet/raw/master/genesis.cosmoshub-4.json.gz
gzip -d genesis.cosmoshub-4.json.gz
mv genesis.cosmoshub-4.json ~/.gaia/config/genesis.json
gaiad start --p2p.seeds bf8328b66dceb4987e5cd94430af66045e59899f@public-seed.cosmos.vitwit.com:26656,cfd785a4224c7940e9a10f6c1ab24c343e923bec@164.68.107.188:26656,d72b3011ed46d783e369fdf8ae2055b99a1e5074@173.249.50.25:26656,ba3bacc714817218562f743178228f23678b2873@public-seed-node.cosmoshub.certus.one:26656,3c7cad4154967a294b3ba1cc752e40e8779640ad@84.201.128.115:26656 --x-crisis-skip-assert-invariants
```
Now wait until the chain reaches block height XXXX. It will panic and you will see a message like:
```
ERR UPGRADE "Gravity-DEX" NEEDED at height: 6910000: v5.0.0-4760cf1f1266accec7a107f440d46d9724c6fd08

panic: UPGRADE "Gravity-DEX" NEEDED at height: 6910000: v5.0.0-4760cf1f1266accec7a107f440d46d9724c6fd08
```
Then you should run the following commands:
```bash
git checkout v5.0.2 -b v5.0.2
make install
gaiad start --x-crisis-skip-assert-invariants --p2p.seeds bf8328b66dceb4987e5cd94430af66045e59899f@public-seed.cosmos.vitwit.com:26656,cfd785a4224c7940e9a10f6c1ab24c343e923bec@164.68.107.188:26656,d72b3011ed46d783e369fdf8ae2055b99a1e5074@173.249.50.25:26656,ba3bacc714817218562f743178228f23678b2873@public-seed-node.cosmoshub.certus.one:26656,3c7cad4154967a294b3ba1cc752e40e8779640ad@84.201.128.115:26656
```

Note:  If your node is unable to connect to any of the seeds listed here, you can find seeds and peers in [this document](https://hackmd.io/@KFEZk8oMTz6vBlwADz0M4A/BkKEUOsZu#) maintained by community members, and at [Atlas](https://atlas.cosmos.network/nodes), which is automatically generated by crawling the network.

If you'd like to save those seeds to your settings put them in ~/.gaia/config/config.toml in the p2p section under seeds in the same comma-separated list format.

**You need to [install gaia](./installation.md) before you go further**

**Gaia nodes on cosmoshub-4 take longer than 45 min to startup. The development team are evaluating [solutions](https://github.com/cosmos/cosmos-sdk/issues/7766).**

## Setting Up a New Node

These instructions are for setting up a brand new full node from scratch.

First, initialize the node and create the necessary config files:

```bash
gaiad init <your_custom_moniker>
```

**Note**
Monikers can contain only ASCII characters. Using Unicode characters will render your node unreachable.

You can edit this `moniker` later, in the `~/.gaia/config/config.toml` file:

```toml
# A custom human readable name for this node
moniker = "<your_custom_moniker>"
```

You can edit the `~/.gaia/config/app.toml` file in order to enable the anti spam mechanism and reject incoming transactions with less than the minimum gas prices:

```
# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

##### main base config options #####

# The minimum gas prices a validator is willing to accept for processing a
# transaction. A transaction's fees must meet the minimum of any denomination
# specified in this config (e.g. 10uatom).

minimum-gas-prices = ""
```

Your full node has been initialized!

## Genesis & Seeds

### Copy the Genesis File

Fetch the mainnet's `genesis.json` file into `gaiad`'s config directory.

```bash
mkdir -p $HOME/.gaia/config
wget https://github.com/cosmos/mainnet/raw/master/genesis.cosmoshub-4.json.gz
gzip -d genesis.cosmoshub-4.json.gz
mv genesis.cosmoshub-4.json $HOME/.gaia/config
```

If you want to connect to the public testnet instead, click [here](./join-testnet.md)

To verify the correctness of the configuration run:

```bash
gaiad start
```

### Add Seed Nodes

Your node needs to know how to find peers. You'll need to add healthy seed nodes to `$HOME/.gaia/config/config.toml`. The [`launch`](https://github.com/cosmos/launch) repo contains links to some seed nodes.

If those seeds aren't working, you can find more seeds and persistent peers on a Cosmos Hub explorer (a list can be found on the [launch page](https://cosmos.network/launch)).



## A Note on Gas and Fees

On Cosmos Hub mainnet, the accepted denom is `uatom`, where `1atom = 1.000.000uatom`

Transactions on the Cosmos Hub network need to include a transaction fee in order to be processed. This fee pays for the gas required to run the transaction. The formula is the following:

```
fees = ceil(gas * gasPrices)
```

The `gas` is dependent on the transaction. Different transaction require different amount of `gas`. The `gas` amount for a transaction is calculated as it is being processed, but there is a way to estimate it beforehand by using the `auto` value for the `gas` flag. Of course, this only gives an estimate. You can adjust this estimate with the flag `--gas-adjustment` (default `1.0`) if you want to be sure you provide enough `gas` for the transaction.

The `gasPrice` is the price of each unit of `gas`. Each validator sets a `min-gas-price` value, and will only include transactions that have a `gasPrice` greater than their `min-gas-price`.

The transaction `fees` are the product of `gas` and `gasPrice`. As a user, you have to input 2 out of 3. The higher the `gasPrice`/`fees`, the higher the chance that your transaction will get included in a block.

For mainnet, the recommended `gas-prices` is `0.025uatom`.

## Set `minimum-gas-prices`

Your full-node keeps unconfirmed transactions in its mempool. In order to protect it from spam, it is better to set a `minimum-gas-prices` that the transaction must meet in order to be accepted in your node's mempool. This parameter can be set in the following file `~/.gaia/config/app.toml`.

The initial recommended `min-gas-prices` is `0.025uatom`, but you might want to change it later.

## Pruning of State

There are three strategies for pruning state, please be aware that this is only for state and not for block storage:

1. `PruneEverything`: This means that all saved states will be pruned other than the current.
2. `PruneNothing`: This means that all state will be saved and nothing will be deleted.
3. `PruneSyncable`: This means that only the state of the last 100 and every 10,000th blocks will be saved.

By default every node is in `PruneSyncable` mode. If you would like to change your nodes pruning strategy then you must do so when the node is initialized. For example, if you would like to change your node to the `PruneEverything` mode then you can pass the `---pruning everything` flag when you call `gaiad start`.

> Note: When you are pruning state you will not be able to query the heights that are not in your store.

## Run a Full Node

Start the full node with this command:

```bash
gaiad start
```

Check that everything is running smoothly:

```bash
gaiad status
```

View the status of the network with the [Cosmos Explorer](https://cosmos.network/launch).

## Export State

Gaia can dump the entire application state to a JSON file, which could be useful for manual analysis and can also be used as the genesis file of a new network.

Export state with:

```bash
gaiad export > [filename].json
```

You can also export state from a particular height (at the end of processing the block of that height):

```bash
gaiad export --height [height] > [filename].json
```

If you plan to start a new network from the exported state, export with the `--for-zero-height` flag:

```bash
gaiad export --height [height] --for-zero-height > [filename].json
```

## Verify Mainnet

Help to prevent a catastrophe by running invariants on each block on your full
node. In essence, by running invariants you ensure that the state of mainnet is
the correct expected state. One vital invariant check is that no atoms are
being created or destroyed outside of expected protocol, however there are many
other invariant checks each unique to their respective module. Because invariant checks
are computationally expensive, they are not enabled by default. To run a node with
these checks start your node with the assert-invariants-blockly flag:

```bash
gaiad start --assert-invariants-blockly
```

If an invariant is broken on your node, your node will panic and prompt you to send
a transaction which will halt mainnet. For example the provided message may look like:

```bash
invariant broken:
    loose token invariance:
        pool.NotBondedTokens: 100
        sum of account tokens: 101
    CRITICAL please submit the following transaction:
        gaiad tx crisis invariant-broken staking supply

```

When submitting a invariant-broken transaction, transaction fee tokens are not
deducted as the blockchain will halt (aka. this is a free transaction).

## Upgrade to Validator Node

You now have an active full node. What's the next step? You can upgrade your full node to become a Cosmos Validator. The top 125 validators have the ability to propose new blocks to the Cosmos Hub. Continue onto [the Validator Setup](../validators/validator-setup.md).
