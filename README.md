# Simplified Blockchain Implementation

Blockchain is a groundbreaking technology that revolutionizes how data is stored and verified in a decentralized manner. This project aims to explore the fundamental concepts of blockchain by building a simplified cryptocurrency prototype.

## Table of Contents

- [Introduction](#introduction)
- [Project Overview](#project-overview)
  - [Block](#1-block)
  - [Blockchain](#2-blockchain)
  - [Proof of work](#3-proof-of-work)
- [Objectives](#objectives)
- [Example](#example)
- [License](#license)

## Introduction

Blockchain, at its core, is a distributed database where data is stored in blocks that are linked together in a chronological order. Each block contains valuable information, such as transactions in cryptocurrencies, along with technical details like timestamps and references to previous blocks.

#### Tools used:
- Golang
- Boltdb
- 

## Project Overview

In this project, we focus on two primary components of a blockchain:

### 1. Block

Blocks in a blockchain store essential data and technical information:
- **Timestamp**: When the block is created.
- **Data**: Valuable information, such as transactions.
- **PrevBlockHash**: Hash of the previous block, creating a link in the chain.
- **Hash**: Hash of the current block, ensuring data integrity and security.

### 2. Blockchain

The blockchain itself is an ordered list of blocks:
- **Structure**: A sequential chain where each block is linked to its predecessor.
- **Functionality**: Allows adding new blocks, ensuring the integrity of the entire chain through cryptographic hashing and consensus mechanisms.
- **Storage**: The chain is persistently stored in a database, where each block and its hash are key-value pairs.

### 3. Proof of Work

Proof of work is a mechanism used by blockchain networks to validate transactions and add new blocks to the chain. It requires miners (network nodes) to solve complex computational puzzles by repeatedly guessing values until they find a solution that meets specific criteria. The first miner to solve the puzzle and broadcast the valid solution to the network is rewarded with cryptocurrency. This process ensures the security and integrity of the blockchain by making it computationally expensive to tamper with the ledger. 

- **Time to solve**: Increasing the number of bits of the target exponentially increases how long it takes to solve the puzzle. I.e. Setting the target bits to 16 takes a few seconds, but, setting it to 24 takes a few minutes.

## Objectives

This project serves to:
- **Understand Blockchain Basics**: Explore how blocks and chains work together.
- **Implement Cryptographic Principles**: Utilize hashing algorithms for data security.
- **Prototype a Simplified Cryptocurrency**: Develop a basic blockchain model for educational purposes.

## Example

```

$ go build -o blockchain.exe
$ ./blockchain.exe printchain
No existing blockchain found. Creating a new one...
Mining the block containing "Genesis Block"
0000e4cb28adb98c0d32ec5cc3fcea495bf9f635664f180c497d07336779649c

Prev. hash:
Data: Genesis Block
Hash: 0000e4cb28adb98c0d32ec5cc3fcea495bf9f635664f180c497d07336779649c
PoW: true

$ ./blockchain.exe addblock -data "Send 1 BTC to Abhi"
Mining the block containing "Send 1 BTC to Abhi"
000085e7c4ce313a5a4d472bbb92000c341e39a2276f0fb32be1b04793add050

Success!

$ ./blockchain.exe addblock -data "Pay 0.5 BTC for pizza"
Mining the block containing "Pay 0.5 BTC for pizza"
0000205381a1a5fe0447997d46ba01e3668d5bbd7e1c13dc8a99ef3694bafb31

Success!

$ ./blockchain.exe printchain
Prev. hash: 000085e7c4ce313a5a4d472bbb92000c341e39a2276f0fb32be1b04793add050
Data: Pay 0.5 BTC for pizza
Hash: 0000205381a1a5fe0447997d46ba01e3668d5bbd7e1c13dc8a99ef3694bafb31
PoW: true

Prev. hash: 0000e4cb28adb98c0d32ec5cc3fcea495bf9f635664f180c497d07336779649c
Data: Send 1 BTC to Abhi
Hash: 000085e7c4ce313a5a4d472bbb92000c341e39a2276f0fb32be1b04793add050
PoW: true

Prev. hash:
Data: Genesis Block
Hash: 0000e4cb28adb98c0d32ec5cc3fcea495bf9f635664f180c497d07336779649c
PoW: true

```
## License

Distributed under the MIT License. See LICENSE for more information.