package calculation

import (
	"delivery-app/input"
	"delivery-app/pkg"
)

func getPackages(noOfPackages int) []pkg.Package {
	pkgs := make([]pkg.Package, noOfPackages)
	for i := 0; i < noOfPackages; i++ {
		pkg := pkg.New()
		pkg.GetPackageData()
		pkg.Index = i
		pkg.DeliveryTime = -1
		pkgs[i] = pkg
	}
	return pkgs
}

func getBaseCostAndNoOfPkgs() (float64, int) {
	var basePrice float64
	var noOfPackages int
	floatIp := input.FloatInput{}
	intIp := input.IntInput{}
	basePrice = floatIp.GetFloatInput("Please enter the base price value : ")
	noOfPackages = intIp.GetIntInput("Please enter the number of packages : ")
	return basePrice, noOfPackages
}

func getDeliveryData() (float64, float64, int) {
	floatIp := input.FloatInput{}
	intIp := input.IntInput{}
	vechiles := input.GetPositiveIntValue("Please enter the number of vehicles : ", intIp)
	speedLimit := input.GetPositiveFloatValue("Please enter the speed of the vechicle : ", false, floatIp)
	weightLimit := input.GetPositiveFloatValue("Please enter the maximum weight for the vehicle : ", false, floatIp)
	return speedLimit, weightLimit, vechiles
}
