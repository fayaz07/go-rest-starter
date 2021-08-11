<!--
(#how-to-run-the-project)
-->

## How to run the project?

> Note: Run the below commands in the root directory of the project.

1. Setup keys and `.env` file

```bash
make setup
```

2. Install the required dependencies by the following command

```bash
make install
```

### Run Server

```bash
make dev
```

### Run watchable server

```bash
reflex -r '\.go$' -s -- sh -c "go run src/main.go"
```
