package domain

type Thing struct {
	Field  string `json:"field" bson:"field" msgpack:"field" xml:"field"`
	URL    string `json:"url" bson:"url" msgpack:"url" xml:"url" validate:"empty=false & format=url"`
	Number int64  `json:"number" bson:"number" msgpack:"number" xml:"number"`
}
