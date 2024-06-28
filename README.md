# Simplified Blockchain Implementation

Blockchain is a groundbreaking technology that revolutionizes how data is stored and verified in a decentralized manner. This project aims to explore the fundamental concepts of blockchain by building a simplified cryptocurrency prototype.

## Table of Contents

- [Introduction](#introduction)
- [Project Overview](#project-overview)
  - [Block](#1-block)
  - [Blockchain](#2-blockchain)
- [Objectives](#objectives)
- [Usage](#usage)
- [License](#license)

## Introduction

Blockchain, at its core, is a distributed database where data is stored in blocks that are linked together in a chronological order. Each block contains valuable information, such as transactions in cryptocurrencies, along with technical details like timestamps and references to previous blocks.

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

## Objectives

This project serves to:
- **Understand Blockchain Basics**: Explore how blocks and chains work together.
- **Implement Cryptographic Principles**: Utilize hashing algorithms for data security.
- **Prototype a Simplified Cryptocurrency**: Develop a basic blockchain model for educational purposes.

## Usage

```

$ go build -o blockchain.exe
$ ./blockchain.exe
Prev. hash: 
Data: Genesis Block
Hash: 78f3299b2c9d92db10b9c6ea68d9a09a821210994cd9d4d818acb9a182e0a8f5

Prev. hash: 78f3299b2c9d92db10b9c6ea68d9a09a821210994cd9d4d818acb9a182e0a8f5
Data: Send 1 BTC to Abhi
Hash: 24dca38c5a99da479a0208b9b72800c9c0b101ec7eb1c4eab09e86a91a24bb7b

Prev. hash: 24dca38c5a99da479a0208b9b72800c9c0b101ec7eb1c4eab09e86a91a24bb7b
Data: Send 2 BTC to Abhi
Hash: ad57897468481391f982c01b853f1c598025975c49ae7fcea098587d70daf43c
```
## License

Distributed under the MIT License. See LICENSE for more information.