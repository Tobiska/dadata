package suggest

import "context"

type BankByIDParams struct {
	Query string  `json:"query"`
	KPP   *string `json:"kpp,omitempty"`
}

func NewBankByIDParams(query string) *BankByIDParams {
	return &BankByIDParams{
		Query: query,
	}
}

func (p *BankByIDParams) SetKpp(kpp string) *BankByIDParams {
	p.KPP = &kpp
	return p
}

func (p *BankByIDParams) SetQuery(query string) *BankByIDParams {
	p.Query = query
	return p
}

// Bank find banks by ID
func (a *Api) BankByID(ctx context.Context, params *BankByIDParams) (ret []*BankSuggestion, err error) {
	var result = &BankResponse{}

	err = a.Client.Post(ctx, "findById/bank", params, result)
	if err != nil {
		return nil, err
	}
	return result.Suggestions, nil
}

// Bank try to return suggest banks by params
func (a *Api) Bank(ctx context.Context, params *RequestParams) (ret []*BankSuggestion, err error) {
	var result = &BankResponse{}
	err = a.Client.Post(ctx, "suggest/bank", params, result)
	if err != nil {
		return
	}
	ret = result.Suggestions
	return
}
