package pingpp

type ContactParam struct {
	User                  string `json:"user"`                     // required
	AccNo                 string `json:"acc_no"`                   // 用户编号
	ContactType           int    `json:"contact_type"`             // required 联系人类型。1：第一紧急联系人、2：第二紧急联系人、3：第三紧急联系人、4、法人、5、财务人员（4、5必传）。
	ContactName           string `json:"contact_name"`             // required
	ContactCertType       string `json:"contact_cert_type"`        // required
	ContactCertNo         string `json:"contact_cert_no"`          // required
	ContactCertValidFrom  string `json:"contact_cert_valid_from"`  // required yyyy-mm-dd
	ContactCertValidUntil string `json:"contact_cert_valid_until"` // required
	ContactCertMobile     string `json:"contact_cert_mobile"`      // required
	ContactJobType        int
	ContactRelationShip   int
}

type Contact struct {
	User      string `json:"user"`
	LiveMode  bool   `json:"live_mode"`
	AccNo     string `json:"acc_no"`
	ContactNo string `json:"contact_no"`
}
