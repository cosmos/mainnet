# Cosmos Hub Launch

This is the Interchain Foundation's recommendation for the Genesis Block Release
Software and marks the initiation of [phase
one](https://blog.cosmos.network/the-3-phases-of-the-cosmos-hub-mainnet-fdff3a68c4c0) of the Cosmos Hub launch.

Please be aware that there is no guarantee a network will start from this
recommendation - nodes and validators may never come online, the community may disregard the
recommendation and choose different genesis files, and/or they may modify the
software in arbitrary ways. Such outcomes and many more are outside the Interchain
Foundation's control and completely in the hands of the community.

The recommended genesis file is [genesis.json](genesis.json). It has the
following SHA256 hash:

```
$ shasum -a 256 genesis.json 
73a866b21723ecbc28b6d15951b2eb3aa2f2443650ff6df489bf55ac5edceefa  genesis.json
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
Syncnode
2626942148fd39830cb7a3acccb235fab0332d86@173.212.199.36:26656
3028c6ee9be21f0d34be3e97a59b093e15ec0658@91.205.173.168:26656
```

