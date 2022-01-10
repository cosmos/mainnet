# Cosmos Hub Mainnet

## Overview

The current Gaia Version of the Cosmos Hub mainnet is [`v6.0.0`](https://github.com/cosmos/gaia/releases/tag/v6.0.0). To bootstrap a mainnet node, it is possible to sync from `V6.0.0` via Quicksync or via [State Sync](https://hub.cosmos.network/main/hub-tutorials/join-mainnet.html#state-sync).

However if you want to build a node from scratch you need to first run [v4.2.6](https://github.com/cosmos/gaia/releases/tag/v4.2.6) until the node panics at block height [6910000](https://github.com/cosmos/gaia/blob/main/docs/migration/cosmoshub-4-delta-upgrade.md#Upgrade-will-take-place-July-12,-2021). The node should stop running after the panic, if it does not stop automatically, wait for 5-10 minutes and then kill it manually. Now install the latest version of gaia ([v5.0.2](https://github.com/cosmos/gaia/releases/tag/v5.0.2)) and then begin running the binary agian with the optional flag `--x-crisis-skip-assert-invariants`. This will begin syncing the node since the last upgrade until it is at the current height.

## Quickstart

**Preresquisites**
- `make` & `gcc`
- `Go 1.16+`

> **Note**: Make sure to have all prerequisites installed. See the [installation docs](https://hub.cosmos.network/main/getting-started/installation.html) for clarification and a detailed set of instructions.

**Quicksync**

Quicksync.io offers several daily snapshots of the Cosmos Hub with varying levels of pruning (archive 1.4TB, default 540GB, and pruned 265GB). For downloads and installation instructions, visit the [Cosmos Quicksync guide](https://quicksync.io/networks/cosmos.html).

**State Sync**

To enable state sync, visit an [explorer](https://www.mintscan.io/cosmos/blocks) to get a recent block height and corresponding hash. A node operator can choose any height/hash in the current bonding period, but as the recommended snapshot period is 1000 blocks, it is advised to choose something close to current height - 1000. Set these parameters in the code snippet below `<BLOCK_HEIGHT>` and `<BLOCK_HASH>`

For reference, the list of `rpc_servers` and `persistent` peers can be found in the [cosmos hub chain-registry repo](https://github.com/cosmos/chain-registry/blob/master/cosmoshub/chain.json).

```bash
# Build gaiad binary and initialize chain
cd $HOME
git clone -b v6.0.0 https://github.com/cosmos/gaia
cd gaiad
make install
gaiad init <custom moniker>

# Prepare genesis file for cosmoshub-4
wget https://github.com/cosmos/mainnet/raw/master/genesis.cosmoshub-4.json.gz
gzip -d genesis.cosmoshub-4.json.gz
mv genesis.cosmoshub-4.json $HOME/.gaia/config/genesis.json

#Set minimum gas price & peers
sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0.001uatom"/' app.toml
sed -i 's/persistent_peers = ""/persistent_peers = "6e08b23315a9f0e1b23c7ed847934f7d6f848c8b@165.232.156.86:26656,ee27245d88c632a556cf72cc7f3587380c09b469@45.79.249.253:26656,538ebe0086f0f5e9ca922dae0462cc87e22f0a50@34.122.34.67:26656,d3209b9f88eec64f10555a11ecbf797bb0fa29f4@34.125.169.233:26656,bdc2c3d410ca7731411b7e46a252012323fbbf37@34.83.209.166:26656,585794737e6b318957088e645e17c0669f3b11fc@54.160.123.34:26656,5b4ed476e01c49b23851258d867cc0cfc0c10e58@206.189.4.227:26656"/' config.toml

# Configure State sync
cd $HOME/.gaia/config
sed -i 's/enable = false/enable = true/' config.toml
sed -i 's/trust_height = 0/trust_height = <BLOCK_HEIGHT>/' config.toml
sed -i 's/trust_hash = ""/trust_hash = "<BLOCK_HASH>"/' config.toml
sed -i 's/rpc_servers = ""/rpc_servers = "https:\/\/rpc.cosmos.network:443,https:\/\/rpc.cosmos.network:443"/' config.toml

#Start Gaia
gaiad start --x-crisis-skip-assert-invariants
```

**Sync from Scratch**

```bash
# Build gaiad binary and initialize chain
git clone -b v4.2.1 https://github.com/cosmos/gaia
cd gaia
make install
gaiad init <custom moniker>

# Prepare genesis file for cosmoshub-4
wget https://github.com/cosmos/mainnet/raw/master/genesis.cosmoshub-4.json.gz
gzip -d genesis.cosmoshub-4.json.gz
mv genesis.cosmoshub-4.json ~/.gaia/config/genesis.json


#Set minimum gas price & peers
sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0.001uatom"/' app.toml
sed -i 's/persistent_peers = ""/persistent_peers = "6e08b23315a9f0e1b23c7ed847934f7d6f848c8b@165.232.156.86:26656,ee27245d88c632a556cf72cc7f3587380c09b469@45.79.249.253:26656,538ebe0086f0f5e9ca922dae0462cc87e22f0a50@34.122.34.67:26656,d3209b9f88eec64f10555a11ecbf797bb0fa29f4@34.125.169.233:26656,bdc2c3d410ca7731411b7e46a252012323fbbf37@34.83.209.166:26656,585794737e6b318957088e645e17c0669f3b11fc@54.160.123.34:26656,5b4ed476e01c49b23851258d867cc0cfc0c10e58@206.189.4.227:26656"/' config.toml

gaiad start --x-crisis-skip-assert-invariants
```
Now wait until the chain reaches block height 6910000. It will panic and log the following:
```
ERR UPGRADE "Gravity-DEX" NEEDED at height: 6910000: v5.0.0-4760cf1f1266accec7a107f440d46d9724c6fd08

panic: UPGRADE "Gravity-DEX" NEEDED at height: 6910000: v5.0.0-4760cf1f1266accec7a107f440d46d9724c6fd08
```

It's now time to perform the manual Delta upgrade:
```bash
git checkout -b v5.0.2
make install
gaiad start --x-crisis-skip-assert-invariants
```

Once `V5` reaches the upgrade block height, the chain will halt and display the following message:
```
ERR UPGRADE "Vega" NEEDED at height: 8695000

```

This will indicate it is time to perform the Vega upgrade. Similar with the previous upgrade, checkout `V6`, compile the new binary and restart `gaiad`

```bash
git checkout -b v6.0.0
make install
gaiad start --x-crisis-skip-assert-invariants
```

> _NOTE_:  If the node is unable to connect to any of the seeds listed here, find additional seeds and peers in [this document](https://hackmd.io/@KFEZk8oMTz6vBlwADz0M4A/BkKEUOsZu#) maintained by community members, and at [Atlas](https://atlas.cosmos.network/nodes), which is automatically generated by crawling the network. Additionally, node operators can just copy [Quicksync's addressbook](https://quicksync.io/addrbook.cosmos.json) and move it to `$HOME/.gaia/config/addrbook.json`


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

## Exporting State

Gaia can dump the entire application state into a JSON file. This application state dump is useful for manual analysis and can also be used as the genesis file of a new network.

Export state with:

```bash
gaiad export > [filename].json
```

It is also possible to export state from a particular height (at the end of processing the block of that height):

```bash
gaiad export --height [height] > [filename].json
```

If planning to start a new network from the exported state, export with the `--for-zero-height` flag:

```bash
gaiad export --height [height] --for-zero-height > [filename].json
```


## Verify Mainnet

Help to prevent a catastrophe by running invariants on each block on your full
node. In essence, by running invariants the node operator ensures that the state of mainnet is the correct expected state. One vital invariant check is that no atoms are being created or destroyed outside of expected protocol, however there are many other invariant checks each unique to their respective module. Because invariant checks are computationally expensive, they are not enabled by default. To run a node with these checks start your node with the assert-invariants-blockly flag:

```bash
gaiad start --assert-invariants-blockly
```

If an invariant is broken on the node, it will panic and prompt the operator to send a transaction which will halt mainnet. For example the provided message may look like:

```bash
invariant broken:
    loose token invariance:
        pool.NotBondedTokens: 100
        sum of account tokens: 101
    CRITICAL please submit the following transaction:
        gaiad tx crisis invariant-broken staking supply

```

## Upgrade to Validator Node

You now have an active full node. What's the next step? You can upgrade your full node to become a Cosmos Validator. The top 125 validators have the ability to propose new blocks to the Cosmos Hub. Continue onto [the Validator Setup](../validators/validator-setup.md).
