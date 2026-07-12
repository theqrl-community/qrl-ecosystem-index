# Contributing to QRL Ecosystem

Thank you for your interest in contributing your project to the QRL ecosystem index!

This index broadly covers five types of projects.

1. **dApps**: Decentralized, smart-contract based applications with on-chain logic.
2. **Applications**: User-facing client applications such as wallets and explorers.
3. **Infrastructure**: Services that support or extend the QRL network.
4. **Tooling**: Development tools, libraries, and SDKs
5. **Community**: Unofficial community groups, spaces, informational websites, and educational content.

## Requirements

### Soft Requirements

#### Branding 

- To avoid confusion or implied endorsement, submissions should avoid using `QRL` as part of the project name, logo, domain, social handle, token symbol, or primary brand identity.
- If the usage of `QRL` can't be avoided, governance, ownership or non-endorsement should be made clear in the name, for example `QRL Explorer, by Sam`.
- Descriptive usage is allowed where it is factual and non-branded, such as “built on QRL,” “for QRL,” or “supports QRL.”
- Ecosystem-inspired terminology, such as `quanta`, `planck` from the QRL nomenclature, are fine.

### Required Fields

| Field | Type | Description |
|---|---|---|
| `id` | `string` | Unique slug identifier, lowercase, hyphenated (e.g. `my-qrl-dapp`) |
| `name` | `string` | Human-readable project name |
| `project_type` | `enum` | One of: `dapp`, `application`, `infrastructure`, `tooling`, `community` |
| `status` | `enum` | One of: `development`, `production`, `archived` |
| `description` | `string` | Short description, max 280 characters |
| `category` | `enum` | One of: `defi`, `nft`, `wallet`, `explorer`, `infrastructure`, `tooling`, `dao`, `gaming`, `identity`, `oracle`, `bridge`, `social`, `educational`, `news` |
| `tags` | `list[string]` | Freeform tags for filtering/search |
| `author` | `string` | Author name or organization |
| `license` | `string` | SPDX license identifier (e.g. `MIT`, `Apache-2.0`). Use `none` if not applicable |
| `created` | `date` | ISO 8601 date of initial submission (`YYYY-MM-DD`) |
| `updated` | `date` | ISO 8601 date of last meaningful update |

### Optional Fields

| Field | Type | Description |
|---|---|---|
| `url` | `string` | Primary project URL |
| `github` | `string` | GitHub repository URL |
| `discord` | `string` | Discord invite or channel link |
| `twitter` | `string` | Twitter/X handle or URL |
| `docs` | `string` | Documentation URL |
| `logo` | `string` | Relative path to a logo image |
| `screenshots` | `list[object]` | Ordered local screenshots, each with `path` and `caption` |
| `long_description` | `string` | Extended markdown description |
| `features` | `list[string]` | Key features or capabilities of the project |
| `open_source` | `boolean` | True if the source code is public |
| `audited` | `boolean` | True if professionally audited |
| `audits` | `list[object]` | Security audit reports, each with `auditor` and `audit_url` |
| `clients` | `list[object]` | Platform-specific destinations, each with `platform`, optional `url`, optional `github`, and optional `default` |

Use `clients` when one project has separate platform destinations, such as web, iOS, Android, or desktop wallets. Set `default: true` on the client that should be used as the primary external link in project cards and generated indexes. If no client is marked as the default, the site uses the first client with a `url`, then falls back to the top-level `url`.

### Project Type-Specific Nested Blocks

Only one type-specific block should be included per submission, matching the `project_type`. Use `category` and `tags` to provide finer-grained classification within each type — for example, a wallet application would set `project_type: application` with `category: [wallet]`, or a community-run wiki would set `project_type: community` with `category: [educational]`.

| Block Key | Fields | Description |
|---|---|---|
| `dapp` | `network` (enum), `contract_address` (string), `token` (string) | For smart-contract based decentralized applications |
| `application` | `platforms` (list[string]), `supported_networks` (list[string]) | For user-facing client applications such as wallets and explorers |
| `infrastructure` | `supported_networks` (list[string]), `endpoints` (list[string]) | For network services, nodes, indexers, and APIs |
| `tooling` | `languages` (list[string]) | For development tools, libraries, and SDKs |
| `community` | `platforms` (list[string]), `language` (string) | For community groups, social spaces, informational websites, and educational resources |

## Submission Process

1. **Fork** the repository
2. **Copy** `projects/template.yaml` into the appropriate subdirectory:
   - `projects/active/` - Live or in-development projects
   - `projects/archived/` - No longer active projects
3. **Name the file** using your project's `id` field (e.g., `my-project.yaml`)
4. **Fill out** all required fields and any relevant optional fields
5. **Open a Pull Request** using the provided template
6. Automated checks will run (schema validation and linting)
7. A maintainer will review your submission

## Automated Checks

Your PR must pass:

- YAML schema validation
- Filename matches `id` field
- No duplicate `id` values
- Correct subdirectory based on `status`
- Valid date formats (ISO 8601)
- Valid URLs
- Description under 280 characters
- Valid `category` and `status` values
- Logos use local paths (no external URLs)
- Screenshots use local, project-scoped paths and meet count, format, caption, and file-size requirements

## Logos

Logos must be hosted locally in this repository for security reasons. External URLs are not allowed.

