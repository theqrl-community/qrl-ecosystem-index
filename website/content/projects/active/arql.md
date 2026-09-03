---
aliases:
    - /projects/active/arql/
    - /projects/archived/arql/
assets:
    - type: token
      name: USD Coin
      symbol: USDC
      deployment_id: qrl-2-testnet-v2
      identifier: Qadf94bb6e061a9f3b1d54826241eba701d43fb86
      evidence_url: https://zondscan.com/address/Qadf94bb6e061a9f3b1d54826241eba701d43fb86
availability: live
capabilities:
    - payments
    - faucet
categories:
    - interoperability-messaging-data
    - finance
    - assets-tokenization
data_updated_at: "2026-09-02"
deployments:
    - id: qrl-2-testnet-v2
      network: qrl-2-testnet-v2
      operational_state: live
      identifiers:
        - type: token
          value: Qadf94bb6e061a9f3b1d54826241eba701d43fb86
          role: qrc20-usdc
        - type: contract
          value: Q89b8cdf9bddba27ac48938bd82d99f33796bcc05
          role: minter
        - type: contract
          value: Q360b826e5290ae8ebb0b58318c12085cb37d9dd8
          role: sealed-bridge
      evidence:
        - https://github.com/pq-cybarg/arql/blob/main/deployments/qrl-testnet.json
        - https://zondscan.com/address/Qadf94bb6e061a9f3b1d54826241eba701d43fb86
      source_verification: unverified
description: A CCTP-shaped lock-and-mint path that locks native Circle USDC on Arc Testnet and mints a 6-decimal QRC-20 USD Coin on QRL 2.0, with Hyperion contracts, an Iris-style attester API, and a public settlement desk.
display_status: Beta · Testnet
evidence:
    - type: deployment
      url: https://zondscan.com/address/Qadf94bb6e061a9f3b1d54826241eba701d43fb86
      note: QRC-20 USDC on QRL 2.0 Testnet V2
      checked_at: "2026-09-02"
    - type: qrl-relationship
      url: https://github.com/pq-cybarg/arql/blob/main/docs/architecture.md
      note: Public architecture documenting the QRVM Hyperion lock-and-mint path
      checked_at: "2026-09-02"
features:
    - Locks official Arc native USDC; does not mint a second USDC on Arc
    - Hyperion QRC-20 USD Coin (6 decimals) on QRL 2.0 Testnet V2
    - SLH-DSA attestation on Arc and ML-DSA-87 attester on QRVM
    - Public settlement desk and Iris-shaped /v2 message API
id: arql
keywords:
    - usdc
    - qrc20
    - bridge
    - hyperion
    - cctp
last_verified_at: "2026-09-02"
links:
    - type: website
      url: https://github.com/pq-cybarg/arql
    - type: application
      url: https://pq-cybarg.github.io/arql/
      platform: web
      primary: true
    - type: documentation
      url: https://github.com/pq-cybarg/arql/blob/main/docs/architecture.md
listed_at: "2026-09-02"
maintainer_records:
    - name: pq-cybarg
      contact: https://github.com/pq-cybarg
maintainers:
    - pq-cybarg
maintenance: active
maturity: beta
platforms: []
primary_category: interoperability-messaging-data
primary_link:
    type: application
    url: https://pq-cybarg.github.io/arql/
    platform: web
    primary: true
primary_url: https://pq-cybarg.github.io/arql/
project-types:
    - protocols
project_launched_at: "2026-08-31"
project_type: protocol
publisher:
    name: pq-cybarg
    url: https://pq-cybarg.github.io/
publishers:
    - pq-cybarg
qrl_environments:
    - testnet
qrl_generations:
    - "2.0"
qrl_relationship: native
qrl_support:
    - generation: "2.0"
      environments:
        - testnet
      evidence:
        - https://github.com/pq-cybarg/arql
        - https://zondscan.com/address/Qadf94bb6e061a9f3b1d54826241eba701d43fb86
repositories:
    - id: main
      role: contracts
      url: https://github.com/pq-cybarg/arql
      license: NOASSERTION
secondary_categories:
    - finance
    - assets-tokenization
source_availability: full
title: ARQL
url: /projects/arql/
---

ARQL is a CCTP-shaped bridge that locks native Circle USDC on Arc Testnet
and mints a 6-decimal QRC-20 named USD Coin on QRL 2.0. QRL is not a Circle
CCTP domain, so ARQL does not call Circle TokenMessengerV2 toward QRL.

Arc-side contracts lock official native USDC. QRVM-side contracts are
Hyperion and mint or burn the QRC-20. Arc attestation uses SLH-DSA-SHA2-128s;
QRL receive uses ML-DSA-87 `msg.sender`. A public static desk is published at
pq-cybarg.github.io/arql. Current Testnet V2 deployments use 20-byte Q
addresses; the repository documents a later 64-byte address reset.

This listing is informational and is not an endorsement, audit, or claim
that the token is Circle-issued native USDC on QRL.
