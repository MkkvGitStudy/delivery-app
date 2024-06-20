package pkg

import "delivery-app/input"

type Package struct {
	PkgId        string
	PkgWeight    float64
	Distance     float64
	Code         string
	Cost         float64
	Discount     float64
	Index        int
	DeliveryTime float64
}

func New() Package {
	return Package{}
}

func (pkg Package) GetTotalPrice(basePrice float64) float64 {
	return basePrice + (pkg.PkgWeight * 10) + (pkg.Distance * 5)
}

func (pkg *Package) GetPackageData() {
	pkg.PkgId = input.GetStringInput("Please enter the package Id : ")
	pkg.PkgWeight = input.GetPositiveFloatValue("Please enter the package weight : ", false, input.FloatInput{})
	pkg.Distance = input.GetPositiveFloatValue("Please enter the delivery distance : ", false, input.FloatInput{})
	pkg.Code = input.GetCouponCodeInput("Please enter any coupon code if applicable : ")
}
