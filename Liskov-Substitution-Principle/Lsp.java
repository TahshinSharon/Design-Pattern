/*
 * If You have a parent class or interface, any subclass shoud be usable in its place without
 * Throwing unexpected exceptions
 * Removing Functionality
 * Changing behaviour in a surprising way
 */
interface Bird{
    void eat();
}
interface FlyableBird extends Bird{
    void fly();
}

class Sparrow implements FlyableBird{
    @Override
    public void eat(){
        System.out.println("Sparrow Is Eating:");
    }
    @Override
    public void fly(){
        System.out.println("Sparrow Is Flying:");
    }
}

class Eagle implements FlyableBird{
    @Override
    public void eat(){
        System.out.println("Eagle Is Eating:");
    }
    @Override
    public void fly(){
        System.out.println("Eagle Is Flying:");
    }
}
class Penguin implements Bird{
    @Override
    public void eat(){
        System.out.println("Penguine is Eating:");
    }
}
public class Lsp{
    public static void main(String[] args){
        FlyableBird sparrow = new Sparrow();
        FlyableBird eagle = new Eagle();
        Bird penguin = new Penguin();
        sparrow.eat();
        sparrow.fly();
        eagle.eat();
        eagle.fly();
        penguin.eat();
    }
}