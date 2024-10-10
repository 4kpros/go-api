package types

type Permission struct {
	Table  string
	Create bool
	Read   bool
	Update bool
	Delete bool
}
