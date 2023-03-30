package spot

type NewOrderRequest struct {
	Symbol    string `json:"symbol"`
	Side      string `json:"side"`
	Type      string `json:"type"`
	Quantity  string `json:"quantity"`
	Timestamp string `json:"timestamp"`
}

type NewOrderResponse struct {
	Symbol  string `json:"symbol"`
	Side    string `json:"side"`
	Type    string `json:"type"`
	OrderId int64  `json:"orderId"`
	// unless OCO, value will be -1
	//OrderListId   int64  `json:"orderListId"`
	//ClientOrderId string `json:"clientOrderId"`
	//TransactTime  int64  `json:"transactTime"`
	//Price         string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
}

type CancelOrderRequest struct {
	Symbol    string `json:"symbol"`
	OrderId   int64  `json:"orderId"`
	Timestamp int64  `json:"timestamp"`
}

type CancelOrderResponse struct {
	Symbol            string `json:"symbol"`
	Type              string `json:"type"`
	Side              string `json:"side"`
	OrderId           int64  `json:"orderId"`
	OrigClientOrderId string `json:"origClientOrderId"`
	ClientOrderId     string `json:"clientOrderId"`
	Status            string `json:"status"`
}
