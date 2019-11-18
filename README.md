econumi-coin
====================

The goal of this blockchain is to change the mining paradigm to doing actual computational work (i.e. scientific experiments similar to [GridCoin](https://github.com/gridcoin/Gridcoin-Research), machine learning on models to build better models) instead of just hashing and verifying hashes. A similar project to this currently exists called [Golem](https://github.com/golemfactory/golem) but it written in Python and does not have deep integrations with the Docker Platform.

Design Doc
----------
Read the following design doc to learn more about the design of this product: [econumi-spec.md](project/econumi-spec.md). Also see [sprint1.md](project/sprint1.md) for an overview of the latest design in action.

Background material
-------------------

For a primer on bitcoin and cryptocurrencies this is a great read the following resources if you can:
- [Mastering Bitcoin: Programming the Open Blockchain](https://www.amazon.com/Mastering-Bitcoin-Programming-Open-Blockchain/dp/1491954388/ref=sr_1_3?ie=UTF8&qid=1516928723&sr=8-3&keywords=mastering+bitcoin)
- [Bitcoin whitepaper](https://www.bitcoin.com/bitcoin.pdf)
- [Ethereum whitepaper](https://github.com/ethereum/wiki/wiki/White-Paper)
- [Nano whitepaper](https://nano.org/en/whitepaper)
- [BlackCoin Proof of Stake](https://blackcoin.co/blackcoin-pos-protocol-v2-whitepaper.pdf)

MVP (Demo)
----------
- Run a tasks with varying degrees of difficulty as containers and get rewarded for mining
- Verify signed containers via notary in a decentralized way via notary
- Setup peer to peer network and keep a node running (https://github.com/libp2p/go-libp2p)
- Build image, endpoint into blockchain hashmap[image hash][(difficulty, endpoint)]
- Build a way to measure amount of computation and work done on miner (https://github.com/bobrik/collectd-docker)
- Run as systemd background process: https://raspberrypi.stackexchange.com/questions/52052/how-to-run-a-golang-program-in-the-background

Backlog
-------

- GPU support for containers
- Support for Apache Spark on K8s: http://spark.apache.org/releases/spark-release-2-3-0.html
- Incorporate k8s and Calico to run a container with networking across clusters. Run an nginx application that is reachable from another cluster and accessible on the internet.
- Client GUI using React (https://github.com/cgrant/go-react, https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server)
- Find a business model (provide a commercial service that will heavily be utilized by businesses)
  - Possible use cases for business
      - Allow financial instutions to be involved in a financial network in exchange for providing compute power to companies. Generate support contracts to get first customers.
      - Virtual Currency for Games, E-Game Tournaments, and Livestreaming - allows game companies to mine and distribute virtual currency paid for by grid consumers
      - Provide computation services for machine learning engines
      - Provide computation services for synthetic browser testing
      - Provide security solution for storing crypto currencies
      - Provide monitoring and management console for servers running blockchain
      - Run p2p services like PeerTube and get compensated: https://github.com/Chocobozzz/PeerTube
- Use anacrolix for peer to peer network (runs on MIT License): https://github.com/anacrolix/torrent
- Crypto SDK to utilize if need to start to build Byzantine Fault Tolerance: https://github.com/tendermint/tendermint

Prerequisites (Install before building)
---------------------------------------

- [Docker CE](https://github.com/docker/docker): Version 17.12.x ([link](https://download.docker.com/mac/stable/Docker.dmg))
- [Golang](https://golang.org/dl/): Version 1.10 (`brew install go`)
- [Glide](https://github.com/Masterminds/glide): Version 0.13.1 (`brew install glide`)

Usage
-----

Create Wallet
```
econumi createwallet
```

Start blockchain
```
econumi createblockchain
```

Print the contents of the chain
```
econumi printchain
```

Get balance
```
econumi getbalance <address>
```

Send
```
econumi send <from address> <to address> <amount>
```

Getting Started
---------------

The easiest way to get started is to clone the repository and create your own env.sh file:

Example env.sh file

```
export GOPATH=/Users/yongshin/projects/econumi-coin
export PATH=$PATH:$GOPATH/bin
export NODE_ID=3000
export DOCKER_API_VERSION='1.35'
export DATA_DIR=$GOPATH/data
```

Commands to get started:

```bash
# Add ssh key for convenience or insert into .zshrc file
ssh-add -K /Users/yongshin/.ssh/id_rsa
# Get the latest snapshot
git clone git@gitlab.com:econumi/econumi-coin.git econumi-coin
cd econumi-coin
source env.sh
cd src/econumi.org/
# Get all packages from Glide
glide install
cd $GOPATH
make
make start
```
