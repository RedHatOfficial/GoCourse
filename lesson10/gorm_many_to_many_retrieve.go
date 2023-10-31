// Retrieve user list with eager loading languages
func GetAllUsers(db *gorm.DB) ([]User, error) {
  var users []User
  err := db.Model(&User{}).Preload("Languages").Find(&users).Error
  return users, err
}

// Retrieve language list with eager loading users
func GetAllLanguages(db *gorm.DB) ([]Language, error) {
  var languages []Language
  err := db.Model(&Language{}).Preload("Users").Find(&languages).Error
  return languages, err
}
