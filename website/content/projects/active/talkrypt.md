---
aliases:
    - /projects/active/talkrypt/
    - /projects/archived/talkrypt/
availability: live
capabilities:
    - cli
    - key-management
categories:
    - interoperability-messaging-data
    - social-creator-content
    - identity-naming-privacy
data_updated_at: "2026-09-02"
description: Post-quantum P2P chat for side-channel and resilient communications across more than one network type. Public CLI, TUI, and Android builds use ML-KEM-1024 and ML-DSA-87. QRL 2.0 account login is planned and is not in the published tree yet.
display_status: Prototype
features:
    - Post-quantum P2P chat with P2P, hub, and hybrid topologies
    - Cross-network transports for side-channel and resilient communications
    - ML-KEM-1024 and ML-DSA-87 double ratchet; Tor/Arti among the transports
    - Planned QRL 2.0 account login; not present in current public releases
id: talkrypt
keywords:
    - p2p
    - chat
    - tor
    - side-channel
    - ml-dsa
last_verified_at: "2026-09-02"
links:
    - type: website
      url: https://github.com/pq-cybarg/talkrypt
    - type: application
      url: https://pq-cybarg.github.io/talkrypt/
      platform: web
      primary: true
listed_at: "2026-09-02"
maintainer_records:
    - name: pq-cybarg
      contact: https://github.com/pq-cybarg
maintainers:
    - pq-cybarg
maintenance: active
maturity: prototype
platforms:
    - desktop
    - mobile
primary_category: interoperability-messaging-data
primary_link:
    type: application
    url: https://pq-cybarg.github.io/talkrypt/
    platform: web
    primary: true
primary_url: https://pq-cybarg.github.io/talkrypt/
project-types:
    - applications
project_type: application
publisher:
    name: pq-cybarg
    url: https://pq-cybarg.github.io/
publishers:
    - pq-cybarg
qrl_environments: []
qrl_generations:
    - "2.0"
qrl_relationship: integrated
qrl_support:
    - generation: "2.0"
      environments: []
repositories:
    - id: main
      role: client
      url: https://github.com/pq-cybarg/talkrypt
      license: Apache-2.0
secondary_categories:
    - social-creator-content
    - identity-naming-privacy
source_availability: full
title: talkrypt
url: /projects/talkrypt/
---

talkrypt is intended as a post-quantum P2P messenger that can move across
network types for side-channel and resilient communications, not as a
single-transport onion-service toy. Current public code includes CLI, TUI,
and Android clients. QRL 2.0 login is planned and is not available yet.
