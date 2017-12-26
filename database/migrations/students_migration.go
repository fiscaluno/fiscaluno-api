package migrations

type Student struct {
    Id      uint
    Name    string  `gorm:"size:100"`
    Email   string  `gorm:"not null;unique"`
    Age     int32   `gorm:type:tinyint;unsigned`
}

func init() {
    Create(&Student{})
}
