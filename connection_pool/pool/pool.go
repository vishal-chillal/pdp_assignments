package pool

import (
"database/sql"
_"github.com/mattn/go-sqlite3"
"sync"
)
var db_path string
var con_size,pull,push,req int
var connection_channel
var busy_connection = make([]int ,4)
var mu sync.Mutex

func Init(path string) int {
	db_path = path
	con_size = 4
	push,req = -1,-1
	connection_channel = make([]*sql.DB, 4)
	return Insert_conn(con_size)
}

func Insert_conn(num int) int {
	for i:=0 ; i< num ; i++{
		conn := make_connection(i)
		if(conn != nil){
			connection_channel[i] = conn

		}else{
			return 0
		}
	}
	return 1
}

func make_connection(i int) *sql.DB {
	db, err := sql.Open("sqlite3", db_path)
	if (err != nil) {
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
}

func Get_Conn() *sql.DB {
	req +=1
	pull = req%con_size
	if(check_available(pull)){
		busy_connection[pull] = 1
		return connection_channel[pull]
	}else{
		return nil
	}

}