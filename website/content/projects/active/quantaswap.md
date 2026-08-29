---
aliases:
    - /projects/active/quantaswap/
    - /projects/archived/quantaswap/
availability: live
capabilities:
    - atomic-swap
    - dex
    - wallet-connector
categories:
    - interoperability-messaging-data
    - finance
data_updated_at: "2026-07-14"
deployments:
    - id: qrl-2-testnet-v2
      network: qrl-2-testnet-v2
      operational_state: live
      identifiers:
        - type: contract
          value: Q94cd8e406d2bb4ea251dce3f0558941f2ac056ee
      evidence:
        - https://github.com/DigitalGuards/QuantaSwap
      source_verification: unverified
description: 'Cross-chain atomic swaps between native QRL and native ETH using hashed timelock contracts on both chains. No bridge, no wrapped assets, no custodian: each leg settles natively and the swap secret never leaves the client.'
display_status: Beta · Testnet
features:
    - HTLC atomic swaps between QRL v2 testnet and Ethereum Sepolia
    - Ownerless, non-upgradeable contracts with no pause switch; one artifact deploys to both chains
    - Two-party order book for coordination only; clients re-verify every step on-chain
    - QRL leg connects via the MyQRLWallet relay or the QRL browser extension, ETH leg via EIP-6963 wallets
gallery:
    - type: youtube
      id: rIwqXEEpaaM
      caption: Full testnet demo of a cross-chain atomic ETH to QRL swap.
    - type: image
      path: quantaswap/screenshot1.png
      caption: QuantaSwap main dashboard showing active swaps and order book.
    - type: image
      path: quantaswap/screenshot2.png
      caption: Play both sides of an atomic swap from one browser and watch the HTLC handshake happen live on both chains.
    - type: image
      path: quantaswap/screenshot3.png
      caption: Plain language guide on how QuantaSwap works and how to use it.
id: quantaswap
keywords:
    - atomic-swaps
    - htlc
    - cross-chain
    - ethereum
links:
    - type: website
      url: https://quantaswap.io/
    - type: application
      url: https://quantaswap.io/
      platform: web
      primary: true
    - type: social
      url: https://x.com/DigitalGuards
listed_at: "2026-07-11"
logos:
    - path: quantaswap/icon.svg
      description: QuantaSwap logo
maintainer_records:
    - name: DigitalGuards
      contact: https://github.com/DigitalGuards/QuantaSwap
maintainers:
    - DigitalGuards
maturity: beta
platforms: []
primary_category: interoperability-messaging-data
primary_link:
    type: application
    url: https://quantaswap.io/
    platform: web
    primary: true
primary_url: https://quantaswap.io/
project-types:
    - protocols
project_type: protocol
publisher:
    name: DigitalGuards
    url: https://github.com/DigitalGuards/QuantaSwap
publishers:
    - DigitalGuards
qrl_environments:
    - testnet
qrl_generations:
    - "2.0"
qrl_relationship: native
qrl_support:
    - generation: "2.0"
      environments:
        - testnet
repositories:
    - id: main
      role: contracts
      url: https://github.com/DigitalGuards/QuantaSwap
      license: GPL-3.0
secondary_categories:
    - finance
source_availability: full
title: QuantaSwap
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
