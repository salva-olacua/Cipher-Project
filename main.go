package main

import ("fmt";"Cipher_Project/xor";"Cipher_Project/ioutils")

func main(){
	var path string
	var key string
	var data string

	//fmt.Println("Ingrese el texto a cifar: ")
	//fmt.Scan(&data)

	fmt.Println("Ingrese la ruta del archivo a encriptar")
	fmt.Scan(&path)

	data=string(ioutils.Read_File(path))

	fmt.Println("-----------------------------")
	fmt.Println("Ingrese su clave: ")
	fmt.Scan(&key)

	fmt.Println("-----------------------------")
	encrypted:=xor.XorToAllElements(data,key)
	fmt.Println("Encrypted",encrypted)

	decrypted:=xor.XorToAllElements(encrypted,key)
	fmt.Println("Decrypted",decrypted)

	ioutils.Write_File(`C:\Practica_en_GO\encrypted.txt`,[]byte(encrypted))
}