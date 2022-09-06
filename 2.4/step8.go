package main

import "fmt"

type Fatima struct {
	On          bool
	Ammo, Power int
}

func (u *Fatima) Shoot() bool {
	if u.On == false || u.Ammo == 0 {
		return false
	}
	u.Ammo--
	return true
}
func (u *Fatima) RideBike() bool {
	if u.On == false || u.Power == 0 {
		return false
	}
	u.Power--
	return true
}
func main() {
	testStruct := new(Fatima)
	fmt.Println(testStruct)
}
