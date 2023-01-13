package solveDoubts

import (
	"awesomeProject/modules/auxiliars"
	"fmt"
	"golang.org/x/exp/maps"
	"reflect"
)

// SwitchWithFallthrough
// Método utilizado para solucionar el problema con la falta de break
func SwitchWithFallthrough(status int) {

	switch status := status; {
	case status < 2:
		fmt.Println("Status is not done")
	case status < 4:
		fmt.Println("Status is done")
		fallthrough
	case status < 6:
		fmt.Println("Status is done but has some issues")
		fallthrough
	case status == 7:
		fmt.Println("Status Extra")
	default:
		fmt.Println("Status default")
	}
}

// GetKeysAndValuesOfAMap
// Método utilizado para solucionar la forma de obtener las keys de un 'Map'
func GetKeysAndValuesOfAMap() {
	map_aux := map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": "dasd",
	}
	valueOf := reflect.ValueOf(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	})
	keys := valueOf.MapKeys()

	keys_2 := maps.Keys(map_aux)

	fmt.Println("Map: ", map_aux)
	fmt.Println("Value Of: ", valueOf)
	fmt.Println("Keys from Map: ", keys)
	fmt.Println("Keys another way: ", keys_2)
}

// Way to check if a key exist
func ExistKeyInMap() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	key_to_check := "asd"

	if a, ok := m[key_to_check]; ok {
		fmt.Println(m[key_to_check])
		fmt.Println("Value from the key \""+key_to_check+"\": ", a, "\nExist?: ", ok)
	}

}

// FindIfExistSomeElemInArray
// Método utilizado para solucionar el conocer la existencia de un elemento en un arreglo
func FindIfExistSomeElemInArray(find int) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Get Index of Element:", auxiliars.Index(arr, find))
	fmt.Println("Get if Exist an Element:", auxiliars.IsExist(arr, find))
}
