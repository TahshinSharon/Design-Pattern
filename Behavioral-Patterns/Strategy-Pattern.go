package main

import "fmt"

type IDBconnection interface { // struct can be assigned any type in runtime
	Connect()
}

type DBConnection struct {
	Db IDBconnection // Compatible to accept any type in runtime
}

func (con DBConnection) DBConnect() { //Receiver Function For struct DBConnection
	con.Db.Connect()
}

// Lets implement First behaviour to connect to MySql
type MySqlConnection struct {
	ConnectionString string
}

func (con MySqlConnection) Connect() { //Receiver Function for struct MySqlConnection
	fmt.Println(("MySql " + con.ConnectionString))
}

// Second behaviour to connect to Postgress
type PostgressConnection struct {
	ConnectionString string
}

func (con PostgressConnection) Connect() { //Receiver Function for struct MySqlConnection
	fmt.Println("Postgress " + con.ConnectionString)
}

// Third behaviour to connect to MongoDB
type MongoDbConnection struct {
	ConnectionString string
}

func (con MongoDbConnection) Connect() { //Receiver Function for struct MySqlConnection
	fmt.Println("MongoDb " + con.ConnectionString)
}
func main() {
	MySqlConnection := MySqlConnection{ConnectionString: " Is Connected"}
	con := DBConnection{Db: MySqlConnection}
	con.DBConnect()

	PostgressConnection := PostgressConnection{ConnectionString: " Is Connected"}
	con = DBConnection{Db: PostgressConnection}
	con.DBConnect()

	MongoDbConnection := MongoDbConnection{ConnectionString: " Is Connected"}
	con = DBConnection{Db: MongoDbConnection}
	con.DBConnect()
}
