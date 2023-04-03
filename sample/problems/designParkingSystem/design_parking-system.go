package designParkingSystem

type ParkingSystem struct {
	Parking1 [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{[3]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.Parking1[carType-1] > 0 {
		this.Parking1[carType-1]--
		return true
	} else {
		return false
	}
}
