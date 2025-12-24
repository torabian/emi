update users set 
    name = valueof(?, 'string', 'Nameee', true),
    lastnamevalue = ?,
    last_order_total = valueof(
        (
        SELECT total
        FROM orders
        WHERE orders.user_id = users.id
        ORDER BY id DESC
        LIMIT 1
        ),
        'string',
        'LastOrderTotal',
        true
    )
