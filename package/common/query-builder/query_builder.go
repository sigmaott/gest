package query_builder

type IQueryBuilder interface {
	Parser(queryStr string, query any) error
}
