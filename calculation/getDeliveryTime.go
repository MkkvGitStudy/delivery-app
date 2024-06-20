package calculation

import (
	"delivery-app/input"
	"delivery-app/output"
	"delivery-app/pkg"
	"delivery-app/util"
	"fmt"
	"math"
)

func GetDeliveryTime() {
	// Internally checks delivery price as well
	basePrice, noOfPackages := getBaseCostAndNoOfPkgs()
	pkgs := GetDeliveryPriceDiscount(basePrice, noOfPackages)
	maxWeight := GetMaxWeightOfPkgs(pkgs)
	speed, limit, vechiles := getDeliveryData()
	limit = validateMaxWeight(maxWeight, limit)
	CalculateDeliveryTime(pkgs, speed, limit, vechiles)
	output.OutPutDeliveryTimeAndPrice(pkgs)
}

func CalculateDeliveryTime(pkgs []pkg.Package, speed, limit float64, vechiles int) {

	updatedPkgSlice := pkgs
	var vehicleList = make([]float64, int(vechiles))

	for len(updatedPkgSlice) > 0 {

		nextShipmentList := GetNextShipmentsList(updatedPkgSlice, limit)
		nextDelivery := GetMinDistShipment(nextShipmentList, pkgs)
		// next vechile availabilty time
		nextAvailableSlot, nextAvailableSlotInd := util.Min(vehicleList...)
		// Total trip duration for all the pacakges in this shipment
		tripDuration := 0.0

		for _, pkgInd := range nextDelivery {

			time := util.Round(pkgs[pkgInd].Distance / speed)
			pkgs[pkgInd].DeliveryTime = util.Round(time + nextAvailableSlot)
			tripDuration = math.Max(tripDuration, pkgs[pkgInd].DeliveryTime)
		}
		updatedPkgSlice = []pkg.Package{}
		for _, v := range pkgs {
			// If not calculated delivery time will be -1
			if v.DeliveryTime == -1 {
				updatedPkgSlice = append(updatedPkgSlice, v)
			}
		}
		//Updatig availablity of current vehicle to twice the total shipment duration
		vehicleList[nextAvailableSlotInd] = 2 * tripDuration
	}
}

func GetNextShipmentsList(pkgList []pkg.Package, maxWeight float64) [][]int {
	nextShipments := [][]int{}
	n := len(pkgList)
	var curMaxWeight float64
	subsetMap := make(map[float64][][]pkg.Package)

	// check through all possible subsets
	for i := 1; i < (1 << n); i++ {
		var subset []pkg.Package
		totalWeight := 0.0

		for j := 0; j < n; j++ {

			//Checks if the item is included in the current subset
			if i&(1<<j) != 0 {
				subset = append(subset, pkgList[j])
				totalWeight += pkgList[j].PkgWeight
			}
		}

		if totalWeight <= maxWeight && totalWeight >= curMaxWeight {
			curMaxWeight = totalWeight
			subsetMap[curMaxWeight] = append(subsetMap[curMaxWeight], subset)
		}
	}

	// creates a list of subsets indices with current max weight for further processing
	for _, subset := range subsetMap[curMaxWeight] {
		temp := []int{}
		for _, elem := range subset {
			temp = append(temp, elem.Index)
		}
		nextShipments = append(nextShipments, temp)
	}
	return nextShipments
}

// Returns the minimum distance shipment from the current shipment list
func GetMinDistShipment(shipmentList [][]int, pkgs []pkg.Package) []int {
	if len(shipmentList) == 1 {
		return shipmentList[0]
	}

	minDistanceInd := 0
	minDistance := 0.0
	for ind, v := range shipmentList {
		distance := 0.0
		for _, i := range v {
			distance = math.Max(distance, pkgs[i].Distance)
		}
		if distance < minDistance || minDistance == 0.0 {
			minDistance = distance
			minDistanceInd = ind
		}
	}
	return shipmentList[minDistanceInd]
}

// The vehicles maxweight capacity should be greater than or equal to the maximum package weight
func validateMaxWeight(maxWeight, limit float64) float64 {
	for limit < maxWeight {
		fmt.Printf("Error the vechicle weight capacity should be atleast same as the maximum package weight %f\n", maxWeight)
		limit = input.GetPositiveFloatValue("Please enter the maximum weight for the vehicle : ", false, input.FloatInput{})
	}
	return limit
}

func GetMaxWeightOfPkgs(pkgs []pkg.Package) float64 {
	maxWeight := pkgs[0].PkgWeight
	for i := 1; i < len(pkgs); i++ {
		if pkgs[i].PkgWeight > maxWeight {
			maxWeight = pkgs[i].PkgWeight
		}
	}
	return maxWeight
}