### Adding Logos

1. Create a subdirectory under `images/logos/<your-project-id>/`
2. Add your logo files (PNG, SVG, or WebP recommended)
3. Reference them in your YAML using the `logos` array or `logo` shorthand:

```yaml
# Shorthand (single logo)
logo: your-project/icon.png

# Full format (multiple variants)
logos:
  - path: your-project/icon.png
    description: Project icon
  - path: your-project/logo-full.svg
    description: Full-width logo
```
### Requirements

- **Path format**: `images/logos/<project-id>/<filename>`
- **Supported formats**: PNG, SVG, WebP
- **Icon size**: 128x128 or 256x256 recommended
- **No external URLs**: All logo references must be local paths
- **PR includes both YAML and logo files**: Both must be in the same PR

## Screenshots

Screenshots are optional and must be hosted locally in this repository. External URLs are not allowed. When provided, screenshots appear in a full-width gallery on the project page in the same order as the YAML entries.

### Adding Screenshots

1. Create a subdirectory under `images/screenshots/<your-project-id>/`
2. Add between 1 and 10 PNG, JPEG, or WebP images
3. Reference each image with a short caption and place your best screenshot first:

```yaml
screenshots:
  - path: your-project/main-dashboard.webp
    caption: Main dashboard showing current network activity
  - path: your-project/transaction-details.png
    caption: Transaction details and confirmation status
```

### Requirements

- **Path format**: `images/screenshots/<project-id>/<filename>`
- **Supported formats**: lowercase `.png`, `.jpg`, `.jpeg`, or `.webp`
- **File size**: 2 MB maximum per screenshot
- **Count**: 1–10 screenshots; 3–6 is a good range for most projects
- **Captions**: Required, concise, plain text, and no more than 160 characters
- **No external URLs**: All screenshot references must be local paths
- **PR includes both YAML and screenshot files**: Both must be in the same PR

Use current screenshots with populated, readable content. Avoid sensitive information, entire desktop backgrounds, promotional text overlays, or stale interfaces. Desktop captures around or below 2000x1400 pixels are recommended so interface details remain legible when scaled; other aspect ratios, including portrait mobile screenshots, are welcome.

## Review Criteria

Maintainers will review for:

- Schema compliance
- Accurate categorization
- Unique `id` values
- No malicious or misleading entries

Maintainers do **not** endorse or audit projects for security.

## Questions?

Open an issue or reach out on the QRL Discord.

## Examples

### Quanta Swap

```yaml
id: quanta-swap
name: Quanta Swap
project_type: dapp
status: development
description: >
  A decentralized token swap protocol built natively
  on QRL 2.0 smart contracts.
category: defi
tags:
  - swap
  - amm
  - tokens
author: Community
license: MIT
created: 2024-11-01
updated: 2025-01-15

url: https://qrlswap.example.com
github: https://github.com/alicedev/qrl-swap
docs: https://docs.qrlswap.example.com
discord: https://discord.gg/example
twitter: https://twitter.com/qrlswap

audited: true
audits:
  - auditor: Example Security
    audit_url: https://security.example.com/audit/qrl-swap

features:
  - Automated market maker (AMM)
  - Permissionless liquidity pools
  - Post-quantum secure transactions

dapp:
  networks: testnet
  contract_address: "Q0104..."

long_description: |
  QRL Swap is the first AMM protocol designed specifically
  for the QRL 2.0 ecosystem. It leverages the post-quantum
  cryptographic guarantees of the QRL chain to provide
  a swap protocol that is secure against both classical
  and quantum adversaries.
```

### Quanta Wallet

```yaml
id: quanta-wallet
name: Quanta Wallet
project_type: application
status: development
description: >
  A post-quantum secure wallet for storing, sending,
  and receiving QRL tokens with advanced privacy features.
category: wallet
tags:
  - wallet
  - privacy
  - tokens
author: Community
license: MIT
created: 2025-01-01
updated: 2025-05-28

url: https://quanta-wallet.example.com
github: https://github.com/qrl/quanta-wallet
docs: https://docs.quanta-wallet.example.com
discord: https://discord.gg/quanta-wallet
twitter: https://twitter.com/quantawallet

clients:
  - platform: web
    url: https://quanta-wallet.example.com
    github: https://github.com/qrl/quanta-wallet-web
    default: true
  - platform: android
    url: https://play.google.com/store/apps/details?id=com.example.quantawallet
    github: https://github.com/qrl/quanta-wallet-android
  - platform: ios
    url: https://apps.apple.com/app/quanta-wallet/id0000000000
    github: https://github.com/qrl/quanta-wallet-ios

audited: true
audits:
  - auditor: Example Security
    audit_url: https://security.example.com/audit/quanta-wallet

features:
  - Post-quantum secure key generation
  - Multi-signature support
  - Hardware wallet integration
  - Token swap integration
  - Address book with encryption
  - Testnet and mainnet support

application:
  platforms:
    - web
    - android
    - ios
  supported_networks:
    - testnet
    - mainnet

long_description: |
  Quanta Wallet is a secure, user-friendly wallet for the QRL
  ecosystem. Built with post-quantum cryptography from the ground
  up, it ensures your funds are protected against both classical
  and quantum adversaries. Features include multi-signature support,
  hardware wallet integration, and seamless token swaps.
```
