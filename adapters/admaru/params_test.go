package admaru

import (
	"encoding/json"
	"github.com/prebid/prebid-server/openrtb_ext"
	"testing"
)

// This file actually intends to test static/bidder-params/admaru.json
//
// These also validate the format of the external API: request.imp[i].ext.admaru

// TestValidParams makes sure that the admaru schema accepts all imp.ext fields which we intend to support.
func TestValidParams(t *testing.T) {
	validator, err := openrtb_ext.NewBidderParamsValidator("../../static/bidder-params")
	if err != nil {
		t.Fatalf("Failed to fetch the json-schemas. %v", err)
	}

	for _, validParam := range validParams {
		if err := validator.Validate(openrtb_ext.BidderAdmaru, json.RawMessage(validParam)); err != nil {
			t.Errorf("Schema rejected admaru params: %s", validParam)
		}
	}
}

// TestInvalidParams makes sure that the admaru schema rejects all the imp.ext fields we don't support.
func TestInvalidParams(t *testing.T) {
	validator, err := openrtb_ext.NewBidderParamsValidator("../../static/bidder-params")
	if err != nil {
		t.Fatalf("Failed to fetch the json-schemas. %v", err)
	}

	for _, invalidParam := range invalidParams {
		if err := validator.Validate(openrtb_ext.BidderAdmaru, json.RawMessage(invalidParam)); err != nil {
			t.Errorf("Schema allowed unexpected params: %s", invalidParam)
		}
	}
}

var validParams = []string{
	`{"adspaceid": "ad-83444226E44368D1E32E49EEBE6D29","pubid": "par-2EDDB423AA24474188B843EE4842932"}`,
}

var invalidParams = []string{
	`{"adspaceid": "","pubid": ""}`,
}
