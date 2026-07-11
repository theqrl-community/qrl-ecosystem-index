---
aliases:
    - /projects/active/myqrlwallet/
    - /projects/archived/myqrlwallet/
audited: false
audits: []
author: DigitalGuards
categories:
    - wallet
category: wallet
clients:
    - platform: web
      url: https://myqrlwallet.com/
      default: true
    - platform: web wallet
      url: https://qrlwallet.com/
      github: https://github.com/DigitalGuards/myqrlwallet-frontend/
    - platform: android
      url: https://play.google.com/store/apps/details?id=com.chiefdg.myqrlwallet
    - platform: desktop
      url: https://github.com/DigitalGuards/myqrlwallet-desktop/releases
      github: https://github.com/DigitalGuards/myqrlwallet-desktop
    - platform: browser extension
      github: https://github.com/DigitalGuards/myqrlwallet-extension
created: "2026-06-09"
default_client:
    platform: web
    url: https://myqrlwallet.com/
    default: true
default_client_github: https://github.com/DigitalGuards/myqrlwallet-frontend/
default_client_url: https://myqrlwallet.com/
description: 'A family of wallet applications for QRL 2.0: web wallet, Android app, desktop wallet, and browser extension, with post-quantum dApp connectivity over the QRL Connect relay protocol.'
discord: ""
docs: ""
ecosystem_type: application
features:
    - Web wallet for account creation, transactions, tokens, and NFTs on the quantum-resistant ledger
    - Android app and hardened Electron desktop wallet
    - Browser extension (MIT fork of the QRL web3 wallet)
    - dApp connectivity via the QRL Connect relay protocol (ML-KEM-768 + AES-256-GCM, ML-DSA-87 signing)
github: https://github.com/DigitalGuards/myqrlwallet-frontend/
license: none
logos: []
open_source: true
platforms:
    - web
    - android
    - desktop
    - browser extension
project-types:
    - applications
project_type: application
project_url: https://myqrlwallet.com/
status: development
supported_networks:
    - testnet
tags:
    - wallet
title: MyQRLWallet
twitter: https://x.com/DigitalGuards
updated: "2026-07-11"
url: /projects/myqrlwallet/
---

MyQRLWallet is a DigitalGuards family of wallet applications for QRL 2.0.
It started as a browser-based web wallet (qrlwallet.com) for creating
accounts and managing transactions on the quantum-resistant ledger, and has
grown into a suite: an Android app on Google Play, a hardened Electron
desktop wallet with an isolated signer process, and a browser extension.

dApps connect to the wallets through the open-source QRL Connect relay
protocol (ML-KEM-768 key encapsulation with AES-256-GCM transport and
ML-DSA-87 message signing) or through the browser extension via EIP-6963.

