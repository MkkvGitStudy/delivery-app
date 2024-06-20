package test

import (
	"delivery-app/calculation"
	"delivery-app/offers"
	"delivery-app/pkg"
	"reflect"
	"testing"
)

func TestCalculateDeliveryTime(t *testing.T) {
	type test struct {
		pkgs               []pkg.Package
		speed, weightLimit float64
		noOfVechiles       int
		resPkg             []pkg.Package
	}
	pkgsTest1 := []pkg.Package{
		{PkgId: "pk1", PkgWeight: 50.0, Distance: 30.0, Code: "OFR001", Index: 0, DeliveryTime: -1},
		{PkgId: "pk2", PkgWeight: 75.0, Distance: 125.0, Code: "OFFR0008", Index: 1, DeliveryTime: -1},
		{PkgId: "pk3", PkgWeight: 175.0, Distance: 100.0, Code: "OFR003", Index: 2, DeliveryTime: -1},
		{PkgId: "pk4", PkgWeight: 110.0, Distance: 60.0, Code: "OFR002", Index: 3, DeliveryTime: -1},
		{PkgId: "pk5", PkgWeight: 155.0, Distance: 95.0, Code: "test", Index: 4, DeliveryTime: -1},
	}
	resPkg1 := []pkg.Package{
		{PkgId: "pk1", PkgWeight: 50.0, Distance: 30.0, Code: "OFR001", Index: 0, DeliveryTime: 4.01, Discount: 0, Cost: 0},
		{PkgId: "pk2", PkgWeight: 75.0, Distance: 125.0, Code: "OFFR0008", Index: 1, DeliveryTime: 1.79, Discount: 0, Cost: 0},
		{PkgId: "pk3", PkgWeight: 175.0, Distance: 100.0, Code: "OFR003", Index: 2, DeliveryTime: 1.43, Discount: 0, Cost: 0},
		{PkgId: "pk4", PkgWeight: 110.0, Distance: 60.0, Code: "OFR002", Index: 3, DeliveryTime: 0.86, Discount: 0, Cost: 0},
		{PkgId: "pk5", PkgWeight: 155.0, Distance: 95.0, Code: "test", Index: 4, DeliveryTime: 4.22, Discount: 0, Cost: 0},
	}
	tests := []test{{pkgsTest1, 70, 200, 2, resPkg1}}

	for _, tc := range tests {
		calculation.CalculateDeliveryTime(tc.pkgs, tc.speed, tc.weightLimit, tc.noOfVechiles)
		if !reflect.DeepEqual(tc.pkgs, tc.resPkg) {
			t.Errorf("TestCalculateDeliveryTime error, wanted %v, got %v", tc.pkgs, tc.resPkg)
		}
	}
}

func TestCalculateDiscount(t *testing.T) {
	type test struct {
		name        string
		total       float64
		pkg         pkg.Package
		offerCodes  map[string]offers.Offer
		expectedRes float64
	}
	codes := map[string]offers.Offer{
		"OFR001": {
			DistanceCriteria: offers.Criteria{Min: 0, Max: 200},
			WeightCriteria:   offers.Criteria{Min: 70, Max: 200},
			Discount:         10,
		},
		"OFR002": {
			DistanceCriteria: offers.Criteria{Min: 50, Max: 150},
			WeightCriteria:   offers.Criteria{Min: 100, Max: 250},
			Discount:         7,
		},
	}
	tests := []test{
		{"no offer case", 200, pkg.Package{PkgId: "pk1", PkgWeight: 5, Distance: 5, Code: "OFR001"}, codes, 200},
		{"offer applied", 700, pkg.Package{PkgId: "pk2", PkgWeight: 100, Distance: 100, Code: "OFR001"}, codes, 630},
	}
	for _, tc := range tests {
		res := calculation.CalculateDiscount(tc.total, tc.pkg, tc.offerCodes)
		if res != tc.expectedRes {
			t.Errorf("%s - TestCalculateDiscount error, expected %f and got %f", tc.name, tc.expectedRes, res)
		}
	}
}

