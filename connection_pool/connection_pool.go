package main

import (
"pool"
"os"
)

func main() {
	//the pool size is of FOUR CONNECTIONS
	os.Remove("./myDb.db")
	pool.Init("./myDb.db")
	
	check1 := pool.Get_Conn()
	check2 := pool.Get_Conn()
	check3 := pool.Get_Conn()

	pool.Release_Conn(check1)
	pool.Get_Conn()
	pool.Get_Conn()
	pool.Get_Conn()

	pool.Release_Conn(check3)
	pool.Release_Conn(check2)
	pool.Get_Conn()
	pool.Get_Conn()
}