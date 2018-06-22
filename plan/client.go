/*************************************************************************
 * 
 * █████╗ ██╗   ██╗ ██████╗██╗     
 * ██╔══██╗██║   ██║██╔════╝██║     
 * ███████║██║   ██║██║     ██║     
 * ██╔══██║╚██╗ ██╔╝██║     ██║     
 * ██║  ██║ ╚████╔╝ ╚██████╗███████╗
 * ╚═╝  ╚═╝  ╚═══╝   ╚═════╝╚══════╝
 * __________________
 * 
 *  [2012] - [2017] AVCL group 
 *  All Rights Reserved.
 * 
 * NOTICE:  All information contained herein is, and remains
 * the property of Avincel group and its suppliers,
 * if any.  The intellectual and technical concepts contained
 * herein are proprietary to Avincel group
 * and its suppliers and may be covered by U.S. and Foreign Patents,
 * patents in process, and are protected by trade secret or copyright law.
 * Dissemination of this information or reproduction of this material
 * is strictly forbidden unless prior written permission is obtained
 * from Avincel Group.
 */

package plan

import (
	"net/url"
	"github.com/hadv/go-chargebee"
	"strconv"
)

// Client is used to invoke /invoices APIs.
// https://apidocs.chargebee.com/docs/api/invoices
type Client struct {
	B   chargebee.Backend
	Key string
}


func (c Client) ListPlans(params *chargebee.ListPlansParams) ([]*chargebee.PlanList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Limit > 0 {
			body.Add("limit", strconv.Itoa(params.Limit))
		}
	}
	result := &chargebee.PlanResult{}
	err := c.B.Call("GET", "/plans", c.Key, body, result)
	if err != nil {
		return nil, err
	}
	invoice := result.List
	return invoice, nil
}

func ListPlans(params *chargebee.ListPlansParams) ([]*chargebee.PlanList, error) {
	return getC().ListPlans(params)
}


func getC() Client {
	return Client{
		B:   chargebee.NewBackend(chargebee.SiteName),
		Key: chargebee.Key,
	}
}



