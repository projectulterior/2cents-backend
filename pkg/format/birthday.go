package format

type Birthday struct {
	Day   int `bson:"day"`
	Month int `bson:"month"`
	Year  int `bson:"year"`
}
