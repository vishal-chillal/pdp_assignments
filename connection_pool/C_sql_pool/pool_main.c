#include <stdio.h>
#include <sqlite3.h> 
#include "pool.c"

static int callback(void *data, int argc, char **argv, char **azColName){
   int i;
   fprintf(stderr, "%s: ", (const char*)data);
   for(i=0; i<argc; i++){
      printf("%s = %s\n", azColName[i], argv[i] ? argv[i] : "NULL");
   }
   printf("\n");
   return 0;
}

int main()
{
   sqlite3 *db;
   char *zErrMsg = 0;
   int size = 4,rc,check_flag;
   char *path = "./test.db";
   if(Init(size,path) == 1)
      printf("pool created\n");
   else
      printf("pool Failed\n");
   db = aquire();
   const char* data = "called";
   char *sql = "select * from abc;";
   rc = sqlite3_exec(db, sql, callback, (void*)data, &zErrMsg);
   if( rc != SQLITE_OK ){
      fprintf(stderr, "SQL error: %s\n", zErrMsg);
      sqlite3_free(zErrMsg);
   }else{
      fprintf(stdout, "Operation done successfully\n");
   }
   sqlite3_close(db);
   return 0;
}