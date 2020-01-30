package engine

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Item     []interface{}
}

func NilParseResult([]byte) ParseResult {
	return ParseResult{}
}
