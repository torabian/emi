// Predefined definitions shown in the playground's file explorer.
//
// To add an example, append an entry to `examples` below - nothing else is
// needed. `kind: "yaml"` examples are listed for every compiler target except
// QueryPredict(SQL); `kind: "sql"` examples only for that target. The first
// example of each kind is the default. Selecting an example switches the
// playground to the example's `target` compiler and `tags` options (the tag names
// and descriptions themselves come from the wasm compiler, see getCompilerTags).
// The Allegro SDK definitions live in the repo's examples/allegro-sdk and are pulled
// in as raw text, so the playground always shows the real files.
import allegroOfferManagement from "../../examples/allegro-sdk/definitions/offer/offer-management.emi.yml?raw";
import allegroOfferTranslations from "../../examples/allegro-sdk/definitions/offer/offer-translations.emi.yml?raw";
import allegroUserOfferInformation from "../../examples/allegro-sdk/definitions/offer/user-offer-information.emi.yml?raw";

export type ExampleKind = "yaml" | "sql";

export interface Example {
  /** Unique, stable id (also used as the editor's model path). */
  id: string;
  /** File name shown in the explorer. */
  label: string;
  kind: ExampleKind;
  /** Compiler (playground target) this example is meant to be compiled with. */
  target: string;
  /** Compiler options (--tags) the example is meant to be compiled with. */
  tags: string[];
  /** Options to apply instead of `tags` when the example is viewed with another compiler. */
  tagsByTarget?: Record<string, string[]>;
  content: string;
}

