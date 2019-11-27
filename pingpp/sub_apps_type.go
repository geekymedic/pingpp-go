package pingpp

type (
	SubAppParams struct {
		DisplayName string                 `json:"display_name"`
		User        interface{}            `json:"user"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
		Description string                 `json:"description,omitempty"`
	}
	SubApp struct {
		ID               string                 `json:"id"`
		Object           string                 `json:"object"`
		App              string                 `json:"app"`
		Created          string                 `json:"created"`
		DisplayName      string                 `json:"display_name"`
		Account          string                 `json:"account"`
		Description      string                 `json:"description"`
		Metadata         map[string]interface{} `json:"metadata"`
		AvailableMethods string                 `json:"available_methods"`
		User             User                   `json:"user"`
		Level            int                    `json:"level"`
		ParentApp        string                 `json:"parent_app"`
	}

	SubAppUpdateParams struct {
		DisplayName string                 `json:"display_name,omitempty"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
		Description string                 `json:"description,omitempty"`
		ParentApp   string                 `json:"parent_app,omitempty"`
	}

	// SubAppList 商户列表
	SubAppList struct {
		ListMeta
		Values []*SubApp `json:"data"`
	}
)

type DepositoryUser struct {
	Id           string       `json:"id"`     //  用户 id，1~64 位。
	Name         string       `json:"name"`   // 公司名称，1~32 位，中英文长度都算 1（后面添加结算账号需与此名字一致）。
	Mobile       string       `json:"mobile"` // 法人手机号，1~20 位。
	IdentifyInfo IdentityInfo // 证件信息
	Address      string       `json:"address,omitempty"` // 办公地址
	Email        string       `json:"email,omitempty"`
	Gender       string       `json:"gender"` // MALE FEMALE
}
