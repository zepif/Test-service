package data

type LinkQ interface {
	Insert(FullURL, ShortURL string) error
	Get(ShortURL string) (string, error)
}

type Link struct {
	FullURL  string `db:"FullURL" structs:"FullURL"`
	ShortURL string `db:"ShortURL" structs:"ShortURL"`
}
