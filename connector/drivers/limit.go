package drivers

import "github.com/lab210-dev/dbkit/specs"

type limit struct {
	limit  int
	offset int
}

func (l *limit) Offset() int {
	return l.offset
}

func (l *limit) Limit() int {
	return l.limit
}

func (l *limit) SetOffset(index int) specs.DriverLimit {
	l.offset = index
	return l
}

func (l *limit) SetLimit(index int) specs.DriverLimit {
	l.limit = index
	return l
}

func NewLimit() specs.DriverLimit {
	return new(limit)
}