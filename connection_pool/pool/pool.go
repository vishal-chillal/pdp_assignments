package pool

import (
"database/sql"
"fmt"
_"github.com/mattn/go-sqlite3"
"logger"
"sync"
"os"
"strconv"
)
var db_path string
var con_size,pull,push,req int
var connection_channel = make([]*sql.DB, 4)
var busy_connection = make([]int ,4)
var mu sync.Mutex

func Init(path string) {
	os.Remove("pool_log.text")
	logger.Init("pool_log.text")
	db_path = path
	con_size = 4
	push,req = -1,-1
	if (Insert_conn(con_size) == 1){
		fmt.Println("SUCCESS")
		logger.Info("generated the connection")
	}else{
		fmt.Println("FAILED")
		logger.Error("generation of pool is failed")
	}
}

func Insert_conn(num int) int {
	fmt.Println(con_size)	
	for i:=0 ; i< num ; i++{
		conn := make_connection(i)
		if(conn != nil){
			connection_channel[i] = conn
			fmt.Println("connection "+strconv.Itoa(i)+" done successfully")
			logger.Info("connection "+strconv.Itoa(i)+" done successfully")

		}else{
			fmt.Println("connection "+strconv.Itoa(i)+" Failed")
			logger.Error("connection "+strconv.Itoa(i)+" Failed")
			return 0
		}
	}
	fmt.Println("database pool array = ", connection_channel)
	return 1
}

func make_connection(i int) *sql.DB {
	db, err := sql.Open("sqlite3", db_path)
	if (err != nil) {
		db.Close()
		return nil
	}
	return db
}

func check_available(position int) bool {
	
	mu.Lock()
	if(busy_connection[position] == 1){
		mu.Unlock()
		return false
	}
	mu.Unlock()
	return true
}
func Release_Conn(db *sql.DB) {
	push += 1
	position := push%con_size
	busy_connection[position] = 0
	connection_channel[position] = db
	fmt.Println("connection "+strconv.Itoa(position)+" release")
	logger.Info("connection "+strconv.Itoa(position)+" release")
}

func Get_Conn() *sql.DB {
	req +=1
	pull = req%con_size
	if(check_available(pull)){
		busy_connection[pull] = 1
		fmt.Println("connection accuire")
		 logger.Info("connection "+strconv.Itoa(pull)+" accuire successfully")
		return connection_channel[pull]
	}else{
		fmt.Println("Pool is full")
		logger.Warning("Pool is fool")
		return nil
	}

}