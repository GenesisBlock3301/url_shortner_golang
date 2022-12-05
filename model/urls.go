package model

type URL struct {
	Key       string `bson:"key omitempty"`
	SecretKey string `bson:"secretKey omitempty"`
	TargetUrl string `bson:"targetUrl omitempty"`
	IsActive  string `bson:"isActive omitempty"`
	Clicks    string `bson:"clicks omitempty"`
}
