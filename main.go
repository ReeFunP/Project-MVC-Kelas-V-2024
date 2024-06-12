package main

import (
	"Project_MVC/Model"
	"fmt"
)

func main() {
	Model.InsertMember(1, "Rifan", "Surabaya", 1000)
	Model.InsertMember(2, "Firmansyah", "Nganjuk", 200)
	Model.InsertMember(3, "Priono", "Pacitan", 1500)
	Model.InsertMember(4, "Icha", "Bojonegoro", 75)

	Model.SearchMember(1)

	member := Model.CariMember(2)
	fmt.Println("Test Cari Member : ", member.Data)

	Model.UpdateMember(3, "Abel", "Tengger Raya")

	Model.DeleteMember(4)

	Model.ReadAllMember()
	fmt.Println("--------------------------")
	if member != nil {
		for member.Next != nil {
			fmt.Println(member.Next.Data)
			member = member.Next
		}
	}
	fmt.Println("--------------------------")

}
