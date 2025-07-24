# Blockchain Architecture Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                        USER INTERFACE                          │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │  CLI Commands (cli/cli.go)                              │    │
│  │  • createwallet    • getbalance    • send              │    │
│  │  • listaddresses   • printchain    • createblockchain  │    │
│  └─────────────────────────────────────────────────────────┘    │
└─────────────────┬───────────────────────────────────────────────┘
                  │
┌─────────────────▼───────────────────────────────────────────────┐
│                    WALLET SYSTEM                               │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │  wallet/                                                │    │
│  │  • wallet.go: Individual wallets & key management      │    │
│  │  • wallets.go: Multiple wallet management              │    │
│  │  • utils.go: Address generation & validation           │    │
│  └─────────────────────────────────────────────────────────┘    │
└─────────────────┬───────────────────────────────────────────────┘
                  │
┌─────────────────▼───────────────────────────────────────────────┐
│                  BLOCKCHAIN CORE                               │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │  blockchain/                                            │    │
│  │  • blockchain.go: Main chain management                │    │
│  │  • block.go: Individual block structure                │    │
│  │  • transaction.go: Transaction handling                │    │
│  │  • proof.go: Proof of Work mining                      │    │
│  │  • utxo.go: Unspent output tracking                    │    │
│  │  • chain_iter.go: Block iteration                      │    │
│  │  • merkle.go: Transaction verification                 │    │
│  └─────────────────────────────────────────────────────────┘    │
└─────────────────┬───────────────────────────────────────────────┘
                  │
┌─────────────────▼───────────────────────────────────────────────┐
│                 STORAGE & NETWORK                              │
│  ┌─────────────────┐           ┌─────────────────────────────┐  │
│  │  BadgerDB       │           │  Network (network/)         │  │
│  │  • Blocks       │           │  • Peer-to-peer comm.      │  │
│  │  • UTXO Set     │           │  • Node synchronization    │  │
│  │  • Wallet files │           │  • Transaction broadcast   │  │
│  └─────────────────┘           └─────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘

FLOW OF OPERATIONS:
1. User runs CLI command
2. CLI calls appropriate wallet/blockchain functions
3. Blockchain validates and processes transactions
4. New blocks are mined using Proof of Work
5. Data is stored in BadgerDB
6. Updates are broadcast to network peers
```

## Transaction Flow Example:

```
Alice wants to send 5 coins to Bob:

1. CLI: send -from Alice -to Bob -amount 5
   │
2. Wallet: Sign transaction with Alice's private key
   │
3. Blockchain: Validate Alice has 5+ coins (check UTXO)
   │
4. Transaction: Create new transaction
   │
5. Mining: Find nonce for new block (Proof of Work)
   │
6. Block: Add transaction to new block
   │
7. Storage: Save block to BadgerDB
   │
8. UTXO: Update who owns which coins
   │
9. Network: Broadcast new block to other nodes
```

## Block Structure:

```
┌─────────────────────────────────────────┐
│               BLOCK HEADER              │
├─────────────────────────────────────────┤
│ Hash: 0004aa5ef... (Block fingerprint)  │
│ PrevHash: 000123... (Previous block)    │
│ Nonce: 12345 (Mining proof)             │
├─────────────────────────────────────────┤
│              TRANSACTIONS               │
├─────────────────────────────────────────┤
│ Transaction 1: Alice → Bob (5 coins)    │
│ Transaction 2: Mining Reward (20 coins) │
│ Transaction 3: Bob → Charlie (2 coins)  │
└─────────────────────────────────────────┘
```
