# Start the game

Deployed

```shell
forge create \
  --rpc-url $RPC_URL \
  --private-key $PRIVATE_KEY \
  --legacy \
  --gas-price 1000 \
  --broadcast \
  foundry/src/AlephGameState.sol:AlephGameState \
  --constructor-args 10000000000000000 1100000000000000000 

Deployer: 0x36615Cf349d7F6344891B1e7CA7C72883F5dc049
Transaction hash: 0x3783c4167f703754a587d2c187e645b708fab865f8c194422e378566f3e2ec88
export DEPLOYED_TO=0xDABa4e2Eef03b13a153a88B7E53846b55190a778

solc --optimize --bin --abi -o build foundry/src/AlephGameState.sol --overwrite
```

```sh
export CONTRACT_ADDRESS=0x36615Cf349d7F6344891B1e7CA7C72883F5dc049
export PRIVATE_KEY=0x7726827caac94a7f9e1b160f7ea819f172f7b6f9d2a97f992c38edeab82d4110
export RPC_URL=http://server1.xprts.xyz:8545
export ABI_PATH=out/AlephGameState.sol/AlephGameState.json

cast send $DEPLOYED_TO "setGreeting(string)" "Hello, Y'all" --private-key $PRIVATE_KEY --rpc-url $RPC_URL

export DEPLOYED_TO=0xDa7410efd68ea8C0b24cb89816585D34E5Bd93Df

#cast send $DEPLOYED_TO "initGame(address)" $AGENT_ADDRESS \
cast send $DEPLOYED_TO "initGame(address)" $CONTRACT_ADDRESS \
  --private-key $PRIVATE_KEY \
  --rpc-url $RPC_URL \
  --gas-price 1000 
  
  \
  --abi $ABI_PATH
  


```
