package market

import "github.com/JamesMike-a11y/go-okx/rest/api"

func NewGetInstruments(param *GetInstrumentsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/public/instruments",
		Method: api.MethodGet,
		Param:  param,
	}, &GetBooksResponse{}
}

type GetInstrumentsParam struct {
	InstId     string `url:"instId"`
	InstType   string `url:"instType"`
	InstFamily string `url:"instFamily"`
}

type GetInstrumentsResponse struct {
	api.Response
	Data []Instrument `json:"data"`
}

type Instrument struct {
	Alias             string   `json:"alias"`             // 交易对别名
	AuctionEndTime    string   `json:"auctionEndTime"`    // 竞价结束时间
	BaseCcy           string   `json:"baseCcy"`           // 基础货币，如BTC
	Category          string   `json:"category"`          // 产品类别
	CtMult            string   `json:"ctMult"`            // 合约乘数 (适用于衍生品)
	CtType            string   `json:"ctType"`            // 合约类型 (适用于衍生品)
	CtVal             string   `json:"ctVal"`             // 合约面值 (适用于衍生品)
	CtValCcy          string   `json:"ctValCcy"`          // 合约面值计价货币 (适用于衍生品)
	ContTdSwTime      string   `json:"contTdSwTime"`      // 连续交易切换时间 (毫秒时间戳)
	ExpTime           string   `json:"expTime"`           // 到期日 (适用于衍生品)
	FutureSettlement  bool     `json:"futureSettlement"`  // 是否为期货结算
	InstFamily        string   `json:"instFamily"`        // 产品族
	InstId            string   `json:"instId"`            // 交易对ID，如BTC-USDT
	InstType          string   `json:"instType"`          // 产品类型 (如SPOT, FUTURES, SWAP)
	Lever             string   `json:"lever"`             // 最大杠杆倍数
	ListTime          string   `json:"listTime"`          // 上线时间 (毫秒时间戳)
	LotSz             string   `json:"lotSz"`             // 数量精度 (最小数量单位)
	MaxIcebergSz      string   `json:"maxIcebergSz"`      // 最大冰山委托数量
	MaxLmtAmt         string   `json:"maxLmtAmt"`         // 限价单最大委托金额
	MaxLmtSz          string   `json:"maxLmtSz"`          // 限价单最大委托数量
	MaxMktAmt         string   `json:"maxMktAmt"`         // 市价单最大委托金额
	MaxMktSz          string   `json:"maxMktSz"`          // 市价单最大委托数量
	MaxStopSz         string   `json:"maxStopSz"`         // 止损单最大委托数量
	MaxTriggerSz      string   `json:"maxTriggerSz"`      // 触发单最大委托数量
	MaxTwapSz         string   `json:"maxTwapSz"`         // TWAP订单最大委托数量
	MinSz             string   `json:"minSz"`             // 最小交易数量
	OptType           string   `json:"optType"`           // 期权类型 (适用于期权)
	OpenType          string   `json:"openType"`          // 开盘类型
	PreMktSwTime      string   `json:"preMktSwTime"`      // 盘前切换时间
	QuoteCcy          string   `json:"quoteCcy"`          // 计价货币，如USDT
	TradeQuoteCcyList []string `json:"tradeQuoteCcyList"` // 交易计价货币列表
	SettleCcy         string   `json:"settleCcy"`         // 结算货币 (适用于衍生品)
	State             string   `json:"state"`             // 状态 (live: 交易中)
	RuleType          string   `json:"ruleType"`          // 规则类型
	Stk               string   `json:"stk"`               // 行权价 (适用于期权)
	TickSz            string   `json:"tickSz"`            // 价格精度 (最小价格单位)
	Uly               string   `json:"uly"`               //  underlying (适用于衍生品)
	InstIdCode        int64    `json:"instIdCode"`        // 产品ID数字代码
}
