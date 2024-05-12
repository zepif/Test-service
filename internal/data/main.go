package data

type MasterQ interface {
	New() MasterQ
	Link() LinkQ
}
