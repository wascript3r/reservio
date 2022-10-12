package pgsql

import (
	"errors"
	"fmt"
	"strings"
)

type QueryType int8

const (
	SelectQuery QueryType = iota + 1
	UpdateQuery
)

type queryDescriptor struct {
	// Initial query word to be placed before the subqueries (WHERE, SET, etc.)
	initialWord string

	// Replace this word in the base query with the combined subqueries (<where>, <set>, etc.)
	replaceWord string

	// Subqueries are joined with this separator (", ", " ", etc.)
	joinSeparator string
}

var (
	ErrDescriptorNotFound = errors.New("descriptor not found")

	descriptors = map[QueryType]*queryDescriptor{
		SelectQuery: {
			initialWord:   "WHERE",
			replaceWord:   "<where>",
			joinSeparator: " ",
		},
		UpdateQuery: {
			initialWord:   "SET",
			replaceWord:   "<set>",
			joinSeparator: ", ",
		},
	}
)

type QueryBuilder struct {
	descriptor   *queryDescriptor
	baseQuery    string
	startParamID int
	subQueries   []string
	params       []any
}

func NewQueryBuilder(t QueryType, baseQuery string, startParamID int) *QueryBuilder {
	d, ok := descriptors[t]
	if !ok {
		panic(ErrDescriptorNotFound)
	}

	return &QueryBuilder{
		descriptor:   d,
		baseQuery:    baseQuery,
		startParamID: startParamID,
	}
}

func (q *QueryBuilder) Add(subQuery string, param any) {
	nextID := fmt.Sprintf("$%d", q.startParamID+len(q.params))
	subQuery = strings.Replace(subQuery, "?", nextID, 1)

	q.subQueries = append(q.subQueries, subQuery)
	if param != nil {
		q.params = append(q.params, param)
	}
}

func (q *QueryBuilder) GetQuery() string {
	subStr := strings.Join(q.subQueries, q.descriptor.joinSeparator)
	combined := fmt.Sprintf("%s %s", q.descriptor.initialWord, subStr)

	return strings.Replace(q.baseQuery, q.descriptor.replaceWord, combined, 1)
}

func (q *QueryBuilder) GetParams(initialParams ...any) []any {
	return append(initialParams, q.params...)
}
