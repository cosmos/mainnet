# Terra Core & Columbus Launch
![banner](launch-banner.png)

This repository contains the configuration parameters for the Columbus Network Genesis. Columbus is the first decentralized network of nodes communicating over [Terra Core](https://github.com/terra-project/core). 

**_The Columbus mainnet will go live on April 23rd, 23:00 PST_**. 

## Important Notes for Validators

1. **For genesis Validators**: Please follow [these instructions](INSTRUCTIONS.md) to prepare for genesis, claim rewards from the drill, and get ready for genesis. 

:: At this time, Please use [this genesis file](./penultimate_genesis.json) to create your gentx.json file, and send a pull request to the /gentx directory. You must do this by  **23 April 2019 02:00 PST** to participate in the genesis.

2. **To stay updated for the genesis**: Please monitor the validator chat on Discord in real time to stay coordinated with Terraform Labs and the rest of the community regarding the launch. 

## Genesis Background

Terra is a project that is made possible thanks to the collective effort of its global community. Those of you interested in finding out the background of its genesis please see [here](./GENESIS.md) and our [blog](https://medium.com/terra-money). Final parameter settings for Columbus can be found [here](./params.README.md). 

## Terra community 

Community channels actively being moderated are here:
- [Website](https://terra.money/)
- [Discord](https://discord.gg/bYfyhUT)
- [Telegram](https://t.me/terra_announcements)
- [Twitter](https://twitter.com/terra_money)
- [YouTube](https://goo.gl/3G4T1z)

We will be making announcements regarding the launch using these channels, so please stay tuned. 

## Reference files

### Genesis files
 
- `/accounts`: Contains addresses and genesis allocations for employees, supporters, terraform labs, and validators. 
- `penultimate_genesis.json`: The near-final genesis file, minus the gentx data validators must create and submit with this as the reference. 
- `/gentx`: Genesis transactions to create validators, submitted by each validator. 
- `/genesis.json`: The final genesis file that will be used to launch columbus.

### Seed Nodes

We request known community members who wish to run public p2p seed nodes make pull requests to add community run seed nodes below.

```
Known seed node list: 


```

## Disclaimer

The foundational software for the Columbus mainnet, Terra Core, is *highly* experimental software. In these early days, we can expect to have issues, updates, and bugs. The existing tools require advanced technical skills and involve risks which are outside of the control of Terraform Labs or its developers. Any use of this open source Apache 2.0 licensed software is done at your *own risk and on a “AS IS” basis, without warranties or conditions of any kind*. **Please exercise extreme caution!**



