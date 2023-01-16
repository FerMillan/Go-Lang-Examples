package main

import (
	"awesomeProject/database"
	"awesomeProject/modules/api_calls"
	"awesomeProject/modules/aritmetic"
	"awesomeProject/modules/concurrency"
	"awesomeProject/modules/interfaces"
	"awesomeProject/modules/solveDoubts"
	"fmt"
)

// Use of a struct to construct interfaces functions
type T struct {
	S string
}

// Function coded since interfaz
func (t T) ShowString() {
	fmt.Println(t.S)
}

// Function coded since interfaz
func (t T) Function() {
	fmt.Println(t.S + " - Function to add")
}

// Menu function can only use in the package main
func menu() {
	var metodo string

	fmt.Println("MENU DE FUNCIONES:\n" +
		"0 -- Example to Sum aritmetic (sum)\n" +
		"1 -- Switch Doubt solution (switch)\n" +
		"2 -- Map Doubt solution (map)\n" +
		"3 -- Array doubt solution (array)\n" +
		"4 -- Example of interfaces (interfaz)\n" +
		"5 -- Example of concurrency (hilos)\n" +
		"6 -- Example of database (database)\n" +
		"7 -- Example of api get (get)\n" +
		"8 -- Example of api post (post)\n" +
		"9 -- Way to check if a key in a map exist (map_exist)\n" +
		"salir -- Exit of menu\n" +
		"\nIngrese una entrada:")
	fmt.Scanln(&metodo)
	fmt.Println("\n||||  Resultado de la llamda a la función", metodo, "  ||||\n")

	// Switch to select some method
	switch metodo {
	case "sum":
		fallthrough
	case "0":
		fmt.Println(aritmetic.Sum(4, 5))
	case "switch":
		fallthrough
	case "1":
		solveDoubts.SwitchWithFallthrough(3)
	case "map":
		fallthrough
	case "2":
		solveDoubts.GetKeysAndValuesOfAMap()
	case "array":
		fallthrough
	case "3":
		solveDoubts.FindIfExistSomeElemInArray(11)
	case "interfaz":
		fallthrough
	case "4":
		var i interfaces.Interface = &T{"hola"}
		i.Function()
		i.ShowString()
	case "hilos":
		fallthrough
	case "5":
		concurrency.Xplosion()
		c := make(chan int, 10)
		go concurrency.Fibonacci(cap(c), c)
		fmt.Println("\nFibonacci con hilos:\n")
		for i := range c {
			fmt.Println(i)
		}
	case "database":
		fallthrough
	case "6":
		menuToDatabase()
	case "salir":
		fmt.Println("Saliendo ...")
	case "7":
		fallthrough
	case "get":
		fmt.Println(api_calls.Get())
	case "8":
		fallthrough
	case "post":
		fmt.Println(api_calls.Post())
	case "9":
		fallthrough
	case "map_exist":
		solveDoubts.ExistKeyInMap()
	default:
		fmt.Println("Función no añadida Intentelo de nuevo\n----------------------\n\n")
		menu() // Recursive call to access again to the menu if this is required
	}
}

func menuToDatabase() {
	var metodo string

	fmt.Println("MENU DE FUNCIONES (base de datos):\n" +
		"0 -- SELECT COUNT(*) FROM tbl_users\n" +
		"1 -- SELECT * FROM tbl_users LIMIT 1 (ORM)\n" +
		"2 -- SELECT * FROM tbl_users LIMIT 1 (native)\n" +
		"3 -- SELECT * FROM tbl_users WHERE email = 'vikram@undostres.com.mx' LIMIT 1\n" +
		"4 -- " + "SELECT * " +
		"FROM tbl_referralcodes tr " +
		"JOIN tbl_users tu " +
		"ON tr.user_id = tu.id " +
		"WHERE tu.mobile IS NULL " +
		"AND tu.facebookid = '' " +
		"LIMIT 1\n" +
		"5 -- SELECT * FROM referralcode_rejection_reason (All rows)\n" +
		"6 -- SELECT * FROM tbl_referralcodes WHERE user_id = 2647712 (Type safe version)")
	fmt.Println("\nIngrese una entrada:")
	fmt.Scanln(&metodo)
	fmt.Println("\n||||  Resultado de la llamda a la función", metodo, "  ||||\n")

	switch metodo {
	case "0":
		database.ExampleSelectCountToUsers()
	case "1":
		database.ExampleSelectLimitORM()
	case "2":
		database.ExampleNativeQuery()
	case "3":
		database.ExampleMoreHardQueryUseORM()
	case "4":
		database.ExampleStillMoreHardQueryUseORM()
	case "5":
		database.ExampleEveryRowInATable()
	case "6":
		database.ExampleQueryToUseSafeConstants()
	default:
		fmt.Println("No es una opción válida")
	}
}
