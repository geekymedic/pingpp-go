package pingpp

// Order相关的数据结构
type (
	// OrderCreateParams 创建Order请求数据类型
	OrderCreateParams struct {
		App               string                 `json:"app"`
		Uid               string                 `json:"uid,omitempty"`
		MerchantOrderNo   string                 `json:"merchant_order_no,omitempty"`
		Amount            int64                  `json:"amount"`
		BalanceSettlement BalanceSettlement      `json:"balance_settlement,omitempty"`
		Currency          string                 `json:"currency,omitempty"`
		ClientIp          string                 `json:"client_ip,omitempty"`
		Subject           string                 `json:"subject,omitempty"`
		Body              string                 `json:"body,omitempty"`
		Description       string                 `json:"description,omitempty"`
		Coupon            string                 `json:"coupon,omitempty"`
		ActualAmount      int64                  `json:"actual_amount,omitempty"`
		TimeExpire        int64                  `json:"time_expire,omitempty"`
		Metadata          map[string]interface{} `json:"metadata,omitempty"`
		ReceiptApp        string                 `json:"receipt_app,omitempty"`
		ServiceApp        string                 `json:"service_app,omitempty"`
		RoyaltyUsers      []RoyaltyUser          `json:"royalty_users,omitempty"`
		RoyaltyTemplate   string                 `json:"royalty_template,omitempty"`
	}
	// 余额结算信息
	BalanceSettlement struct {
		User    int `json:"user"`
		UserFee int `json:"user_fee,omitempty"`
	}
	// RoyaltyUser 分润用户
	RoyaltyUser struct {
		User   string `json:"user"`
		Amount int    `json:"amount"`
	}
	// OrderPayParams 订单支付请求参数
	OrderPayParams struct {
		ChargeAmount  *int64                 `json:"charge_amount,omitempty"`
		Channel       string                 `json:"channel,omitempty"`
		Extra         map[string]interface{} `json:"extra,omitempty"`
		TimeExpire    int64                  `json:"time_expire,omitempty"` // 时间戳
		ChargeOrderNo string                 `json:"charge_order_no,omitempty"`
	}

	// Order 数据类型
	Order struct {
		ID               string                 `json:"id"`
		Object           string                 `json:"object"`
		Created          int64                  `json:"created"`
		Livemode         bool                   `json:"livemode"`
		Refunded         bool                   `json:"refunded"`
		Status           string                 `json:"status"`
		Paid             bool                   `json:"paid"`
		App              string                 `json:"app"`
		Uid              string                 `json:"uid"`
		MerchantOrderNo  string                 `json:"merchant_order_no"`
		Amount           int64                  `json:"amount"`
		CouponAmount     int64                  `json:"coupon_amount"`
		ActualAmount     int64                  `json:"actual_amount"`
		AmountPaid       int64                  `json:"amount_paid"`
		AmountRefunded   int64                  `json:"amount_refunded"`
		Currency         string                 `json:"currency"`
		Subject          string                 `json:"subject"`
		Body             string                 `json:"body"`
		ClientIp         string                 `json:"client_ip"`
		TimePaid         int64                  `json:"time_paid"`
		TimeExpire       int64                  `json:"time_expire"`
		Coupon           string                 `json:"coupon"`
		Charge           string                 `json:"charge"`
		Charges          ChargeList             `json:"charges"`
		ChargeEssentials map[string]interface{} `json:"charge_essentials"`
		AvailableBalance int64                  `json:"available_balance"`
		ReceiptApp       string                 `json:"receipt_app"`
		ServiceApp       string                 `json:"service_app"`
		AvailableMethods []string               `json:"available_methods"`
		Description      string                 `json:"description"`
		Metadata         map[string]interface{} `json:"metadata"`
	}

	// OrderList Order 列表数据类型
	OrderList struct {
		ListMeta
		Values []*Order `json:"data"`
	}

	// Recharge 数据类型
	Recharge struct {
		ID                 string                 `json:"id"`
		Object             string                 `json:"object"`
		App                string                 `json:"app"`
		Created            int64                  `json:"created"`
		Livemode           bool                   `json:"livemode"`
		Amount             int                    `json:"amount"`
		Succeeded          bool                   `json:"succeeded"`
		TimeSucceed        uint64                 `json:"time_succeed"`
		Refunded           bool                   `json:"refunded"`
		User               string                 `json:"user"`
		From               string                 `json:"from"`
		UserFee            int                    `json:"user_fee"`
		Charge             Charge                 `json:"charge"`
		BalanceBonus       BalanceBonus           `json:"balance_bonus"`
		BalanceTransaction string                 `json:"balance_transaction"`
		Description        string                 `json:"description"`
		Metadata           map[string]interface{} `json:"metadata"`
	}

	// RechargeList Recharge 列表数据类型
	RechargeList struct {
		ListMeta
		Values []*Recharge `json:"data"`
	}

	// RechargeRefundParams 充值退款请求参数
	RechargeRefundParams struct {
		Description   string                 `json:"description,omitempty"`
		Metadata      map[string]interface{} `json:"metadata,omitempty"`
		FundingSource string                 `json:"funding_source,omitempty"`
	}
)
