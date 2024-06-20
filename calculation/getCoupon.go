package calculation

import (
	"delivery-app/filemanager"
	"delivery-app/input"
	"delivery-app/offers"
	"delivery-app/output"
	"delivery-app/util"
	"fmt"
)

func fetchNewCouponData() (offers.Offer, string) {
	fp := input.FloatInput{}
	discount := input.GetPositiveFloatValue("Please enter the discount percentage for coupon : ", true, fp)
	minDistance := input.GetPositiveFloatValue("Please enter the minimum distance for coupon : ", true, fp)
	maxDistance := input.GetPositiveFloatValue("Please enter the maximum distance for coupon : ", true, fp)
	minWeight := input.GetPositiveFloatValue("Please enter the minimum weight for coupon : ", true, fp)
	maxWeight := input.GetPositiveFloatValue("Please enter the maximum weight for coupon : ", true, fp)
	code := input.GetCouponCodeInput("Please enter the coupon code : ")
	offer := offers.New(minWeight, maxWeight, minDistance, maxDistance, discount)
	return offer, code
}

func AddNewOfferCode() {
	offer, code := fetchNewCouponData()
	offerCodes, err := filemanager.ReadOfferCodesJson(util.FilePath)
	if err != nil {
		fmt.Println("Unable to update the new offer code -", err)
	}
	// Adding new offer to code to existing offers
	offerCodes[code] = offer
	err = filemanager.WriteOfferCodesJson(offerCodes, util.FilePath)
	if err != nil {
		fmt.Println("Unable to update the new offer code -", err)
	}
}

func GetOfferCodes() {
	offerCodes, err := filemanager.ReadOfferCodesJson(util.FilePath)
	if err != nil {
		fmt.Println("Unable to retrieve the offer codes -", err)
	}
	output.OutputOfferCodes(offerCodes)
}
