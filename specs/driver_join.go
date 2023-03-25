package specs

type JoinMethod int

type DriverJoin interface {
	Validate() error

	FromTableIndex() int
	FromKey() string

	ToTable() string
	ToTableIndex() int
	ToKey() string
	ToDatabase() string

	Method() string

	SetMethod(method JoinMethod) DriverJoin
	SetFromTableIndex(fromTableIndex int) DriverJoin
	SetFromKey(fromKey string) DriverJoin

	SetToTable(toTable string) DriverJoin
	SetToTableIndex(toTableIndex int) DriverJoin
	SetToKey(toKey string) DriverJoin
	SetToDatabase(toModel string) DriverJoin
}
