package migrations

type Student struct {
    Name    string
    Email   string
    Age     int32
}

func init() {
    Create(&Student{})
}
