package data

type LinkQ interface {
	Insert(fullURL, shortURL string) error
	Get(id int64) (string, string, error)
}

type Link struct {
    Id        int64  `db:"id" structs:"-"`
	full_url  string `db:"full_url" structs:"full_url"`
    short_url string `db:"short_url" structs: "short_url"`
}
