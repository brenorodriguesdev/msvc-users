package contracts

type HasherCompare interface {
	Compare(password, hash string) (bool, error)
}
