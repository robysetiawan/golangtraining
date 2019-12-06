package main

import (
	"fmt"
	"time"
)

func hallo() {
	fmt.Println("Hello world routines")
}

func main() {
	go hallo()
	fmt.Println("main function")
	time.Sleep(1 * time.Second)
	//no 1
	// var N int
	// fmt.Print("Masukkan jumlah student : ")
	// fmt.Scan(&N)

	// var student day3.Student
	// for index := 0; index < N; index++ {
	// 	fmt.Print("Input ", index+1, " Student's Name : ")
	// 	var name string
	// 	fmt.Scan(&name)
	// 	student.Name = append(student.Name, name)

	// 	fmt.Print("Input ", index+1, " Student's Score : ")
	// 	var score float64
	// 	fmt.Scan(&score)
	// 	student.Score = append(student.Score, score)
	// }

	// min, max, average := student.Summary()
	// fmt.Println("Average Score : ", average)
	// fmt.Println("Min Score of Students : ", min)
	// fmt.Println("Max Score of Students : ", max)

	// no2
	// fmt.Print("[1] Encrypt\n[2] Decrypt\nChoose your menu ? ")
	// var menu int
	// fmt.Scan(&menu)

	// var student day3.Siswa

	// switch menu {
	// case 1:
	// 	fmt.Print("Input Students Name : ")
	// 	// var name string
	// 	fmt.Scan(&student.Name)

	// 	// student.SetName(name)
	// 	student.Encode()
	// 	fmt.Println(student.NameEncode)
	// case 2:
	// 	fmt.Print("Input Students Encoded Name : ")
	// 	// var encodedName string
	// 	fmt.Scan(&student.NameEncode)

	// 	// student.SetNameEncode(encodedName)
	// 	student.Decode()
	// 	fmt.Println(student.Name)
	// }

}
