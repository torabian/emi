package xxx

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
    "regexp"
)

const VirtualUserOrderSQL = `SELECT 
        u.user_id as user_id,
        u.user_name as UserName,
        u.user_email,
        COUNT(o.order_id) AS total_orders,
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
`

type VirtualUserOrderRow struct {
	UserId string
	UserName string
	UserEmail string
	TotalOrders int64
	TotalSpent string
	MaxOrder string
	BigOrdersCount string
}

type VirtualUserOrderContext struct {
    Filter string
    Having string
    Restriction string
    Params      map[string]interface{}
    Placeholders      []any
}


func VirtualUserOrderPrepreSql(ctx VirtualUserOrderContext) (string, error) {
    replaceUseVal := func(sql string, values map[string]interface{}) string {
        re := regexp.MustCompile(`useval\(\s*['"]([^'"]+)['"]\s*\)`)
        return re.ReplaceAllStringFunc(sql, func(match string) string {
            if m := re.FindStringSubmatch(match); len(m) > 1 {
                if val, ok := values[m[1]]; ok {
                    // escape the value safely
                    switch v := val.(type) {
                    case string:
                        safe := strings.ReplaceAll(v, `'`, `''`)
                        return "'" + safe + "'"
                    default:
                        return fmt.Sprintf("%v", v)
                    }
                }
            }
            return match
        })
    }

    script := replaceUseVal(VirtualUserOrderSQL, ctx.Params)
    filter := "1"
	if ctx.Filter != "" {
		filter = ctx.Filter
	}
	script = strings.ReplaceAll(script, "filter()", "(" +filter+ ")")

    restriction := "1"
    if ctx.Restriction != "" {
        restriction = ctx.Restriction
    }
    script = strings.ReplaceAll(script, "restriction()", "(" +restriction + ")")
    
    having := ""
    if ctx.Having != "" {
        having = ctx.Having
    }
    script = strings.ReplaceAll(script, "having()", having)

    return script, nil
}


func VirtualUserOrder(db *sql.DB, ctx VirtualUserOrderContext,) ([]VirtualUserOrderRow, error) {
    script, err := VirtualUserOrderPrepreSql(ctx)
    if err != nil {
		return nil, err
	}

    log.Default().Println(script)

	rows, err := db.Query(script, ctx.Placeholders...)
	if err != nil {
		return nil, fmt.Errorf("query VirtualUserOrder failed: %w", err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []VirtualUserOrderRow
	for rows.Next() {
		var r VirtualUserOrderRow

		scanArgs := make([]interface{}, len(cols))
		for i, col := range cols {
			switch col {
			case "user_id":
				scanArgs[i] = &r.UserId
			case "UserName":
				scanArgs[i] = &r.UserName
			case "user_email":
				scanArgs[i] = &r.UserEmail
			case "total_orders":
				scanArgs[i] = &r.TotalOrders
			case "total_spent":
				scanArgs[i] = &r.TotalSpent
			case "max_order":
				scanArgs[i] = &r.MaxOrder
			case "big_orders_count":
				scanArgs[i] = &r.BigOrdersCount
			default:
				var discard interface{}
				scanArgs[i] = &discard
			}
		}

		if err := rows.Scan(scanArgs...); err != nil {
			return nil, err
		}
		results = append(results, r)
	}

	return results, nil
}
