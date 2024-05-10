package data

type URLStorage interface {
	New() URLStorage
	Link() LinkQ
}
