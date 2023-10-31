var age int64
result, err := db.QueryRow("SELECT age FROM users WHERE name = $1", name)
result.Scan(&age)
