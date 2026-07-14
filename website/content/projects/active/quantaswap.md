---
aliases:
    - /projects/active/quantaswap/
    - /projects/archived/quantaswap/
audited: false
audits: []
author: DigitalGuards
categories:
    - defi
category: defi
clients:
    - platform: web
      url: https://quantaswap.io/
      github: https://github.com/DigitalGuards/QuantaSwap
      default: true
contract_address: Q94cd8e406d2bb4ea251dce3f0558941f2ac056ee
created: "2026-07-11"
default_client:
    platform: web
    url: https://quantaswap.io/
    github: https://github.com/DigitalGuards/QuantaSwap
    default: true
default_client_github: https://github.com/DigitalGuards/QuantaSwap
default_client_url: https://quantaswap.io/
description: 'Cross-chain atomic swaps between native QRL and native ETH using hashed timelock contracts on both chains. No bridge, no wrapped assets, no custodian: each leg settles natively and the swap secret never leaves the client.'
discord: ""
docs: ""
ecosystem_type: dapp
features:
    - HTLC atomic swaps between QRL v2 testnet and Ethereum Sepolia
    - Ownerless, non-upgradeable contracts with no pause switch; one artifact deploys to both chains
    - Two-party order book for coordination only; clients re-verify every step on-chain
    - QRL leg connects via the MyQRLWallet relay or the QRL browser extension, ETH leg via EIP-6963 wallets
gallery:
    - type: image
      path: quantaswap/screenshot1.png
      caption: QuantaSwap main dashboard showing active swaps and order book.
    - type: image
      path: quantaswap/screenshot2.png
      caption: Play both sides of an atomic swap from one browser and watch the HTLC handshake happen live on both chains.
    - type: image
      path: quantaswap/screenshot3.png
      caption: Plain language guide on how QuantaSwap works and how to use it.
github: https://github.com/DigitalGuards/QuantaSwap
license: GPL-3.0
logos:
    - path: quantaswap/icon.svg
      description: ""
network: testnet
open_source: true
project-types:
    - dapps
project_type: dapp
project_url: https://quantaswap.io/
status: development
tags:
    - atomic-swaps
    - htlc
    - cross-chain
    - ethereum
title: QuantaSwap
token: none
twitter: https://x.com/DigitalGuards
updated: "2026-07-11"
url: /projects/quantaswap/
---

QuantaSwap is an open-source dApp for trustless cross-chain swaps between
the two native coins: QRL on the QRL 2.0 network and ETH on Ethereum. Swaps
use the classic hashed timelock contract (HTLC) construction with sha256
hashlocks, so there is no bridge, no wrapped asset, and no operator holding
funds. Either both legs complete or both refund after their timelocks.

The same contract artifact is deployed on both chains: the QRL v2 testnet at
Q94cd8e406d2bb4ea251dce3f0558941f2ac056ee and Ethereum Sepolia at
0x805100Fa4310B9c0dbb0754E14CbDe827E3b8a3c. The contracts have no owner, no
pause, and no upgrade path, and claims are permissionless with the recipient
fixed at lock time. An order book server handles coordination only; both
parties independently verify locks on-chain before acting.

