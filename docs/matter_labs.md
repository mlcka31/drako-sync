# How to work with Matter labs Deployment

## Keys

[rich-wallets.json](rich-wallets.json)
Taken from [here](https://github.com/matter-labs/local-setup/blob/main/rich-wallets.json)


## Test deployment

```sh
export CONTRACT_ADDRESS=0x36615Cf349d7F6344891B1e7CA7C72883F5dc049
export PRIVATE_KEY=0x7726827caac94a7f9e1b160f7ea819f172f7b6f9d2a97f992c38edeab82d4110
export RPC_URL=http://server1.xprts.xyz:8545
export LOCAL_RPC_URL=localhost:8545

# Check remote connection
telnet 66.29.155.254 8081
telnet 66.29.155.254 8545
# or
curl -v http://66.29.155.254:8081
```

Other commands
```shell
export CONTRACT_ADDRESS=0x36615Cf349d7F6344891B1e7CA7C72883F5dc049
export PRIVATE_KEY=0x7726827caac94a7f9e1b160f7ea819f172f7b6f9d2a97f992c38edeab82d4110
export RPC_URL=http://server1.xprts.xyz:8545

# Check balance 
cast balance $CONTRACT_ADDRESS --rpc-url $RPC_URL

# send ETH from one account to another 
export ADDRESS_2=0xa61464658AfeAf65CccaaFD3a512b69A83B77618
cast send --rpc-url $RPC_URL --private-key $PRIVATE_KEY $ADDRESS_2 --value 1ether

export DEPLOYED_TO=0xDa7410efd68ea8C0b24cb89816585D34E5Bd93Df
```

```shell
# How to deploy contract
forge create \
  --rpc-url $RPC_URL \
  --private-key $PRIVATE_KEY \
  --legacy \
  --gas-price 1000 \
  src/HelloWorld.sol:HelloWorld --broadcast

export DEPLOYED_TO=0xDa7410efd68ea8C0b24cb89816585D34E5Bd93Df
cast call $DEPLOYED_TO "greeting()(string)" --rpc-url $RPC_URL
cast send $DEPLOYED_TO "setGreeting(string)" "Hello, Y'all" --private-key $PRIVATE_KEY --rpc-url $RPC_URL

cast send $DEPLOYED_TO "setGreeting(string)" "Hello, Y'all" \
  --private-key $PRIVATE_KEY --rpc-url $RPC_URL \
  --gas-price 1000 
  
cast call $DEPLOYED_TO "greeting()(string)" --rpc-url $RPC_URL

export ADRESS_2=0xa61464658AfeAf65CccaaFD3a512b69A83B77618
cast send --rpc-url $RPC_URL --private-key $PRIVATE_KEY $ADRESS_2 --value 1ether
cast send --rpc-url $RPC_URL --private-key $PRIVATE_KEY $ADRESS_2 \ 
  --gas-price 1000 --value 1ether 
```

```shell
forge create \
  --rpc-url $RPC_URL \
  --private-key $PRIVATE_KEY \
  --legacy \
  --gas-price 1000 \
  --broadcast \
  foundry/src/AlephGameState.sol:AlephGameState \
  --constructor-args 10000000000000000 1100000000000000000 

export DEPLOYED_TO=0xd90d56f99f3F70c990B634b30Ca9FE83956E8A05
Deployer: 0x36615Cf349d7F6344891B1e7CA7C72883F5dc049
Transaction hash: 0xbfb75910b857f3d4d109fe96a50262c53a7d4a56854dd569cfc7e92431b9a028
```

