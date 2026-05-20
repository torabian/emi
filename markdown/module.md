# Module

## Contents

- [DTOs](#dtos) — 1
- [Actions](#actions) — 11

## DTOs

### `FirewallRuleDto`

> One firewall rule belonging to a zone. The rule's expression is the predicate evaluated at the edge for every request; when it matches, the configured action is applied. Rules are evaluated in priority order (lowest priority number first).

| Field | Type | Description |
|---|---|---|
| `uniqueId` | `string` |  |
| `workspaceId` | `string` |  |
| `userId` | `string` |  |
| `zoneId` | `string` | Unique id of the zone this rule belongs to. |
| `description` | `string` | Human-readable description, e.g. "Block bad bots". Also used as the display name in dashboards. |
| `expression` | `string` | The match expression (firewall predicate language), e.g. (http.host eq "api.example.com" and ip.src in {1.2.3.4}). A bare "*" means "match everything". |
| `action` | `string` | What to do when the expression matches. |
| `priority` | `int` | Evaluation order. Lower numbers are evaluated first. Ties are broken by created_at ascending. |
| `enabled` | `bool` | Whether the rule is active. Disabled rules are skipped. |

## Actions

### `QueryFirewallRulesAction`

> Lists firewall rules for the calling workspace. Pass ?zone=<zoneId> to restrict to a single zone.

| | |
|---|---|
| **Method** | `GET` |
| **URL** | `/firewall-rules/query` |
| **CLI** | `rules-query` |
| **CLI short** | `q` |

#### Response

| | |
|---|---|
| **DTO** | `FirewallRuleDto` |
| **Envelope** | `GResponse` |

### `GetFirewallRuleAction`

> Returns a single firewall rule by its unique id.

| | |
|---|---|
| **Method** | `GET` |
| **URL** | `/firewall-rule/request/:rule` |
| **CLI** | `rules-get` |

#### Response

| | |
|---|---|
| **DTO** | `FirewallRule` |
| **Envelope** | `GResponse` |

### `CreateFirewallRuleAction`

> Creates a new firewall rule.

| | |
|---|---|
| **Method** | `POST` |
| **URL** | `/firewall-rule/create` |
| **CLI** | `rules-create` |
| **CLI short** | `c` |

#### Request

| | |
|---|---|
| **DTO** | `FirewallRuleDto` |

#### Response

| | |
|---|---|
| **Envelope** | `GResponse` |

| Field | Type | Description |
|---|---|---|
| `uniqueId` | `string` |  |

### `UnlinkFirewallRuleAction`

> Removes a firewall rule by its unique id.

| | |
|---|---|
| **Method** | `POST` |
| **URL** | `/firewall-rule/unlink` |
| **CLI** | `rules-unlink` |
| **CLI short** | `u` |

#### Request

| Field | Type | Description |
|---|---|---|
| `uniqueId` | `string` |  |

#### Response

| | |
|---|---|
| **Envelope** | `GResponse` |

| Field | Type | Description |
|---|---|---|
| `executed` | `bool` |  |

### `SetFirewallRuleEnabledAction`

> Toggle a single firewall rule on or off without re-creating it.

| | |
|---|---|
| **Method** | `PATCH` |
| **URL** | `/firewall-rule/enabled` |
| **CLI** | `rules-set-enabled` |

#### Request

| Field | Type | Description |
|---|---|---|
| `uniqueId` | `string` |  |
| `enabled` | `bool` | Target state for the rule. |

#### Response

| | |
|---|---|
| **Envelope** | `GResponse` |

| Field | Type | Description |
|---|---|---|
| `enabled` | `bool` |  |

### `GetZoneFirewallConfigurationAction`

> Returns the per-zone firewall settings (security level, challenge TTL, and the simple feature toggles).

| | |
|---|---|
| **Method** | `GET` |
| **URL** | `/zone/get-firewall-config/:zone` |
| **CLI** | `get` |

#### Response

| Field | Type | Description |
|---|---|---|
| `securityLevel` | `string` | Sensitivity threshold for proactive challenges. One of off, low, medium, high, i_am_under_attack. Empty string when the zone has no override yet — treat as "medium". |
| `challengeTtl` | `int` | How long (seconds) a successful challenge solution is honored before the visitor is challenged again. Zero when no override. |
| `browserIntegrityCheckEnabled` | `bool` | Whether the edge rejects requests with invalid headers or user-agents that look like known scrapers. |
| `hotlinkProtectionEnabled` | `bool` | Whether the edge blocks cross-origin embeds of this zone's images (referrer-based). |
| `emailObfuscationEnabled` | `bool` | Whether mailto:/email-looking strings in HTML are rewritten on the fly to slow down scrapers. |

### `SetBrowserIntegrityCheckAction`

> Block requests whose headers or user-agent look invalid or match known bad-bot signatures.

| | |
|---|---|
| **Method** | `PATCH` |
| **URL** | `/zone/browser-integrity-check` |
| **CLI** | `set-browser-integrity-check` |

#### Request

| Field | Type | Description |
|---|---|---|
| `enabled` | `bool` | Enable or disable the browser integrity check |
| `zone` | `string` | The zone id which the configuration will be assigned to. |

#### Response

| | |
|---|---|
| **Envelope** | `GResponse` |

| Field | Type | Description |
|---|---|---|
| `enabled` | `bool` |  |

### `SetHotlinkProtectionAction`

> Prevent other sites from embedding this zone's images by checking the Referer header.

| | |
|---|---|
| **Method** | `PATCH` |
| **URL** | `/zone/hotlink-protection` |
| **CLI** | `set-hotlink-protection` |

#### Request

| Field | Type | Description |
|---|---|---|
| `enabled` | `bool` | Enable or disable hotlink protection |
| `zone` | `string` | The zone id which the configuration will be assigned to. |

#### Response

| | |
|---|---|
| **Envelope** | `GResponse` |

| Field | Type | Description |
|---|---|---|
| `enabled` | `bool` |  |

### `SetEmailObfuscationAction`

> Rewrite plain-text email addresses in HTML to slow down scrapers.

| | |
|---|---|
| **Method** | `PATCH` |
| **URL** | `/zone/email-obfuscation` |
| **CLI** | `set-email-obfuscation` |

#### Request

| Field | Type | Description |
|---|---|---|
| `enabled` | `bool` | Enable or disable email obfuscation |
| `zone` | `string` | The zone id which the configuration will be assigned to. |

#### Response

| | |
|---|---|
| **Envelope** | `GResponse` |

| Field | Type | Description |
|---|---|---|
| `enabled` | `bool` |  |

### `SetChallengeTtlAction`

> Set how long (seconds) a successful challenge solution is remembered before a visitor is challenged again for the same zone.

| | |
|---|---|
| **Method** | `PATCH` |
| **URL** | `/zone/challenge-ttl` |
| **CLI** | `set-challenge-ttl` |

#### Request

| Field | Type | Description |
|---|---|---|
| `amount` | `int` | Challenge TTL in seconds. Common values include 1800 (30m), 3600 (1h), 86400 (1d). |
| `zone` | `string` | The zone id which the configuration will be assigned to. |

#### Response

| | |
|---|---|
| **Envelope** | `GResponse` |

| Field | Type | Description |
|---|---|---|
| `amount` | `int` |  |

### `SetSecurityLevelAction`

> Choose how aggressively the edge proactively challenges visitors based on threat scores.

| | |
|---|---|
| **Method** | `PATCH` |
| **URL** | `/zone/security-level` |
| **CLI** | `set-security-level` |

#### Request

| Field | Type | Description |
|---|---|---|
| `zone` | `string` | The zone id which the configuration will be assigned to. |
| `level` | `string` | Sensitivity level for proactive challenges. |

#### Response

| | |
|---|---|
| **Envelope** | `GResponse` |

| Field | Type | Description |
|---|---|---|
| `level` | `string` |  |

