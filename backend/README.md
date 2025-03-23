# Backend for Web3 D&D: AI Dungeon Master

## How to build go interface for a smart contract

```shell
solc --optimize --bin --abi -o build foundry/src/AlephGameState.sol --overwrite
abigen --bin=build/AlephGameState.bin --abi=build/AlephGameState.abi --pkg=contract --out=backend/contract/AlephGameState.go
```

## How to build docker

```shell
docker build -t backend .

docker compose up -d
```