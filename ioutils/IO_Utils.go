package ioutils

import("io/ioutil")

func check_error(err error){
	if err!=nil{
		panic(err)
	}
}

func Read_File(path string) (data []byte){
	data,err:=ioutil.ReadFile(path)
	check_error(err)
	return
}

func Write_File(pathToSave string,data []byte){
	err:=ioutil.WriteFile(pathToSave,data,0644)
	check_error(err)
}

