---
aliases:
    - /projects/active/qrlwallet-connect/
    - /projects/archived/qrlwallet-connect/
availability: live
capabilities:
    - sdk
    - wallet-connector
categories:
    - developer-experience
    - security-custody-account-management
data_updated_at: "2026-06-11"
description: A TypeScript SDK (@qrlwallet/connect) for connecting dApps to MyQRLWallet over a post-quantum encrypted relay, with QR code and deep-link pairing, an EIP-1193 style provider, and persistent sessions.
display_status: Beta
features:
    - ML-KEM-768 key encapsulation with AES-256-GCM encrypted transport over a Socket.IO relay
    - QR code pairing on desktop and deep-link pairing on mobile
    - EIP-1193 style provider with EIP-6963 announcement, coexisting with browser extension wallets in dApp wallet pickers
    - Session persistence with automatic reconnect and connection liveness probing
    - Post-quantum ML-DSA-87 message and typed-data signing requests (qrl_signMessage, qrl_signTypedData)
id: qrlwallet-connect
keywords:
    - sdk
    - dapp
    - post-quantum
links:
    - type: package
      url: https://www.npmjs.com/package/@qrlwallet/connect
      primary: true
    - type: documentation
      url: https://github.com/DigitalGuards/myqrlwallet-connect/blob/main/docs/JSON-RPC-REFERENCE.md
listed_at: "2026-06-11"
maintainer_records:
    - name: DigitalGuards
      contact: https://github.com/DigitalGuards/myqrlwallet-connect
maintainers:
    - DigitalGuards
maturity: beta
platforms: []
primary_category: developer-experience
primary_link:
    type: package
    url: https://www.npmjs.com/package/@qrlwallet/connect
    primary: true
primary_url: https://www.npmjs.com/package/@qrlwallet/connect
project-types:
    - tooling
project_type: tooling
publisher:
    name: DigitalGuards
    url: https://github.com/DigitalGuards/myqrlwallet-connect
publishers:
    - DigitalGuards
qrl_environments: []
qrl_generations:
    - "2.0"
qrl_relationship: native
qrl_support:
    - generation: "2.0"
      environments: []
repositories:
    - id: main
      role: sdk
      url: https://github.com/DigitalGuards/myqrlwallet-connect
      license: MIT
secondary_categories:
    - security-custody-account-management
source_availability: full
title: QRL Wallet Connect
url: /projects/qrlwallet-connect/
---

QRL Wallet Connect is an open-source TypeScript SDK, published on npm as
@qrlwallet/connect, that lets dApps connect to the MyQRLWallet web and
mobile wallet. Pairing starts from a QR code on desktop or a deep link on
mobile, and traffic between the dApp and the wallet is end-to-end encrypted
with ML-KEM-768 key encapsulation and AES-256-GCM over a Socket.IO relay.

The SDK exposes an EIP-1193 style provider with EIP-6963 announcement, so
it coexists with browser extension wallets in dApp wallet pickers. Sessions
persist across page reloads and app relaunches with automatic reconnection.
A live integration example is hosted at zondscan.com/dapp-example.
