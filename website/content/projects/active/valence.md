---
aliases:
    - /projects/active/valence/
    - /projects/archived/valence/
availability: limited
capabilities:
    - notarization
categories:
    - utility-storage-compute
data_updated_at: "2026-09-02"
deployments:
    - id: qrl-2-local-private
      network: qrl-2-local-private
      operational_state: limited
      identifiers: []
      evidence: []
      source_verification: unverified
description: 'Erasure-coded decentralized storage network using a custom multi-dimensional Reed-Solomon design (ND-encoding): slivers, per-dimension repair, and Merkle commitments. Prototype on a private 64-byte QRL 2.0 localnet; no public contract address.'
display_status: Limited
features:
    - Multi-dimensional systematic Reed-Solomon erasure coding
    - Independent reconstruction paths per coding dimension
    - Fiber-slivers with per-dimension Merkle proofs and self-healing repair
    - Verifiable sliver-to-node assignment; no public network identifiers
id: valence
keywords:
    - erasure-coding
    - reed-solomon
    - p2p
    - storage
    - sliver
last_verified_at: "2026-09-02"
links:
    - type: website
      url: https://pq-cybarg.github.io/
      primary: true
listed_at: "2026-09-02"
maintainer_records:
    - name: pq-cybarg
      contact: https://github.com/pq-cybarg
maintainers:
    - pq-cybarg
maintenance: active
maturity: prototype
platforms: []
primary_category: utility-storage-compute
primary_link:
    type: website
    url: https://pq-cybarg.github.io/
    primary: true
primary_url: https://pq-cybarg.github.io/
project-types:
    - protocols
project_type: protocol
publisher:
    name: pq-cybarg
    url: https://pq-cybarg.github.io/
publishers:
    - pq-cybarg
qrl_environments:
    - local-private
qrl_generations:
    - "2.0"
qrl_relationship: native
qrl_support:
    - generation: "2.0"
      environments:
        - local-private
repositories: []
secondary_categories: []
source_availability: unavailable
title: Valence
url: /projects/valence/
---

Valence is an unpublished storage-network prototype. The encoding is a
custom N-dimensional Reed-Solomon construction rather than a single-axis
code. Discus can use it as a substrate; Valence is not a Discus-only
companion node.
