---
aliases:
    - /projects/active/qrlwallet-connect/
    - /projects/archived/qrlwallet-connect/
audited: false
audits: []
author: DigitalGuards
categories:
    - tooling
category: tooling
clients: []
created: "2026-06-11"
default_client_github: https://github.com/DigitalGuards/myqrlwallet-connect
default_client_url: https://www.npmjs.com/package/@qrlwallet/connect
description: A TypeScript SDK (@qrlwallet/connect) for connecting dApps to MyQRLWallet over a post-quantum encrypted relay, with QR code and deep-link pairing, an EIP-1193 style provider, and persistent sessions.
discord: ""
docs: https://github.com/DigitalGuards/myqrlwallet-connect/blob/main/docs/JSON-RPC-REFERENCE.md
ecosystem_type: tooling
features:
    - ML-KEM-768 key encapsulation with AES-256-GCM encrypted transport over a Socket.IO relay
    - QR code pairing on desktop and deep-link pairing on mobile
    - EIP-1193 style provider with EIP-6963 announcement, coexisting with browser extension wallets in dApp wallet pickers
    - Session persistence with automatic reconnect and connection liveness probing
    - Post-quantum ML-DSA-87 message and typed-data signing requests (qrl_signMessage, qrl_signTypedData)
github: https://github.com/DigitalGuards/myqrlwallet-connect
languages:
    - typescript
license: MIT
logos: []
open_source: true
project-types:
    - tooling
project_type: tooling
project_url: https://www.npmjs.com/package/@qrlwallet/connect
status: development
tags:
    - sdk
    - dapp
    - post-quantum
title: QRL Wallet Connect
twitter: ""
updated: "2026-06-11"
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

