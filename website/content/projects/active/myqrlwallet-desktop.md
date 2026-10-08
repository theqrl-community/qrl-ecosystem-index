---
aliases:
    - /projects/active/myqrlwallet-desktop/
    - /projects/archived/myqrlwallet-desktop/
availability: live
capabilities:
    - wallet
    - wallet-connector
categories:
    - security-custody-account-management
    - assets-tokenization
data_updated_at: "2026-10-05"
description: Hardened Electron desktop wallet for QRL 2.0 on Windows and Linux. Keys live only in an isolated signer process, with a native unlock window, lock, and wipe, and dApp connectivity over QRL Connect.
display_status: Beta · Testnet
features:
    - Four-process Electron design where an isolated signer utility process is the only holder of key material
    - Native unlock window, Lock, and a trusted-confirm Remove wallet that wipes the device
    - Same wallet UI as the web wallet, built from the open-source frontend at release time
    - dApp connectivity through the QRL Connect relay protocol (ML-KEM-768 + AES-256-GCM, ML-DSA-87 signing)
    - Windows and Linux builds published on GitHub releases (binaries are unsigned)
gallery:
    - type: image
      path: myqrlwallet-desktop/onboarding.jpg
      caption: Desktop wallet on first launch, before an account is created or imported.
id: myqrlwallet-desktop
keywords:
    - wallet
    - electron
    - windows
    - linux
    - ml-dsa-87
last_release_at: "2026-09-28"
links:
    - type: application
      url: https://github.com/DigitalGuards/myqrlwallet-desktop/releases/latest
      label: Download (GitHub releases)
      platform: desktop
      primary: true
    - type: website
      url: https://myqrlwallet.com/
    - type: security
      url: https://github.com/DigitalGuards/myqrlwallet-desktop/blob/main/THREAT_MODEL.md
      label: Threat model
    - type: social
      url: https://x.com/DigitalGuards
      platform: x
listed_at: "2026-10-05"
logos:
    - path: myqrlwallet-desktop/icon.png
      description: MyQRLWallet logo
maintainer_records:
    - name: DigitalGuards
      contact: https://github.com/DigitalGuards/myqrlwallet-desktop
maintainers:
    - DigitalGuards
maintenance: active
maturity: beta
platforms:
    - desktop
primary_category: security-custody-account-management
primary_link:
    type: application
    url: https://github.com/DigitalGuards/myqrlwallet-desktop/releases/latest
    label: Download (GitHub releases)
    platform: desktop
    primary: true
primary_url: https://github.com/DigitalGuards/myqrlwallet-desktop/releases/latest
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
repositories:
    - id: main
      role: client
      url: https://github.com/DigitalGuards/myqrlwallet-desktop
      license: MIT
secondary_categories:
    - assets-tokenization
source_availability: full
title: MyQRLWallet Desktop
url: /projects/myqrlwallet-desktop/
---

MyQRLWallet Desktop packages the MyQRLWallet web wallet as a hardened
Electron application. The renderer never holds a seed: an isolated signer
utility process is the sole key holder, the main process brokers requests,
and the renderer runs with context isolation under a strict CSP. A native
unlock window gates access, Lock clears the signer, and Remove wallet wipes
local state after a trusted confirmation. The threat model and invariants
are documented in the repository.

The application talks to the QRL 2.0 testnet through the same RPC proxy as
the web wallet and connects to dApps over the QRL Connect relay protocol.
Releases for Windows and Linux are published on GitHub; the binaries are not
code-signed, so operating systems show an unknown-publisher warning.
