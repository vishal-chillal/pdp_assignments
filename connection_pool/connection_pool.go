package main

import (
"pool"
"fmt"
"os"
)

func main() {
	//the pool size is of FOUR CONNECTIONS
	os.Remove("./myDb.db")
	if (pool.Init("./myDb.db") == 1){
		fmt.Println("pool is initialize")
	}else{
		fmt.Println("Initialization failed;")
		return
	}

	check1 := pool.Get_Conn()
	if(check1 == nil){
		fmt.Println("Cant get connection from pool");
	}else{
		fmt.Println("connection successful");
	}
	check2 := pool.Get_Conn()
	if(check2 == nil){
		fmt.Println("Cant get connection from pool");
	}else{
		fmt.Println("connection successful");
	}
	check3 := pool.Get_Conn()
	if(check3 == nil){
		fmt.Println("Cant get connection from pool");
	}else{
		fmt.Println("connection successful");
	}
	// pool.Release_Conn(check1)
	// pool.Get_Conn()
	// pool.Get_Conn()
	// pool.Get_Conn()

	// pool.Release_Conn(check3)
	// pool.Release_Conn(check2)
	// pool.Get_Conn()
	// pool.Get_Conn()
}