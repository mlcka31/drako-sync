"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
Object.defineProperty(exports, "__esModule", { value: true });
exports.walletService = exports.WalletService = void 0;
const crypto = __importStar(require("crypto"));
const fs = __importStar(require("fs"));
const path = __importStar(require("path"));
const ec = __importStar(require("elliptic"));
const ethers_1 = require("ethers");
// Using elliptic for key derivation
const EC = new ec.ec("secp256k1");
class WalletService {
    constructor(keyDirectory = "keys") {
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
    generateAndSavePrivateKey() {
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
    loadPrivateKey() {
        if (!fs.existsSync(this.keyPath)) {
            throw new Error("Private key file not found. Generate a key first.");
        }
        return fs.readFileSync(this.keyPath, "utf8");
    }
    loadOrGeneratePrivateKey() {
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
    derivePublicKey(privateKey) {
        const keyPair = EC.keyFromPrivate(privateKey);
        const publicKeyUncompressed = keyPair.getPublic().encode("hex", false);
        const publicKeyWithoutPrefix = publicKeyUncompressed.slice(2);
        const hash = (0, ethers_1.keccak256)(Buffer.from(publicKeyWithoutPrefix, "hex"));
        const address = "0x" + hash.slice(-40);
        return address;
    }
    /**
     * Get the wallet address (a simplified version of the public key)
     * @param publicKey The public key as a hex string
     * @returns The wallet address
     */
    getWalletAddress(publicKey) {
        // Create a hash of the public key
        const hash = crypto.createHash("sha256").update(publicKey, "hex").digest();
        // Take the last 20 bytes of the hash
        return hash.slice(hash.length - 20).toString("hex");
    }
    /**
     * Check if a private key exists on disk
     * @returns true if the private key file exists, false otherwise
     */
    hasPrivateKey() {
        return fs.existsSync(this.keyPath);
    }
    /**
     * Get both private and public keys
     * @returns An object containing the private key, public key, and wallet address
     */
    getWallet() {
        const privateKey = this.loadPrivateKey();
        const publicKey = this.derivePublicKey(privateKey);
        const address = this.getWalletAddress(publicKey);
        return { privateKey, publicKey, address };
    }
}
exports.WalletService = WalletService;
exports.walletService = new WalletService();
exports.walletService.loadOrGeneratePrivateKey();
