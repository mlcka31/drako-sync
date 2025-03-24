export const contractAddress =
  (process.env.NEXT_PUBLIC_SMART_CONTRACT_ADDRESS as `0x${string}`) ||
  "0x2334cA4eDc697824BeeAb473045DE1f56d07754C";

export const rpcUrl = "http://server1.xprts.xyz:8545";

export const githubRepo = "https://github.com/mlcka31/drako-sync";

export const backendAttestation =
  (process.env.NEXT_PUBLIC_BACKEND_ATTESTATION_URL as `0x${string}`) ||
  "https://github.com/mlcka31/zk-dungeon-backend";

export const backendApp =
  (process.env.NEXT_PUBLIC_BACKEND_APP_URL as `0x${string}`) ||
  "https://github.com/mlcka31/zk-dungeon-backend";
