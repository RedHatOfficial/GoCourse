user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
result := db.Create(&user) // pass pointer of data to Create

user.ID             // returns inserted data's primary key
result.Error        // returns error
result.RowsAffected // returns inserted records count
