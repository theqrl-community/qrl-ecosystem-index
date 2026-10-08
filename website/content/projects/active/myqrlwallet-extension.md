---
aliases:
    - /projects/active/myqrlwallet-extension/
    - /projects/archived/myqrlwallet-extension/
availability: live
capabilities:
    - wallet
    - wallet-connector
categories:
    - security-custody-account-management
    - assets-tokenization
data_updated_at: "2026-10-05"
description: Chrome extension wallet for QRL 2.0. Self-custody accounts signed with ML-DSA-87, Quanta and QRC-20 transfers, NFTs, an address book, and dApp connections through EIP-6963.
display_status: Beta · Testnet
features:
    - Create a wallet or import one from a recovery phrase, a hex seed, or an encrypted MyQRLWallet wallet file
    - Send and receive Quanta and QRC-20 tokens, view NFTs grouped by collection, keep an address book
    - Transaction history with links to the ZondScan explorer and a notification when a transaction confirms
    - dApp provider announced over EIP-6963, so it coexists with the QRL web3 wallet in dApp wallet pickers
    - Popup and side panel modes, MV3, keys encrypted on the device and locked when the service worker stops
gallery:
    - type: image
      path: myqrlwallet-extension/home.png
      caption: Popup home with the active account balance, Send, History and Receive actions, tokens and NFT collections.
    - type: image
      path: myqrlwallet-extension/welcome.jpg
      caption: Onboarding in tab mode, the first screen before creating or importing an account.
    - type: image
      path: myqrlwallet-extension/import-account.jpg
      caption: Import an existing account from a mnemonic, a hex seed, or an encrypted wallet file.
    - type: image
      path: myqrlwallet-extension/receive.png
      caption: Receive screen showing the account address as a QR code.
    - type: image
      path: myqrlwallet-extension/chrome-web-store.jpg
      caption: The extension listing on the Chrome Web Store.
id: myqrlwallet-extension
keywords:
    - wallet
    - chrome-extension
    - eip-6963
    - ml-dsa-87
last_release_at: "2026-10-01"
links:
    - type: app-store
      url: https://chromewebstore.google.com/detail/myqrlwallet/cafpjgccpmfjefdfeldlhgmgbmpibdgg
      label: Chrome Web Store
      platform: chrome
      primary: true
    - type: application
      url: https://github.com/DigitalGuards/myqrlwallet-extension/releases/latest
      label: Release zip (GitHub)
      platform: browser extension
    - type: website
      url: https://myqrlwallet.com/
    - type: social
      url: https://x.com/DigitalGuards
      platform: x
listed_at: "2026-10-05"
logos:
    - path: myqrlwallet-extension/icon.png
      description: MyQRLWallet logo
maintainer_records:
    - name: DigitalGuards
      contact: https://github.com/DigitalGuards/myqrlwallet-extension
maintainers:
    - DigitalGuards
maintenance: active
maturity: beta
platforms:
    - browser-extension
primary_category: security-custody-account-management
primary_link:
    type: app-store
    url: https://chromewebstore.google.com/detail/myqrlwallet/cafpjgccpmfjefdfeldlhgmgbmpibdgg
    label: Chrome Web Store
    platform: chrome
    primary: true
primary_url: https://chromewebstore.google.com/detail/myqrlwallet/cafpjgccpmfjefdfeldlhgmgbmpibdgg
project-types:
    - applications
project_type: application
publisher:
    name: DigitalGuards
    url: https://myqrlwallet.com/
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
relationships:
    - type: part-of
      project_id: myqrlwallet
    - type: fork-of
      project_id: qrl-web3-wallet
repositories:
    - id: main
      role: client
      url: https://github.com/DigitalGuards/myqrlwallet-extension
      license: MIT
secondary_categories:
    - assets-tokenization
source_availability: full
title: MyQRLWallet Extension
url: /projects/myqrlwallet-extension/
---

MyQRLWallet Extension is the browser member of the MyQRLWallet family. It is
an MIT-licensed fork of the QRL web3 wallet (qrl-web3-wallet v0.3.0),
rebranded and extended with hex seed and wallet-file import so the same
recovery material works across the MyQRLWallet web, desktop and mobile apps.

Accounts are secured with ML-DSA-87 signatures and the keystore is encrypted
on the device with a password. dApps reach the wallet through the EIP-6963
provider announcement under the rdns com.qrlwallet.extension, so it can be
installed alongside the QRL Foundation extension without the two answering
the same requests. It ships on the Chrome Web Store, with the release zip
also attached to each GitHub release.
