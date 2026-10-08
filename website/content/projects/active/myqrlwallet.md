---
aliases:
    - /projects/active/myqrlwallet/
    - /projects/archived/myqrlwallet/
availability: live
capabilities:
    - wallet
    - wallet-connector
categories:
    - security-custody-account-management
    - payments-commerce
    - assets-tokenization
data_updated_at: "2026-10-05"
description: 'A family of wallet applications for QRL 2.0: web wallet, Android app, desktop wallet, and browser extension, with post-quantum dApp connectivity over the QRL Connect relay protocol.'
display_status: Beta · Testnet
features:
    - Web wallet for account creation, transactions, tokens, and NFTs on the quantum-resistant ledger
    - Android app and hardened Electron desktop wallet
    - Browser extension (MIT fork of the QRL web3 wallet)
    - dApp connectivity via the QRL Connect relay protocol (ML-KEM-768 + AES-256-GCM, ML-DSA-87 signing)
gallery:
    - type: image
      path: myqrlwallet/landing.jpg
      caption: MyQRLWallet landing page with the wallet suite for web, desktop, Android, and browser.
    - type: image
      path: myqrlwallet/platforms.jpg
      caption: Platform overview listing the web wallet, desktop app, mobile app, and browser extension.
    - type: image
      path: myqrlwallet/self-custody.jpg
      caption: Self-custody, ML-DSA-87 signing, and encrypted dApp connections explained on the landing page.
    - type: image
      path: myqrlwallet/ecosystem.jpg
      caption: Ecosystem section linking the wallet to QuantaSwap and QuantaPool.
id: myqrlwallet
keywords:
    - wallet
links:
    - type: website
      url: https://myqrlwallet.com/
    - type: application
      url: https://myqrlwallet.com/
      platform: web
      primary: true
    - type: application
      url: https://qrlwallet.com/
      platform: web wallet
    - type: application
      url: https://play.google.com/store/apps/details?id=com.chiefdg.myqrlwallet
      platform: android
    - type: application
      url: https://github.com/DigitalGuards/myqrlwallet-desktop/releases
      platform: desktop
    - type: app-store
      url: https://chromewebstore.google.com/detail/myqrlwallet/cafpjgccpmfjefdfeldlhgmgbmpibdgg
      label: Chrome Web Store
      platform: browser extension
    - type: social
      url: https://x.com/DigitalGuards
listed_at: "2026-06-09"
logos:
    - path: myqrlwallet/icon.png
      description: MyQRLWallet logo
maintainer_records:
    - name: DigitalGuards
      contact: https://github.com/DigitalGuards/myqrlwallet-frontend/
maintainers:
    - DigitalGuards
maturity: beta
platforms:
    - mobile
    - desktop
    - browser-extension
    - web
primary_category: security-custody-account-management
primary_link:
    type: application
    url: https://myqrlwallet.com/
    platform: web
    primary: true
primary_url: https://myqrlwallet.com/
project-types:
    - applications
project_type: application
publisher:
    name: DigitalGuards
    url: https://github.com/DigitalGuards/myqrlwallet-frontend/
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
      role: client
      url: https://github.com/DigitalGuards/myqrlwallet-frontend/
      license: MIT
    - id: desktop
      role: client
      url: https://github.com/DigitalGuards/myqrlwallet-desktop
      license: MIT
    - id: browser-extension
      role: client
      url: https://github.com/DigitalGuards/myqrlwallet-extension
      license: MIT
secondary_categories:
    - payments-commerce
    - assets-tokenization
source_availability: full
title: MyQRLWallet
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
