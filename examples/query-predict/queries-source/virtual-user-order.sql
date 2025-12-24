SELECT 
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
