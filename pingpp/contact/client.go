package contact

import (
	"fmt"
	"log"

	"github.com/geekymedic/pingpp-go/pingpp"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

func New(appId string, params *pingpp.ContactParam) (*pingpp.Contact, error) {
	return getC().New(appId, params)
}

func (c Client) New(appId string, params *pingpp.ContactParam) (*pingpp.Contact, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("OrderCreateParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of create order is :\n %v\n ", string(paramsString))
	}

	contact := &pingpp.Contact{}
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/users/upload_pic", appId), c.Key, nil, paramsString, contact)
	return contact, err
}
