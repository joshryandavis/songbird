package starling

type DirectDebitMandate struct {
	UID            string      `json:"uid"`
	Reference      string      `json:"reference"`
	Status         string      `json:"status"`
	Source         string      `json:"source"`
	Created        DateTime    `json:"created"`
	Cancelled      DateTime    `json:"cancelled"`
	NextDate       string      `json:"nextDate"`
	LastDate       string      `json:"lastDate"`
	OriginatorName string      `json:"originatorName"`
	OriginatorUID  string      `json:"originatorUid"`
	MerchantUID    string      `json:"merchantUid"`
	LastPayment    LastPayment `json:"lastPayment"`
	AccountUID     string      `json:"accountUid"`
	CategoryUID    string      `json:"categoryUid"`
}
