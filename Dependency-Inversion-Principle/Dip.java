/*
 * Dependency Inversion Principle:
 * High Level modules should not depend on low-level modules,Both should depend on abstractions/interface(High Level module are core business logic or pollicies)
 * Abstractions should not depend on details.Details should depend on abstractions.
 * Don't let high-level code depend directly on low level implementations.
 * Instead both depend on interfaces(abstractions).
 */
//Interface
interface Database{
    void connect();
}
// Low-Level Modules
class PostgreSQLDatabase implements Database{
    public void connect(){
        System.out.println("Connected to PostgreSQL:");
    }
}
class MongoDbDatabase implements Database{
    public void connect(){
        System.out.println("Connected To MongoDB:");
    }
}
//High level Modules
class Application{ 
    private Database database;
    public Application(Database database){
        this.database = database;
    }
    public void start(){
        database.connect();
    }
}
 public class Dip{
    public static void main(String[] args){
        PostgreSQLDatabase postgre = new PostgreSQLDatabase();
        MongoDbDatabase mongodb = new MongoDbDatabase();

        Application app1 = new Application(postgre);
        app1.start();
        Application app2 = new Application(mongodb);
        app2.start();
    }
 }