func TestGetNextShipmentsList(t *testing.T) {
	tests := []struct {
		pkgList   []pkg.Package
		maxWeight float64
		expected  [][]int
	}{
		{
			pkgList: []pkg.Package{
				{PkgId: "pkg1", PkgWeight: 10, Index: 0},
				{PkgId: "pkg2", PkgWeight: 20, Index: 1},
				{PkgId: "pkg3", PkgWeight: 30, Index: 2},
			},
			maxWeight: 30,
			expected:  [][]int{{0, 1}, {2}},
		},
		{
			pkgList: []pkg.Package{
				{PkgId: "pkg1", PkgWeight: 10, Index: 0},
				{PkgId: "pkg2", PkgWeight: 20, Index: 1},
				{PkgId: "pkg3", PkgWeight: 15, Index: 2},
			},
			maxWeight: 35,
			expected:  [][]int{{1, 2}},
		},
		{
			pkgList: []pkg.Package{
				{PkgId: "pkg1", PkgWeight: 5, Index: 0},
				{PkgId: "pkg2", PkgWeight: 7, Index: 1},
				{PkgId: "pkg3", PkgWeight: 3, Index: 2},
				{PkgId: "pkg4", PkgWeight: 2, Index: 3},
			},
			maxWeight: 10,
			expected:  [][]int{{1, 2}, {0, 2, 3}},
		},
		{
			pkgList: []pkg.Package{
				{PkgId: "pkg1", PkgWeight: 5, Index: 0},
				{PkgId: "pkg2", PkgWeight: 5, Index: 1},
				{PkgId: "pkg3", PkgWeight: 5, Index: 2},
			},
			maxWeight: 5,
			expected:  [][]int{{0}, {1}, {2}},
		},
	}

	for _, test := range tests {
		result := calculation.GetNextShipmentsList(test.pkgList, test.maxWeight)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For pkgList %v with maxWeight %f, expected %v but got %v", test.pkgList, test.maxWeight, test.expected, result)
		}
	}
}

func TestGetMinDistShipment(t *testing.T) {
	pkgs := []pkg.Package{
		{PkgId: "pkg1", PkgWeight: 10, Distance: 100, Index: 0},
		{PkgId: "pkg2", PkgWeight: 20, Distance: 200, Index: 1},
		{PkgId: "pkg3", PkgWeight: 15, Distance: 150, Index: 2},
		{PkgId: "pkg4", PkgWeight: 25, Distance: 250, Index: 3},
	}

	tests := []struct {
		shipmentList [][]int
		expected     []int
	}{
		{
			shipmentList: [][]int{{0, 1}, {2, 3}},
			expected:     []int{0, 1},
		},
		{
			shipmentList: [][]int{{1, 2}, {0, 3}},
			expected:     []int{1, 2},
		},
		{
			shipmentList: [][]int{{0, 2}, {1, 3}},
			expected:     []int{0, 2},
		},
		{
			shipmentList: [][]int{{0}, {1}, {2}},
			expected:     []int{0},
		},
	}

	for _, test := range tests {
		result := calculation.GetMinDistShipment(test.shipmentList, pkgs)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For shipmentList %v, expected %v but got %v", test.shipmentList, test.expected, result)
		}
	}
}

func TestGetMaxWeightOfPkgs(t *testing.T) {
	tests := []struct {
		pkgs     []pkg.Package
		expected float64
	}{
		{
			pkgs: []pkg.Package{
				{PkgId: "pkg1", PkgWeight: 10},
				{PkgId: "pkg2", PkgWeight: 20},
				{PkgId: "pkg3", PkgWeight: 15},
			},
			expected: 20,
		},
		{
			pkgs: []pkg.Package{
				{PkgId: "pkg1", PkgWeight: 10},
			},
			expected: 10,
		},
	}

	for _, test := range tests {
		result := calculation.GetMaxWeightOfPkgs(test.pkgs)
		if result != test.expected {
			t.Errorf("For pkgs %v, expected %f but got %f", test.pkgs, test.expected, result)
		}
	}
}
