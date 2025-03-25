import * as crypto from "crypto";
import * as fs from "fs";
import * as path from "path";
import * as ec from "elliptic";
import { keccak256 } from "ethers";

// Using elliptic for key derivation
const EC = new ec.ec("secp256k1");

export class WalletService {
  private keyPath: string;

  constructor(keyDirectory: string = "keys") {
    // Ensure the key directory exists
    if (!fs.existsSync(keyDirectory)) {
      fs.mkdirSync(keyDirectory, { recursive: true });
    }

    this.keyPath = path.join(keyDirectory, "private_key.pem");
  }

  /**
   * Generate a new private key and save it to disk
   * @returns The generated private key as a string
   */
  public generateAndSavePrivateKey(): string {
    // Generate a random private key
    const keyPair = EC.genKeyPair();
    const privateKey = keyPair.getPrivate("hex");

    // Save the private key to disk
    fs.writeFileSync(this.keyPath, privateKey, "utf8");

    return privateKey;
  }

  /**
   * Load the private key from disk
   * @returns The private key as a string
   * @throws Error if the private key file doesn't exist
   */
  public loadPrivateKey(): string {
    if (!fs.existsSync(this.keyPath)) {
      throw new Error("Private key file not found. Generate a key first.");
    }

    return fs.readFileSync(this.keyPath, "utf8");
  }

  public loadOrGeneratePrivateKey(): string {
    if (!this.hasPrivateKey()) {
      return this.generateAndSavePrivateKey();
    }
    return this.loadPrivateKey();
  }
  /**
   * Derive the public key from a private key
   * @param privateKey The private key as a hex string
   * @returns The public key as a hex string
   */
  public derivePublicKey(privateKey: string): string {
    const keyPair = EC.keyFromPrivate(privateKey);
    const publicKeyUncompressed = keyPair.getPublic().encode("hex", false);
    const publicKeyWithoutPrefix = publicKeyUncompressed.slice(2);
    const hash = keccak256(Buffer.from(publicKeyWithoutPrefix, "hex"));
    const address = "0x" + hash.slice(-40);
    return address;
  }

  /**
   * Get the wallet address (a simplified version of the public key)
   * @param publicKey The public key as a hex string
   * @returns The wallet address
   */
  public getWalletAddress(publicKey: string): string {
    // Create a hash of the public key
    const hash = crypto.createHash("sha256").update(publicKey, "hex").digest();
    // Take the last 20 bytes of the hash
    return hash.slice(hash.length - 20).toString("hex");
  }

  /**
   * Check if a private key exists on disk
   * @returns true if the private key file exists, false otherwise
   */
  public hasPrivateKey(): boolean {
    return fs.existsSync(this.keyPath);
  }

  /**
   * Get both private and public keys
   * @returns An object containing the private key, public key, and wallet address
   */
  public getWallet(): {
    privateKey: string;
    publicKey: string;
    address: string;
  } {
    const privateKey = this.loadPrivateKey();
    const publicKey = this.derivePublicKey(privateKey);
    const address = this.getWalletAddress(publicKey);

    return { privateKey, publicKey, address };
  }
}

export const walletService = new WalletService();
walletService.loadOrGeneratePrivateKey();
