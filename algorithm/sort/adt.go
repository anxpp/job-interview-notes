package sort

type Sortable interface {
	Less(b Sortable) bool
}
