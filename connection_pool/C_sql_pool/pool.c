#include <sqlite3.h> 
#include <malloc.h>
#include <stdio.h>
#include <string.h>
int pool_size, *arr, push=-1,pull=-1;
char *db_path;
sqlite3** pool;
sqlite3* getConn(){
   sqlite3 *db;
   char *zErrMsg = 0;
   int rc;
   rc = sqlite3_open(db_path, &db);
   if( rc ){
      return NULL;
   }else{
      return db;
   }
}

int Init(int size,char *path){
   pool_size = size;
   db_path = malloc(sizeof(strlen(path)));
   strcpy(db_path,path);
   arr = (int*)malloc(sizeof(int)*size);
   sqlite3* pool[pool_size];
   int i;
   for (i = 0; i <pool_size; ++i){  
      sqlite3 *db = getConn();      
      if(db){
       pool[i] = db;
    }else{
      return 0;
   }
}
return 1;
}

sqlite3 *aquire(){
   if(push >= pull)
      return 0;
   pull += 1;
   return pool[pull];
}

void release(sqlite3 *db){
   push += 1;

}
