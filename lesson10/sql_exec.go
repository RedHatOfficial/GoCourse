result, err := db.Exec(
	"INSERT INTO users (name, age) VALUES ($1, $2)",
	"gopher",
	27,
)
