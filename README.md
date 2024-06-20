# Everest engineering coding challenge

Kiki owns a delivery service which has N vehicles and delivery partners who would deliver the packages

## Running the program

```commands
    go mod tidy
    go run .
```

## App workflow

The user will be provided with 4 options in the CLI
1 - Calculate delivery cost 
2 - Calculate the delivery time 
3 - Add offer code 
4 - Show current offers

### Option 1

User will have to provide
- The base price
- The number of packages 
- The package details (package Id, weight, delivery distance and the coupon code)

total delivery cost is calculated as "basePrice + (pkg.PkgWeight * 10) + (pkg.Distance * 5)"
offer code is applied on the total delivery cost if applicable and the actual cost of the package is calculated by subtracting the discount amount from the total delivery cost
offercodes are stored in json format and new offercodes can be added.

### Option 2

User will have to input all the details that they entered in the first option and in addition need to provide the following
- The number of vehicles
- The speed of the vechicle 
- The maximum weight that vehicle can carry

The delivery time is estimated based on the following criterias
- Heavier packages are chosen among the available packages but below the maximum carriable capacity of the vehicle. 
- The shipment which can be delivered first has to be chosen when their weights are same.

Following assumptions are made for the delivery time calculation
- All vehicles travel at the same speed and in same route
- All destinations are covered in a single route

### Option 3

User can select this to add new offer codes. User will have to provide the following details
- The discount percentage for coupon
- The minimum distance and maximum distance values applicable for coupon
- The minimum distance and maximum weight applicable for coupon

### Option 4
This will list out all the offer codes

## Running the test cases

```commands
    go test ./test
```