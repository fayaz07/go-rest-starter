## Setup go lang

### Set path in `~/.bashrc`

```bash
export GOPATH=$HOME/go
export GOROOT="/snap/go/9848"

export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```

### Install `reflex`

- Install reflex

```bash
go install github.com/cespare/reflex@latest
```
