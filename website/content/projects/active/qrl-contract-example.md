---
aliases:
    - /projects/active/qrl-contract-example/
    - /projects/archived/qrl-contract-example/
availability: live
capabilities:
    - contract-template
    - deployment
categories:
    - developer-experience
data_updated_at: "2026-08-17"
description: A JavaScript reference project for compiling, deploying, and interacting with a Hyperion token contract on QRL 2.0 using @theqrl/web3, including on-chain transfers and off-chain balance queries.
display_status: Beta
features:
    - Hyperion token contract built from ERC20 and IERC20 examples
    - Local contract compilation with @theqrl/hypc
    - Contract deployment from a Dilithium-seeded account with @theqrl/web3
    - On-chain token transfers and read-only balance queries
id: qrl-contract-example
keywords:
    - smart-contract
    - hyperion
    - web3
    - javascript
    - token
links:
    - type: website
      url: https://github.com/theQRL/qrl-contract-example
      primary: true
listed_at: "2026-08-17"
maintainer_records:
    - name: The QRL
      contact: https://github.com/theQRL/qrl-contract-example
maintainers:
    - The QRL
maturity: beta
platforms: []
primary_category: developer-experience
primary_link:
    type: website
    url: https://github.com/theQRL/qrl-contract-example
    primary: true
primary_url: https://github.com/theQRL/qrl-contract-example
project-types:
    - tooling
project_type: tooling
publisher:
    name: The QRL
    url: https://github.com/theQRL/qrl-contract-example
publishers:
    - The QRL
qrl_environments: []
qrl_generations:
    - "2.0"
qrl_relationship: native
qrl_support:
    - generation: "2.0"
      environments: []
repositories:
    - id: main
      role: tooling
      url: https://github.com/theQRL/qrl-contract-example
      license: MIT
secondary_categories: []
source_availability: full
title: QRL Contract Example
url: /projects/qrl-contract-example/
---

QRL Contract Example is an open-source JavaScript reference project from The
QRL for getting started with smart contracts on QRL 2.0. Its upstream README
describes it as a Zond smart contract example and walks developers through
configuring a local node provider and testnet-funded Dilithium wallet.

The included scripts compile a sample Hyperion token contract, deploy it with
@theqrl/web3, submit an on-chain token transfer, and query a token balance
without creating a transaction. The repository provides a compact starting
point for learning the QRL 2.0 contract workflow.
