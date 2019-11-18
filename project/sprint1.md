# econumi blockchain demo (sprint1)

Create two wallets for sender and receiver

```
econumi createwallet -default

econumi createwallet
```

Create blockchain

```
econumi createblockchain
```

Get balance for current default wallet

```
econumi getbalance
```

Show containers currently running

```
econumi ps

CONTAINER ID        IMAGE                       TIME ELAPSED      CPU         MEMORY
econumi-node        econumi/econumi-node:1.0    60 seconds        1 core      1 GB
```

Show containers that have finished

```
econumi ps -a

CONTAINER ID        IMAGE                       TIME ELAPSED      CPU         MEMORY    COINS MINED
econumi-node        econumi/econumi-node:1.0    60 seconds        1 core      1 GB      1
```

Show images available to run in econumi blockchain

```
econumi images
IMAGE                           
econumi/busybox             
econumi/spark
econumi/nginx
```

Start/stop workloads after joining nodes

```
econumi start econumi/nginx

econumi stop econumi/nginx

econumi start econumi/nginx -duration 1

econumi start econumi/spark
```

Show workload running

```
econumi ps

CONTAINER ID        IMAGE                       TIME ELAPSED      CPU         MEMORY
econumi-node        econumi/econumi-node:1.0    60 seconds        1 core      1 GB
```

Get balance for current default wallet

```
econumi getbalance
```

Send money from address 1 to address2

```
econumi address1 address2 100
```

Confirm that money is sent

```
econumi getbalance address2
```

Print the contents of the blockchain

```
econumi printchain
```

## TODO

- Run nginx container
- Create CPU and Memory metrics for containers
- Compute credits per node
- Distribute coins per epoch
- Get send to work

## Example applications

### Big data

- Run Apache Spark Job in a container, retrieve data from remote location and send data back to endpoint stored as environment variable on image file
- Send data to different VM running on vagrant running Consul

### Webserver - nginx

- Run Nginx Web application, have web html file baked into image

### Selenium test

- Run a web page test to check if a specific html element exists, hardcode webpage for now

### Function as a Service

- Run a function get compensated for running the service

### TensorFlow

- Run a Tensorflow computation job

## TODO:

- Run container in detached mode and exit: https://docs.docker.com/engine/reference/run/#detached-vs-foreground
- Build a verifier role for nodes within blockchain
