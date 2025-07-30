// Singleton Pattern Guarantee that only one instance of this class will be available at any point in time

public class SingletonExample {

    // Step 1: Create a private static instance of the same class
    private static volatile SingletonExample instance;
    private String data;
    // Step 2: Make the constructor private to prevent external instantiation
    private SingletonExample(String data) {
        this.data =data;
        System.out.println(this.data);
    }

    // Step 3: Create a public static method to return the instance
    public static SingletonExample getInstance(String data) {
        if(instance == null){
        synchronized(SingletonExample.class){
             if (instance == null) {
            instance = new SingletonExample(data);
        }
    }
  }
  return instance;
}

    // Example method
    public void showMessage() {
        System.out.println("Hello from Singleton!");
    }

    // Main method to test the Singleton
    public static void main(String[] args) {
        // Try to create multiple instances
        SingletonExample obj1 = SingletonExample.getInstance("Singleton Instance Created");
        SingletonExample obj2 = SingletonExample.getInstance("2nd Singleton Instance Created");// this string should not be printted cause Instance is already created

        // Check if both objects are the same
        System.out.println("obj1 hash: " + obj1.hashCode());
        System.out.println("obj2 hash: " + obj2.hashCode());

        // Call a method
        obj1.showMessage();
    }
}