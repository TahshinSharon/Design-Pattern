/*
 *Interface Segregation Principle:
 * Definition (SOLID - “I”) Clients should not be forced to depend on interfaces they do not use.
 * Don't create fat interfaces with too many methods. Instead ,split them into smaller, more specific ones so that classes only implement what they actually need.
 */

interface Workable{
    void work();
}
interface Eatable{
    void eat();
}
class HumanWorker implements Workable,Eatable{
    @Override
    public void work(){
        System.out.println("Human is working.");
    }
    @Override
    public void eat(){
        System.out.println("Human is Eating Lunch");
    }
}
class RobotWorker implements Workable{
    @Override
    public void work(){
        System.out.println("Robot is Working");
    }
}

public class Isp{
    public static void main(String[] args){
        Workable human = new HumanWorker();
        Workable robot = new RobotWorker();

        Eatable eatingHuman = new HumanWorker();
        
        human.work();
        robot.work();
        eatingHuman.eat();

        //
    }
}

/*
 * Why This Respects ISP:
	•	Each class only implements interfaces with methods it needs.
	•	RobotWorker doesn’t implement Eatable → avoids being forced to override irrelevant methods.
	•	Code is more maintainable, modular, and extensible.
 */