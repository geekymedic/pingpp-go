package royalty

import (
	"fmt"
	"log"
	"net/url"

	pingpp "github.com/geekymedic/pingpp-go/pingpp"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

func New(orderId string, params pingpp.RoyaltyCreateParams) (*pingpp.RoyaltyCreateReply, error) {
	return getC().New(orderId, params)
}

func (c Client) New(orderId string, params pingpp.RoyaltyCreateParams) (*pingpp.RoyaltyCreateReply, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("UserParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of create user is :\n %v\n ", string(paramsString))
	}

	royaltyReply := &pingpp.RoyaltyCreateReply{}
	err := c.B.Call("POST", fmt.Sprintf("/orders/%s/royalty_datas", orderId), c.Key, nil, paramsString, royaltyReply)
	return royaltyReply, err
}

func BatchUpdate(params *pingpp.RoyaltyBatchUpdateParams) (*pingpp.RoyaltyList, error) {
	return getC().BatchUpdate(params)
}

func (c Client) BatchUpdate(params *pingpp.RoyaltyBatchUpdateParams) (*pingpp.RoyaltyList, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("RoyaltyBatchUpdateParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of create user is :\n %v\n ", string(paramsString))
	}

	royaltyList := &pingpp.RoyaltyList{}
	err := c.B.Call("PUT", "royalties", c.Key, nil, paramsString, royaltyList)
	return royaltyList, err
}

func Get(royaltyId string) (*pingpp.Royalty, error) {
	return getC().Get(royaltyId)
}

func (c Client) Get(royaltyId string) (*pingpp.Royalty, error) {
	royalty := &pingpp.Royalty{}

	err := c.B.Call("GET", fmt.Sprintf("/royalties/%s", royaltyId), c.Key, nil, nil, royalty)
	return royalty, err
}

func List(params *pingpp.PagingParams) (*pingpp.RoyaltyList, error) {
	return getC().List(params)
}

func (c Client) List(params *pingpp.PagingParams) (*pingpp.RoyaltyList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	royaltyList := &pingpp.RoyaltyList{}
	err := c.B.Call("GET", "royalties", c.Key, body, nil, royaltyList)
	return royaltyList, err
}
