package main

import (
	"fmt"
	"os"
	"strconv"
)

type students struct {
	data []student
}

type student struct {
	id      int
	name    string
	address string
	job     string
	reason  string
}

type get interface {
	getStudent(i int) student
}

func (s students) getStudent(i int) student {
	for _, v := range s.data {
		if v.id == i {
			return v
		}
	}
	return student{}
}

// 	id      int
// 	name    string
// 	address string
// 	job     string
// 	reason  string

func main() {
	student_id := os.Args[1]
	i, err := strconv.Atoi(student_id)

	var murid = []student{
		{
			id:      1,
			name:    "nama1",
			address: "alamat1",
			job:     "wirausahawan",
			reason:  "jati diri",
		},
		{
			id:      2,
			name:    "nama2",
			address: "alamat2",
			job:     "pegawai",
			reason:  "cari jodoh",
		},
	}

	// Condition if argument not integer
	if err != nil {
		fmt.Println("Argument must be Integer!")
		return
	}

	// Condition if argument outside data
	if i > len(murid) {
		fmt.Println("ID not Found")
		return
	}

	var student_list get = students{data: murid}

	print(i, student_list)
}

func print(i int, g get) {
	v := g.getStudent(i)
	fmt.Printf("Nama\t: %s\nAddress\t: %s\nJob\t: %s\nReason\t: %s\n", v.name, v.address, v.job, v.reason)
	// if v.id != 0 {
	// 	fmt.Printf("Nama\t: %s\nAddress\t: %s\nJob\t: %s\nReason\t: %s\n", v.name, v.address, v.job, v.reason)
	// } else {
	// 	fmt.Println("ID not Found")
	// }

}
