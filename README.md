# Cosmos Hub Upgrade 1

**Transfer enablement**

According to governance proposal #3 and #5, the state of the `cosmoshub-1` chain was exported at block 500.000 and migrated to `cosmoshub-2` using [this guide](https://github.com/cosmos/cosmos-sdk/wiki/Cosmos-Hub-1-Upgrade).

The new genesis file is called `genesis.cosmoshub-2.json`.


```
$ shasum -a 256 genesis.cosmoshub-2.json
1e349fb39b85f7707ee78d39879f9d5d61f4d30f67980bb0bf07bd35b2f8bf30  genesis.cosmoshub-2.json
$ b2sum genesis.json
1910abe394fc80e0e4ebc7d9388219363a94cf67bc19544f6bd33e01f5afbf38282e832ae321ba62cfd201179febf9b4a7380c7e4a1b1dabaf85ce2649831e24  genesis.cosmoshub-2.json
```


# Cosmos Hub Launch

This is the Interchain Foundation's recommendation for the Genesis Block Release
Software and marks the initiation of [phase
one](https://blog.cosmos.network/the-3-phases-of-the-cosmos-hub-mainnet-fdff3a68c4c0) of the Cosmos Hub launch.

Please be aware that there is no guarantee a network will start from this
recommendation - nodes and validators may never come online, the community may disregard the
recommendation and choose different genesis files, and/or they may modify the
software in arbitrary ways. Such outcomes and many more are outside the Interchain
Foundation's control and completely in the hands of the community.

The recommended genesis file is [genesis.json](https://raw.githubusercontent.com/cosmos/launch/master/genesis.json). It has the
following SHA256 hash:

```
$ shasum -a 256 genesis.json 
73a866b21723ecbc28b6d15951b2eb3aa2f2443650ff6df489bf55ac5edceefa  genesis.json
$ b2sum genesis.json 
8c90b58efe9e0959953fe27ba431137c24c514e357b8025f5252c85ea7401247a909fac95313b907bb48579c6e389b4bbf06df626bff19aae554028964fa189d  genesis.json
```

It includes a genesis time of `2019-03-13 23:00:00 UTC`.
Please read [GENESIS.md](GENESIS.md) for details on how it was generated and
to recompute it for yourself.

The recommended software version is [v0.33.0 of the
Cosmos-SDK](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.33.0).
See the [installation
instructions](https://cosmos.network/docs/gaia/installation.html)
and the [guide to joining mainnet](https://cosmos.network/docs/gaia/join-mainnet.html).

Users wishing to interact with the network should carefully review [how to
protect themselves](https://cosmos.network/atom-protection) and the security
advisories in the recent blog post on 
[preparing for main net
launch](https://blog.cosmos.network/cosmos-hub-to-launch-mainnet-a453d2247a34).

Please note that this is *highly* experimental software. In these early days, we can
expect to have issues, updates, and bugs. The existing tools require advanced
technical skills and involve risks which are outside of the control of the
Interchain Foundation and/or the Tendermint team (see also the risk section in
the Interchain Cosmos Contribution Terms). Any use of this open source Apache
2.0 licensed software is done at your *own risk and on a “AS IS” basis, without
warranties or conditions of any kind*, and any and all liability of the
Interchain Foundation and/or the Tendermint team for damages arising in
connection to the software is excluded. **Please exercise extreme caution!**

Further, please note that it remains in the community's sole discretion to
adopt or not to adopt the recommended Genesis Block Release Software. Therefore, the Interchain
Foundation cannot guarantee that (i) ATOMs will be created and (ii) the recommended
allocation as set forth in GENESIS.md will actually take place. The recommended Genesis Block
Release Software has no support for interoperability (IBC), and the Atoms will not be
transferable.


## Seed Nodes

We request known community members who wish to run public p2p seed nodes make pull requests to add community run seed nodes below.

```
Seed nodes

- `3e16af0cead27979e1fc3dac57d03df3c7a77acc@3.87.179.235:26656` - Bison Trails
- `ba3bacc714817218562f743178228f23678b2873@public-seed-node.cosmoshub.certus.one:26656` - Certus One
- `2626942148fd39830cb7a3acccb235fab0332d86@173.212.199.36:26656` - syncnode
- `3028c6ee9be21f0d34be3e97a59b093e15ec0658@91.205.173.168:26656` - syncnode
- `89e4b72625c0a13d6f62e3cd9d40bfc444cbfa77@34.65.6.52:26656` - Cryptium Labs (@adrianbrink, @awasunyin, @cwgoes)
```
