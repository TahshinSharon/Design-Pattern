interface Transport{
    void deliver();
}
class Truck implements Transport{
    @Override
    public void deliver(){
        System.out.println("Delivering goods by truck on the road.");
    }
}
class Ship implements Transport{
    @Override
    public void deliver(){
        System.out.println("Delivering goods by ship across the sea");
    }
}
class Air implements Transport{
    @Override
    public void deliver(){
        System.out.println("Delivering goods by Air across the Air");
    }
}
abstract class Logistics{
    protected abstract Transport createTransport();
    public void planDelivery(){
        System.out.println("Planning delivery...");
        Transport transport= createTransport();
        transport.deliver();
    }
}
class RoadLogistics extends Logistics{
    @Override
    protected Transport createTransport(){
        return new Truck();
    }
}
class SeaLogistics extends Logistics{
    @Override
    protected Transport createTransport(){
        return new Ship();
    }
}
class AirLogistics extends Logistics{
    @Override
    protected Transport createTransport(){
        return new Air();
    }
}
public class FactoryPattern {
    public static void main(String[] args){
        Logistics logistics;
        logistics=new RoadLogistics();
        logistics.planDelivery();
        logistics=new SeaLogistics();
        logistics.planDelivery();
        logistics=new AirLogistics();
        logistics.planDelivery();
    }
}
