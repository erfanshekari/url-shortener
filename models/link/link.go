package link

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/erfanshekari/url-shortener/db"
)

type Link struct {
	Target string
	Slug   string
}

func (l *Link) Save() error {
	db := db.GetInstance()
	err := db.Set([]byte(l.Target), []byte(l.Slug))
	if err != nil {
		return err
	}
	return db.Set([]byte(l.Slug), []byte(l.Target))
}

func (l *Link) ToJson() map[string]interface{} {
	asJson := make(map[string]interface{})
	asJson["target"] = l.Target
	asJson["slug"] = l.Slug
	return asJson
}

func FindByTarget(target string) (*Link, error) {
	db := db.GetInstance()

	val, err := db.Get([]byte(target))

	if err == badger.ErrKeyNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if val != nil {
		value := string(*val)
		link := Link{
			Target: target,
			Slug:   value,
		}
		return &link, nil
	}

	return nil, nil
}

func FindBySlug(slug string) (*Link, error) {
	db := db.GetInstance()

	val, err := db.Get([]byte(slug))

	if err == badger.ErrKeyNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if val != nil {
		value := string(*val)
		link := Link{
			Target: value,
			Slug:   slug,
		}
		return &link, nil
	}

	return nil, nil
}
