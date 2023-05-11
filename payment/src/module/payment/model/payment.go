package model

type Payment struct {
	Name string `json:"name" bson:"name" filterable:"true" sortable:"true"`
	Bs   Bs     `json:"bs" bson:"bs" filterable:"true" `
}

type Bs struct {
	Name string `json:"name" bson:"name" filterable:"true"`
	Age  string `json:"age" bson:"age" sortable:"true"`
}

//&filter.name=$not:$in:"2","5","7"&filter.giang=123
