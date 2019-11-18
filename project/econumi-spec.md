# econumi - A distributed block chain for distributed computation and work (Draft)

The initial experience for someone using econumi is to download the econumi cli binary (via apt or brew) and run a command to start a econumi node as a container. The assumption is that the user already has a Docker engine installed, which can be pre-installed with brew and apt if not already found on the system.

```
econumi start
```

The econumi node will generate its own CA, and public and private keys or optionally utilize an existing PKI. The econumi cli is heavily dependent on the Docker engine and will utilize REST API calls to run an initial econumi image that serves as both a client and server node.

```
econumi ps

CONTAINER ID        IMAGE
econumi-node        econumi/econumi-node:1.0
```

A console command can also be used to interact with that initial econumi container to select which workloads to pick from the available images that exist for download across the network.

```
econumi console
> show images
IMAGE                           
econumi/busybox             
econumi/spark
econumi/redis
> start econumi/redis
```

Upon choosing to run using the `start` command within the console to choose a specific workload, the image gets added to a config file called `econumi.config` which is essentially a JSON file to store the type of workload selected for future use.

Create wallet in blockchain:

```
econumi createwallet

```

Get balance for addresses

```
econumi getbalance <address>
```

Send coins

```
econumi send <from address> <to address> <amount>
```

## Proof of work using containers

Currently we will be using a Proof of Work model (read [Proof of work vs Proof of Stake](https://blockgeeks.com/guides/proof-of-work-vs-proof-of-stake/)). Each node is both server and client and downloads entire blockchain. The servers take care of handling new transactions that contribute to the building of blockchain after the work has been verified by the container the receives results from the producer container.

### Validating work on each node

In order to prove that work was indeed done on the container, the econumi blockchain daemon will perform the following on the node that the container runs on:

1) Generate a private and public key set build specifically for the container and expose the private and public key via a Docker secret to the container. If the container were to exit, the container will produce a JSON file that includes the length of time the container ran, the start and stop times, resources, and its public key to validate that work had indeed been done.

2) Randomly poll the container for a container healthcheck once every 5 seconds within the 5 second interval. If the container is not healthy then the node will not be eligible for a reward during the epoch.

Based on the JSON file produce from Step 1, we will compute the amount of compute units for that container within the epoch.

### Rewarding a block

In Bitcoin, miners participate in a race to compute a block header hash using a nonce, and compare that hash to a certain target to "win" a new block that will be added to the blockchain. The number of transactions is determined for the block based on the current size of the blockchain.

With econumi since we are not technically hashing to win a race between other nodes for a proof of work, we will rely on the computation resources calculated to run a container as a proof of work. Once every epoch, a lottery will be held based on the number of compute units that are generated and one is picked from the lottery to be added as a block as the winner of the lottery.

The number of compute units are determined for a node is computed as an integer score (rounded down) based on the length of time the container ran and the resources utilized for running the container on the node (i.e. cpu and memory reservations set on the container). After the epoch completes, the blockchain node calculates the total number of compute units it generated and then participates in a distributed lottery. The winner is selected and a block is generated with a hash from the winning compute unit.

### Block size over time

You should be able to get the current number of transactions allocated per block by running the following command, which defines the number of coins that are issued per block:

```
econumi getblockreward
```

Below is a script that determines how many blocks are issued over time for Bitcoin as a reference:

```
# Original block reward for miners was 50 BTC for Bitcoin
start_block_reward = 50
# 210000 is around every 4 years with a 10 minute block interval
reward_interval = 210000

def max_money():
    # 50 BTC = 50 0000 0000 Satoshis
    current_reward = 50 * 10**8
    total = 0
    while current_reward > 0:
        total += reward_interval * current_reward
        current_reward /= 2
    return total

print "Total BTC to ever be created:", max_money(), "Satoshis”
```

## TODO

- Determine transaction fees (keep low)
- Determine whether a spec for executing containers and how to send back and verify results. Determine how to prevent a result where consumer always submits a result that is APPROVED.
- Research another attempt to verify work using bytecode instrumentation via an agent per container.
