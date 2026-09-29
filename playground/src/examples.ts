// Predefined definitions shown in the playground's file explorer.
//
// To add an example, append an entry to `examples` below - nothing else is
// needed. `kind: "yaml"` examples are listed for every compiler target except
// QueryPredict(SQL); `kind: "sql"` examples only for that target. The first
// example of each kind is the default.
export type ExampleKind = "yaml" | "sql";

export interface Example {
  /** Unique, stable id (also used as the editor's model path). */
  id: string;
  /** File name shown in the explorer. */
  label: string;
  kind: ExampleKind;
  content: string;
}

export const examples: Example[] = [
  {
    id: "definition",
    label: "definition.yml",
    kind: "yaml",
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
    id: "query",
    label: "query.sql",
    kind: "sql",
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
];

export const kindForTarget = (target: string): ExampleKind =>
  target === "sqlQueryPredict" ? "sql" : "yaml";

export const examplesForTarget = (target: string) =>
  examples.filter((e) => e.kind === kindForTarget(target));

export const getExample = (id: string) =>
  examples.find((e) => e.id === id) ?? examples[0];
