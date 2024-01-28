package model

type UrlModel struct {
	Id          string `json:"id"`
	OriginalUrl string `json:"url" bson:"OriginalUrl" dynamodbav:"originalUrl"`
	EncodedUrl  string `json:"encodedUrl" bson:"EncodedUrl" dynamodbav:"encodedURL"`
	Owner       string `json:"user" bson:"User" dynamodbav:"userId"`
}

func (u *UrlModel) GetUrlModel() (UrlModel, error) {
	return *u, nil
}

func CreateUrlModel(id string, url string, encoding string, owner string) (*UrlModel, error) {
	u := &UrlModel{
		OriginalUrl: url,
		EncodedUrl:  encoding,
		Owner:       owner,
	}

	return u, nil
}
