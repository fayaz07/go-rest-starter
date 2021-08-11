## Setup environment

1. [Install Go & set path](#installing-go)
2. [Install MongoDb](#mongodb-setup)
3. [Install Redis](#redis-setup)

### Installing Go

Go ahead to [Golang](https://golang.org/doc/install)'s official website and install it.

#### Set Go path & GOPATH

Preferably `.bashrc` or `profilerc`, set path variables as shown in the example below.

> **GOPATH**: Where all of your dependencies will be stored (those which you download from internet),
> **Go Path**: Where GoLang is installed

```bash
export GOPATH=/home/fayaz/GoPath
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:/usr/local/go/bin
```

### MongoDB setup

Go ahead to [MongoDB](https://docs.mongodb.com/manual/installation/)'s official installation docs page and install it.

#### Start MongoDB service

For linux/mac:

```bash
sudo systemctl start mongod
```

### Redis setup

Go ahead to [Redis Labs](https://redis.io/download)'s official installation docs page and install it.

#### Start Redis service

For linux/mac:

```bash
sudo systemctl start redis-cli
```
