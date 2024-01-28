package model

import "context"

type UrlStorageDB interface {
	Get(ctx context.Context, encodedUrl string, Owner string) (UrlModel, error)
	Set(ctx context.Context, url UrlModel) error
	Delete(encodedUrl string) error
	ListAllKeys()
}