export const examples: Example[] = [
  {
    id: "definition",
    label: "definition.yml",
    kind: "yaml",
    target: "jsGenModule",
    tags: ["typescript", "react", "nestjs", "no-sdk", "no-package"],
    content: `name: sampleModule
entities:
  - name: entity1
    fields:
    - name: field1
      type: string
    - name: users
      type: array
      fields:
      - name: firstname1
        type: string
actions:
  - name: getSinglePost
    url: https://jsonplaceholder.typicode.com/posts/1
    cliName: get-single-post
    method: post
    description: Get's an specific post from the endpoint
    in:
      headers:
        - name: accept-language
          type: string
    out:
      headers:
        - name: content-type
          type: string
      fields:
        - name: userId
          type: int64?
        - name: id
          type: int64
        - name: title
          type: string
        - name: body
          type: string
        - name: user
          type: object
          fields:
          - name: firstName
            type: string?
          - name: age
            type: int64
        - name: histories
          type: array
          fields:
          - name: firstName
            type: string?
          - name: age
            type: int64
          - name: info
            type: object
            fields:
            - name: memorySize
              type: int64
        `,
  },
  {
    id: "minimal",
    label: "minimal.yml",
    kind: "yaml",
    target: "goGen",
    tags: [],
    content: `name: minimalModule
actions:
  - name: getGreeting
    url: https://example.com/greeting
    method: get
    out:
      fields:
        - name: message
          type: string
`,
  },
  {
    id: "interfaces",
    label: "interfaces.yml",
    kind: "yaml",
    target: "jsGenModule",
    tags: ["typescript", "no-sdk", "no-package"],
    content: `name: interfacesModule

# An interface is a named set of fields. Every dto that lists it under
# \`implements\` gets those fields, and the compiler also generates an interface
# type (Go: an interface of Get<Field>() methods, TypeScript: an \`interface\`)
# that all of those dtos satisfy - so one function can accept any of them.

# A complex is a type you provide yourself (here TString: one text per language)
# and tell each compiler where to import it from.
complexes:
  - compiler: go
    name: TString
    namespace: complexes
    location: github.com/torabian/fireback/modules/fireback/complexes
  - compiler: ts
    name: TString
    location: "@fireback/complexes"

interfaces:
  - name: titlable
    description: Anything that has a title and some content.
    fields:
      - name: title
        type: string
        description: Short title shown in lists.
      - name: displayName
        type: complex
        complex: TString
        description: Localized name - one text per language, not a single string.
      - name: content
        type: object
        description: The body. Its type belongs to the interface, so it is shared.
        fields:
          - name: text
            type: string
          - name: language
            type: string

dtos:
  # Both dtos get title + content from the interface, plus their own fields.
  - name: cloth
    implements:
      - titlable
    fields:
      - name: size
        type: int64

  - name: shoe
    implements:
      - titlable
    fields:
      - name: color
        type: string
`,
  },
  {
    id: "interface-entity",
    label: "interface-entity.yml",
    kind: "yaml",
    target: "goGen",
    tags: [],
    content: `name: interfaceEntityModule

# Entities can implement interfaces too. The interface's fields become columns of
# the entity (ahead of its own). In Go the entity also gets the Get<Field>() methods,
# so InvoiceEntity satisfies Auditable - as does the plain InvoiceDto derived from it
# and the hand-written dto below. One function taking Auditable accepts all of them,
# and the audit fields are written once instead of on every record type.
interfaces:
  - name: auditable
    description: Who created a record, and an optional note about it.
    fields:
      - name: createdBy
        type: string
        description: Id of the user who created the record.
      - name: note
        type: string?
        description: Optional free text about the record.

entities:
  - name: invoice
    description: A billed invoice. createdBy and note come from auditable.
    implements:
      - auditable
    fields:
      - name: number
        type: string
      - name: total
        type: int64

dtos:
  # A dto implementing the same interface, next to the entity.
  - name: invoiceSummary
    implements:
      - auditable
    fields:
      - name: number
        type: string
`,
  },
  {
    id: "purchasable",
    label: "purchasable.yml",
    kind: "yaml",
    target: "goGen",
    tags: [],
    content: `name: shopModule

# Three different products, each with fields of its own, that all implement the
# \`purchasable\` interface. The interface's fields (sku, title, priceCents,
# currency) become columns of every product, and in Go each entity gets the
# Get<Field>() methods - so one function taking Purchasable can put a book, a
# laptop or a subscription in a cart, without knowing which one it is.
interfaces:
  - name: purchasable
    description: Anything a customer can put in a cart and pay for.
    fields:
      - name: sku
        type: string
        description: Stock keeping unit, unique per product.
      - name: title
        type: string
        description: Name shown in the shop.
      - name: priceCents
        type: int64
        description: Price in the smallest currency unit.
      - name: currency
        type: string
        description: ISO 4217 currency code, e.g. EUR.

entities:
  - name: book
    description: A physical or digital book.
    implements:
      - purchasable
    fields:
      - name: author
        type: string
      - name: isbn
        type: string
      - name: pageCount
        type: int64

  - name: laptop
    description: A laptop with hardware specs and a warranty.
    implements:
      - purchasable
    fields:
      - name: cpu
        type: string
      - name: ramGb
        type: int64
      - name: screenInches
        type: float64
      - name: warrantyMonths
        type: int64

  - name: subscription
    description: A recurring plan, billed every period.
    implements:
      - purchasable
    fields:
      - name: billingPeriod
        type: string
        description: For example monthly or yearly.
      - name: trialDays
        type: int64
      - name: autoRenew
        type: bool
`,
  },
  {
    id: "mcp-server-tool",
    label: "mcp-server-tool.yml",
    kind: "yaml",
    target: "goGen",
    tags: [],
    tagsByTarget: { jsGenModule: ["typescript", "no-sdk", "no-package"] },
    content: `name: billingModule
namespace: billing

# An invoice entity plus one extra action, exposed to AI agents as MCP tools.
#
# \`intent: true\` on the entity turns every action generated for it into an intent
# (an MCP tool): createInvoice, updateInvoice, getInvoice, listInvoices,
# previewDeleteInvoices and deleteInvoices. The compiler generates the typed tool
# signatures (name, description, input/output schema, behavior hints) to wire into
# an MCP server.
entities:
  - name: invoice
    intent: true
    description: A billed invoice.
    fields:
      - name: number
        type: string
      - name: customer
        type: string
      - name: total
        type: int64

actions:
  # A hand-written action next to the entity's own ones: renders an invoice.
  - name: printInvoice
    method: post
    url: /invoices/:uniqueId/print
    description: Renders an invoice as a printable document.
    in:
      fields:
        - name: format
          type: string?
          description: Output format, pdf (default) or html.
    out:
      fields:
        - name: downloadUrl
          type: string
        - name: pageCount
          type: int64

# Actions are not tools by themselves: an intent derives its input and output
# schema from the action it names in \`from\`, and carries the text the model
# reads to decide when to call it.
intents:
  - name: printInvoice
    from: printInvoice
    title: Print invoice
    description: Renders one invoice, by uniqueId, as a printable document and returns where to download it. Find the uniqueId with listInvoices first.
    annotations:
      readOnlyHint: true
      idempotentHint: true
`,
  },
  {
    id: "events-permissions",
    label: "events-permissions.yml",
    kind: "yaml",
    target: "jsGenModule",
    tags: ["typescript", "no-sdk", "no-package"],
    content: `name: blogModule

# The minimal version: one permission and one event, each with typed \`params\`
# that scope it. Params compile like a dto (TypeScript class, Go struct, ...).
# See events-permissions-full.yml for complexes, interfaces and JSON Schema.

permissions:
  - name: post
    key: post
    children:
      # Generates PostPermissionsPublishParams
      - name: publish
        key: publish
        params:
          fields:
            - name: workspaceId
              type: string

events:
  # Generates PostPublishedEventParams
  - key: postPublished
    permissions:
      - with: ["post.publish"]
    params:
      fields:
        - name: workspaceId
          type: string
`,
  },
  {
    id: "events-permissions-full",
    label: "events-permissions-full.yml",
    kind: "yaml",
    target: "jsGenModule",
    tags: ["typescript", "json-schema", "no-sdk", "no-package"],
    content: `name: blogModule

# Full version of events-permissions.yml: complexes and interfaces inside params,
# nested objects, a payload, and (with the json-schema tag) JSON Schema.
#
# Permissions and events can both declare \`params\`: a typed shape that scopes
# them (e.g. the workspace a grant applies within, or an event fired in). Each is
# compiled like a dto - Go struct, TypeScript class, Kotlin/Swift class - and with
# the json-schema tag the TypeScript class also carries its JSON Schema.
# Params can use any field type, including a complex you provide yourself.

complexes:
  - compiler: go
    name: TString
    namespace: complexes
    location: github.com/torabian/fireback/modules/fireback/complexes
  - compiler: ts
    name: TString
    location: "@fireback/complexes"

# Params can implement interfaces just like dtos: the interface's fields are
# included, and the generated type satisfies the interface (Go: Get<Field>()
# methods, TypeScript: an \`implements\` clause).
interfaces:
  - name: scoped
    description: Anything that applies within one workspace.
    fields:
      - name: workspaceId
        type: string
        description: The workspace this applies within.

permissions:
  - name: post
    key: post
    title:
      en: Posts
    children:
      # Generates PostPermissionsPublishParams
      - name: publish
        key: publish
        title:
          en: Publish posts
        params:
          implements:
            - scoped
          fields:
            - name: label
              type: complex
              complex: TString
              description: Localized label - one text per language.

events:
  # Generates PostPublishedEventParams (plus the payload type)
  - key: postPublished
    name:
      en: Post published
    permissions:
      - with: ["post.publish"]
    params:
      implements:
        - scoped
      fields:
        - name: channel
          type: object
          fields:
            - name: name
              type: string
            - name: title
              type: complex
              complex: TString
              description: Localized channel title.
    payload:
      fields:
        - name: postId
          type: string
`,
  },
  {
    id: "allegro-offer-management",
    label: "allegro-offer-management.yml",
    kind: "yaml",
    target: "jsGenModule",
    tags: ["typescript", "no-sdk"],
    content: allegroOfferManagement,
  },
  {
    id: "allegro-offer-translations",
    label: "allegro-offer-translations.yml",
    kind: "yaml",
    target: "jsGenModule",
    tags: ["typescript", "no-sdk"],
    content: allegroOfferTranslations,
  },
  {
    id: "allegro-user-offer-information",
    label: "allegro-user-offer-information.yml",
    kind: "yaml",
    target: "jsGenModule",
    tags: ["typescript", "no-sdk"],
    content: allegroUserOfferInformation,
  },
  {
    id: "query",
    label: "query.sql",
    kind: "sql",
    target: "sqlQueryPredict",
    tags: [],
    content: `SELECT 
        u.user_id as user_id,
        field(u.user_name, 'string') as UserName,
        u.user_email,
        field(COUNT(o.order_id), 'int64') AS total_orders,
        COALESCE(SUM(o.total), 0) AS total_spent,
        MAX(o.total) AS max_order,
        (
            SELECT COUNT(*) 
            FROM (
                SELECT 101 AS order_id, 1 AS user_id, 120.5 AS total
                UNION ALL
                SELECT 102, 1, 50.0
                UNION ALL
                SELECT 103, 2, 75.0
                UNION ALL
                SELECT 104, 3, 200.0
                UNION ALL
                SELECT 105, 3, 25.0
            ) o2
            WHERE o2.user_id = u.user_id AND o2.total > 50
        ) AS big_orders_count
    FROM (
        SELECT 1 AS user_id, 'Alice' AS user_name, 'alice@example.com' AS user_email
        UNION ALL
        SELECT 2, 'Bob', 'bob@example.com'
        UNION ALL
        SELECT 3, 'Carol', 'carol@example.com'
    ) u
    LEFT JOIN (
        SELECT 101 AS order_id, 1 AS user_id, 120.5 AS total
        UNION ALL
        SELECT 102, 1, 50.0
        UNION ALL
        SELECT 103, 2, 75.0
        UNION ALL
        SELECT 104, 3, 200.0
        UNION ALL
        SELECT 105, 3, 25.0
    ) o
    ON u.user_id = o.user_id
    -- WHERE  u.user_name != 'Alice'        -- filter rows before aggregation
    WHERE filter()
    GROUP BY u.user_id, u.user_name, u.user_email
    HAVING MAX(o.total) > 100        -- filter groups after aggregation
    ORDER BY total_spent DESC
    limit useval('limit')
`,
  },
  {
    id: "entity-postgres",
    label: "entity-postgres.yml",
    kind: "yaml",
    target: "entitySqlGen",
    tags: [],
    content: `name: shopModule

# Run the compiler as often as you like: the generated sql only creates what is
# missing (create table if not exists, add column if not exists, constraints added
# only when absent). Postgres is the default - pick the "sqlite" tag in the compiler
# options for the sqlite version of the same tables. Check the output for:
#   object      -> flattened columns        (address_street, address_geo_lat)
#   array       -> child table + linker_id  (order_entity_items, nested ..._items_discounts)
#   collection  -> many2many join table     (product_tags)
#   one         -> foreign key constraint,  one? -> column + index only
#   enum        -> CHECK constraint
entities:
  - name: category
    description: Categories form a tree, a category may have a parent.
    fields:
      - name: title
        type: string
      - name: parent
        type: one?
        target: CategoryEntity

  - name: tag
    fields:
      - name: label
        type: string

  - name: customer
    fields:
      - name: email
        type: string
      - name: tier
        type: enum
        of:
          - k: standard
          - k: silver
          - k: gold
      - name: birthYear
        type: int?
      # object nested in an object: address_street, address_city, address_geo_lat, ...
      - name: address
        type: object
        fields:
          - name: street
            type: string
          - name: city
            type: string
          - name: geo
            type: object?
            fields:
              - name: lat
                type: float64
              - name: lng
                type: float64
      # array: every customer owns any number of phones
      - name: phones
        type: array
        fields:
          - name: number
            type: string
          - name: primary
            type: bool
            default: false

  - name: product
    fields:
      - name: title
        type: string
      - name: description
        type: string?
      - name: status
        type: enum
        of:
          - k: draft
          - k: published
          - k: archived
      - name: category
        type: one
        target: CategoryEntity
      - name: dimensions
        type: object?
        fields:
          - name: width
            type: float64
          - name: height
            type: float64
      - name: attributes
        type: map
      # array inside an array: product -> variants -> prices
      - name: variants
        type: array
        fields:
          - name: sku
            type: string
          - name: stock
            type: int
            default: 0
          - name: prices
            type: array
            fields:
              - name: currency
                type: string
              - name: amountCents
                type: int64
      - name: tags
        type: collection
        target: TagEntity
      # a collection pointing to the entity itself
      - name: relatedProducts
        type: collection?
        target: ProductEntity

  - name: order
    fields:
      - name: number
        type: string
      - name: status
        type: enum
        of:
          - k: pending
          - k: paid
          - k: shipped
          - k: cancelled
      - name: customer
        type: one
        target: CustomerEntity
      # one?: nullable reference, indexed but without a constraint
      - name: referrer
        type: one?
        target: CustomerEntity
      - name: shipping
        type: object
        fields:
          - name: method
            type: string
          - name: costCents
            type: int64
      - name: items
        type: array
        fields:
          - name: quantity
            type: int
          - name: product
            type: one
            target: ProductEntity
          - name: discounts
            type: array?
            fields:
              - name: code
                type: string
              - name: percent
                type: float32
`,
  },
];

export const kindForTarget = (target: string): ExampleKind =>
  target === "sqlQueryPredict" ? "sql" : "yaml";

export const examplesForTarget = (target: string) =>
  examples.filter((e) => e.kind === kindForTarget(target));

export const getExample = (id: string) =>
  examples.find((e) => e.id === id) ?? examples[0];
