package main

import "fmt"

type Segment struct {
	tattoo     int8
	vendorName string
}

type TravelInfo struct {
	travelled  bool
	travelDate string
}

type travelleable interface {
	Travel()
}

func (tf *TravelInfo) Travel(date string) {
	tf.travelled = true
	tf.travelDate = date
}

// Type embedding - Composition over inheritance
type AirSegment struct {
	Segment
	TravelInfo
}

type CarSegment struct {
	Segment
}

type HotelSegment struct {
	Segment
}

func main() {
	air := AirSegment{Segment{}, TravelInfo{travelled: false}}
	fmt.Println(air.tattoo)
	fmt.Println(air.vendorName)
	fmt.Println(air.travelled)
	fmt.Println(air.travelDate)
	// Duck typing - If it walks like a duck and quacks like a duck, then it must be a duck
	// In this case, if AirSegment has a travelInfo field, then it must be travelleable
	// Nice one
	air.Travel("2023-12-25")
	fmt.Println(air.travelled)
}
