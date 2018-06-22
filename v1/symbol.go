package bitfinex

type SymbolService struct {
	client *Client
}

// Get - return symbol list
func (s *SymbolService) Get() ([]string, error) {
	req, err := s.client.newRequest("GET", "symbols", nil)

	if err != nil {
		return nil, err
	}

	var v []string
	_, err = s.client.do(req, &v)

	return v, err
}

type SymbolDetail struct {
	Pair string
	PricePrecision int 		`json:"price_precision"`
	InitialMargin string 	`json:"initial_margin"`
	MinumumMargin string 	`json:"minimum_margin"`
	MaximumOrderSize string `json:"maximum_order_size"`
	MinimumOrderSize string `json:"minimum_order_size"`
}

func (s *SymbolService) Details() ([]SymbolDetail, error) {
	req, err := s.client.newRequest("GET", "symbols_details", nil)

	if err != nil {
		return nil, err
	}

	var v []SymbolDetail
	_, err = s.client.do(req, &v)

	return v, err
}
