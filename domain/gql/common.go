package gql

type SortOrder struct {
	Sort  *string `json:"sort"`
	Order *Order  `json:"order"`
}

type Order string

const (
	OrderAsc  Order = "ASC"
	OrderDesc Order = "DESC"
)
