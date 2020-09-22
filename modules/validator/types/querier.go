package types

// query endpoints supported by the validator Querier
const (
	QueryValidators = "validators"
	QueryValidator  = "validator"
	QueryParameters = "parameters"
)

type QueryValidatorsParams struct {
	Page   int
	Limit  int
	Jailed string
}

func NewQueryValidatorsParams(page, limit int, jailed string) QueryValidatorsParams {
	return QueryValidatorsParams{page, limit, jailed}
}

type QueryValidatorParams struct {
	ID string
}

func NewQueryValidatorParams(id string) QueryValidatorParams {
	return QueryValidatorParams{
		ID: id,
	}
}
