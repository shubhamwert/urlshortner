package model

type UrlModel struct {
	Id          string `json:"id"`
	OriginalUrl string `json:"url" bson:"OriginalUrl"`
	EncodedUrl  string `json:"encodedUrl" bson:"EncodedUrl"`
}

func (u *UrlModel) GetUrlModel() (UrlModel, error) {
	return *u, nil
}

func CreateUrlModel(id string, url string, encoding string) (*UrlModel, error) {
	u := &UrlModel{
		OriginalUrl: url,
		EncodedUrl:  encoding,
	}

	return u, nil
}
