---
aliases:
    - /projects/active/acornswap/
    - /projects/archived/acornswap/
assets:
    - type: token
      name: Acorn V2 LP
      symbol: Acorn V2 LP
      deployment_id: qrl-2-testnet-v2
availability: live
capabilities:
    - dex
categories:
    - finance
data_updated_at: "2026-08-17"
deployments:
    - id: qrl-2-testnet-v2
      network: qrl-2-testnet-v2
      operational_state: live
      identifiers:
        - type: contract
          value: QA19FA6B49A10F1B57E303ECE95eF21acE640d312
      evidence:
        - https://github.com/PhuocNG0308/acorn-swap
      source_verification: unverified
description: A Hyperion port of Uniswap V2 and V3 for the QRL 2.0 testnet. Its deployed V2 contracts provide constant-product pairs and routing, while the concentrated- liquidity V3 contracts are built and tested but not yet deployed.
display_status: Beta · Testnet
features:
    - Deployed constant-product Acorn V2 factory and router
    - Wrapped QRL integration for native-asset trading pairs
    - Concentrated-liquidity Acorn V3 contracts and local test suite
    - Documented QRL address, signature, compiler, and deployment adaptations
id: acornswap
keywords:
    - dex
    - amm
    - liquidity
    - hyperion
links:
    - type: website
      url: https://github.com/PhuocNG0308/acorn-swap
      primary: true
listed_at: "2026-08-17"
maintainer_records:
    - name: PhuocNG0308
      contact: https://github.com/PhuocNG0308/acorn-swap
maintainers:
    - PhuocNG0308
maturity: beta
platforms: []
primary_category: finance
primary_link:
    type: website
    url: https://github.com/PhuocNG0308/acorn-swap
    primary: true
primary_url: https://github.com/PhuocNG0308/acorn-swap
project-types:
    - protocols
project_type: protocol
publisher:
    name: PhuocNG0308
    url: https://github.com/PhuocNG0308/acorn-swap
publishers:
    - PhuocNG0308
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
      url: https://github.com/PhuocNG0308/acorn-swap
      license: NOASSERTION
secondary_categories: []
source_availability: full
title: AcornSwap
url: /projects/acornswap/
---
