package structs

type Framework struct { name string }
type Language struct { name string }
type DB struct { name string }

type IT struct {
	L *Language
	F *Framework
	D *DB
}
