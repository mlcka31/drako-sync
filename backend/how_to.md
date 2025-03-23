# How to 


```shell
solc --optimize --bin --abi -o build foundry/src/AlephGameState.sol --overwrite
abigen --bin=build/AlephGameState.bin --abi=build/AlephGameState.abi --pkg=contract --out=backend/contract/AlephGameState.go
```