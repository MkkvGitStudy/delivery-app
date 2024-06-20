package output

import (
	"delivery-app/offers"
	"delivery-app/pkg"
	"fmt"
	"os"
	"text/tabwriter"
)

func OutPutDiscountedPrice(pkgs []pkg.Package) {
	fmt.Println("\nShipment discount and price")
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(writer, "\nID\tDiscount\tTotalCost")
	for _, pkg := range pkgs {
		fmt.Fprintf(writer, "%s\t%.2f\t%.2f\t\n", pkg.PkgId, pkg.Discount, pkg.Cost)
	}
	writer.Flush()
}

func OutPutDeliveryTimeAndPrice(pkgs []pkg.Package) {
	fmt.Println("\nShipment discount, price & Delivey time in Hour")
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(writer, "\nID\tDiscount\tTotalCost\tDeliveyTime")
	for _, pkg := range pkgs {
		fmt.Fprintf(writer, "%s\t%.2f\t%.2f\t%.2f\t\n", pkg.PkgId, pkg.Discount, pkg.Cost, pkg.DeliveryTime)
	}
	writer.Flush()
}

func OutputOfferCodes(offerCodes map[string]offers.Offer) {
	fmt.Println("\nCurrently available offer code details")
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(writer, "\nCode\tDiscount\tMinimumDistance\tMaximumDistance\tMinimumWeight\tMaximumWeight")
	for code, offer := range offerCodes {
		fmt.Fprintf(writer, "%s\t%.2f\t%.2f\t%.2f\t%.2f\t%.2f\t\n", code, offer.Discount, offer.DistanceCriteria.Min,
			offer.DistanceCriteria.Max, offer.WeightCriteria.Min, offer.WeightCriteria.Max)
	}
	writer.Flush()
}
