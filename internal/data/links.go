package data

type LinkQ interface  {
    Get() (*Link, error)
	Insert(value Link) (*Link, error)
}

type LinkQ struct {
    Id        int64  `db:"id" structs:"-"`
	full_url  string `db:"full_url" structs:"full_url"`
    short_url string `db:"short_url" structs: "short_url"`
}
