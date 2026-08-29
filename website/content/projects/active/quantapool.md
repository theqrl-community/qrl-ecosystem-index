---
aliases:
    - /projects/active/quantapool/
    - /projects/archived/quantapool/
assets:
    - type: token
      name: stQRL
      symbol: stQRL
      deployment_id: qrl-2-testnet-v2
availability: live
capabilities:
    - liquid-staking
    - staking
categories:
    - finance
    - network-operations
    - assets-tokenization
data_updated_at: "2026-08-25"
deployments:
    - id: qrl-2-testnet-v2
      network: qrl-2-testnet-v2
      operational_state: live
      identifiers:
        - type: contract
          value: Q109d7C528a67b80eb638D4C85e7C4545ef9Bb9aC
      evidence:
        - https://github.com/DigitalGuards/QuantaPool
      source_verification: unverified
description: A decentralized QRL liquid staking protocol for staking QRL, receiving stQRL, and earning validator rewards while staying liquid.
display_status: Beta · Testnet
features:
    - Liquid staking for QRL through stQRL pool shares
    - Validator rewards reflected through the stQRL to QRL exchange rate
    - Withdrawal request and claim flow with protocol-enforced delay
    - Public pool, validator, reward, and contract statistics
    - QRL Wallet browser extension connectivity
gallery:
    - type: image
      path: quantapool/staking.jpg
      caption: Liquid-staking screen showing QRL input, stQRL output, exchange rate, minimum deposit, and protocol status.
    - type: image
      path: quantapool/protocol-stats.jpg
      caption: Protocol statistics for the staking pool, validators, rewards, withdrawal reserve, and QRL 2.0 testnet contracts.
    - type: image
      path: quantapool/how-it-works.jpg
      caption: Plain-language guide explaining pooled validators, stQRL shares, and reward-driven exchange-rate growth.
id: quantapool
keywords:
    - liquid-staking
    - staking
    - stqrl
links:
    - type: website
      url: https://quantapool.io/
    - type: application
      url: https://quantapool.io/
      platform: web
      primary: true
    - type: documentation
      url: https://quantapool.io/how-it-works
listed_at: "2026-06-10"
logos:
    - path: quantapool/icon.svg
      description: QuantaPool logo
maintainer_records:
    - name: DigitalGuards
      contact: https://github.com/DigitalGuards/QuantaPool
maintainers:
    - DigitalGuards
maturity: beta
platforms: []
primary_category: finance
primary_link:
    type: application
    url: https://quantapool.io/
    platform: web
    primary: true
primary_url: https://quantapool.io/
project-types:
    - protocols
project_type: protocol
publisher:
    name: DigitalGuards
    url: https://github.com/DigitalGuards/QuantaPool
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
      url: https://github.com/DigitalGuards/QuantaPool
      license: GPL-3.0
secondary_categories:
    - network-operations
    - assets-tokenization
source_availability: full
title: QuantaPool
url: /projects/quantapool/
---

QuantaPool is an open-source liquid staking dApp for QRL. Users stake QRL
into a shared pool and receive stQRL, a token representing their share of the
pool. As validators earn rewards, the stQRL to QRL exchange rate increases.

The application describes a protocol that pools deposits toward QRL
validators, exposes pool and validator statistics, supports withdrawal
requests and claims, and connects through the QRL Wallet browser extension.
The deployed web app defaults to the QRL 2.0 Testnet and links to the public
DigitalGuards QuantaPool repository.
