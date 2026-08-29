---
aliases:
    - /projects/active/qrl-ledger-wallet/
    - /projects/archived/qrl-ledger-wallet/
availability: live
capabilities:
    - wallet
    - hardware-wallet
    - key-management
categories:
    - security-custody-account-management
    - payments-commerce
data_updated_at: "2026-08-25"
description: The QRL application for Ledger Nano hardware wallets, keeping XMSS keys offline while supporting QRL transfers, messages, multiple address trees, OTS tracking, and optional passphrase-protected address spaces.
display_status: Stable · Mainnet
features:
    - Support for Ledger Nano X and Nano S Plus devices
    - XMSS private keys retained on the hardware secure element
    - Dual address-tree support with OTS index tracking
    - Native QRL transfers and message transactions
gallery:
    - type: image
      path: qrl-ledger-wallet/qrl-ready.jpg
      caption: Ledger Nano displaying QRL READY with 256 XMSS one-time signature keys remaining.
    - type: image
      path: qrl-ledger-wallet/ledger-transfer.png
      caption: Ledger Nano reviewing a 13 QRL transfer amount on-device before signing.
    - type: image
      path: qrl-ledger-wallet/address-verification.png
      caption: Ledger Nano verifying a QRL receiving address directly on the hardware display.
id: qrl-ledger-wallet
keywords:
    - wallet
    - hardware-wallet
    - ledger
    - xmss
    - signing
links:
    - type: website
      url: https://docs.theqrl.org/use/wallet/ledger/overview/
      primary: true
listed_at: "2026-08-17"
maintainer_records:
    - name: The QRL
      contact: https://github.com/theQRL/ledger-qrl-app
maintainers:
    - The QRL
maturity: stable
platforms: []
primary_category: security-custody-account-management
primary_link:
    type: website
    url: https://docs.theqrl.org/use/wallet/ledger/overview/
    primary: true
primary_url: https://docs.theqrl.org/use/wallet/ledger/overview/
project-types:
    - applications
project_type: application
publisher:
    name: The QRL
    url: https://github.com/theQRL/ledger-qrl-app
publishers:
    - The QRL
qrl_environments:
    - mainnet
qrl_generations:
    - 1.x
qrl_relationship: native
qrl_support:
    - generation: 1.x
      environments:
        - mainnet
repositories:
    - id: main
      role: client
      url: https://github.com/theQRL/ledger-qrl-app
      license: Apache-2.0
secondary_categories:
    - payments-commerce
source_availability: full
title: QRL Ledger Wallet
url: /projects/qrl-ledger-wallet/
---
