package calculation

import (
	"delivery-app/filemanager"
	"delivery-app/offers"
	"delivery-app/output"
	"delivery-app/pkg"
	"delivery-app/util"
	"fmt"
)

// Gets the final price and prints the details
func GetDeliveryPrice() {
	basePrice, noOfPackages := getBaseCostAndNoOfPkgs()
	pkgs := GetDeliveryPriceDiscount(basePrice, noOfPackages)
	output.OutPutDiscountedPrice(pkgs)
}

func GetDeliveryPriceDiscount(basePrice float64, noOfPackages int) []pkg.Package {
	pkgs := getPackages(noOfPackages)
	var err error

	offerCodes, err := filemanager.ReadOfferCodesJson(util.FilePath)
	for i := 0; i < noOfPackages; i++ {
		totalPrice := pkgs[i].GetTotalPrice(basePrice)
		// Getting error means, unable to do file related operations for coupon code processing.
		// So skipping discount calculation for all iterations and notifying user
		if err == nil {
			pkgs[i].Cost = CalculateDiscount(totalPrice, pkgs[i], offerCodes)
		} else {
			pkgs[i].Cost = totalPrice
		}
		pkgs[i].Discount = totalPrice - pkgs[i].Cost
	}
	if err != nil {
		fmt.Println(err)
	}
	return pkgs
}

func CalculateDiscount(totalPrice float64, pkg pkg.Package, offerCodes map[string]offers.Offer) float64 {
	if details, ok := offerCodes[pkg.Code]; ok {
		if !(details.DistanceCriteria.Min <= pkg.Distance && pkg.Distance <= details.DistanceCriteria.Max) ||
			!(details.WeightCriteria.Min <= pkg.PkgWeight && pkg.PkgWeight <= details.WeightCriteria.Max) {
			return totalPrice
		}
		return totalPrice * (100 - details.Discount) / 100
	}
	return totalPrice
}
