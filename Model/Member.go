package Model

import (
	"Project_MVC/Database"
	"Project_MVC/Node"
	"fmt"
)

func InsertMember(id int, nama string, alamat string, point float32) {
	newDataMember := Node.TableMember{
		Data: Node.Member{Id: id, Nama: nama, Alamat: alamat, Point: point},
		Next: nil,
	}
	var tempLL *Node.TableMember
	tempLL = &Database.DBmember
	if Database.DBmember.Next == nil {
		Database.DBmember.Next = &newDataMember
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = &newDataMember
	}
}

func ReadAllMember() *Node.TableMember {
	var tempLL *Node.TableMember
	tempLL = &Database.DBmember
	if Database.DBmember.Next == nil {
		return nil
	} else {
		return tempLL
	}
}

func SearchMember(id int) {
	var tempLL *Node.TableMember
	tempLL = &Database.DBmember
	cek := false
	if Database.DBmember.Next == nil {
		fmt.Println("                    ")
		fmt.Println("DATA MEMBER KOSONG")
		fmt.Println("                    ")
	} else {
		for tempLL != nil {
			if id == tempLL.Data.Id {
				cek = true
				fmt.Println("=================================")
				fmt.Println("Nomor ID member : ", tempLL.Data.Id)
				fmt.Println("Nama member : ", tempLL.Data.Nama)
				fmt.Println("Alamat member : ", tempLL.Data.Alamat)
				fmt.Println("Point member : ", tempLL.Data.Point)
				fmt.Println("=================================")
			}
			tempLL = tempLL.Next
		}
		if cek != true {
			fmt.Println("ID tidak ada pada DATA MEMBER")
		}
	}
}

func CariMember(id int) *Node.TableMember {
	var tempLL *Node.TableMember
	tempLL = &Database.DBmember
	cek := false
	if Database.DBmember.Next == nil {
		return nil
	} else {
		for tempLL != nil {
			if id == tempLL.Data.Id {
				cek = true
				return tempLL
			}
			tempLL = tempLL.Next
		}
		if cek != true {
			return nil
		}
	}
	return nil
}

func UpdateMember(id int, nama string, alamat string) {
	member := CariMember(id)
	if member != nil {
		member.Data.Nama = nama
		member.Data.Alamat = alamat
		fmt.Println("Update Berhasil")
	} else {
		fmt.Println("Tidak ada data yang dicari")
	}
}

func DeleteMember(id int) {
	var tempLL *Node.TableMember
	tempLL = &Database.DBmember
	if tempLL.Next == nil {
		//cek data kosong
		fmt.Println("data kosong")
	} else {
		for tempLL.Next != nil {
			if tempLL.Next.Data.Id == id {
				tempLL.Next = tempLL.Next.Next
				fmt.Println("Data Berhasil dihapus")
				return
			}
			tempLL = tempLL.Next
		}
	}
}
