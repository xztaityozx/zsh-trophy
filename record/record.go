package record

type Record struct {
	Name   string
	Args   map[string]string
	Status map[int]bool
}
