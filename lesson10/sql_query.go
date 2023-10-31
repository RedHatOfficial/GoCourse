rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
if err != nil {
	log.Fatal(err)
}

defer rows.Close()
for rows.Next() {
	var name string
	if err := rows.Scan(&name); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s is %d\n", name, age)
}

if err := rows.Err(); err != nil {
	log.Fatal(err)
}
