/*
 * Software entities (Classes ,modules,methods etc) should be open for
 * extensionn but closed for modification.
 */
interface Shape {
    double calculateArea();
}

class Circle implements Shape {
    private double radious;

    public Circle(double radious) {
        this.radious = radious;
    }

    @Override
    public double calculateArea() {
        return Math.PI * radious * radious;
    }
}

class Rectangle implements Shape {
    private double length;
    private double width;

    public Rectangle(double length, double width) {
        this.length = length;
        this.width = width;
    }

    @Override
    public double calculateArea() {
        return length * width;
    }
}

class Triangle implements Shape {
    private double base;
    private double height;

    public Triangle(double base, double height) {
        this.base = base;
        this.height = height;
    }

    @Override
    public double calculateArea() {
        return 0.5 * base * height;
    }
}

class Ocp {
    public static void main(String[] args) {
        Shape circle = new Circle(5.0);
        Shape rectangle = new Rectangle(4, 5);
        Shape triangle = new Triangle(3, 4);

        System.out.println("Area Of Circle: " + circle.calculateArea());
        System.out.println("Area Of Rectangle: " + rectangle.calculateArea());
        System.out.println("Area Of Triangle: " + triangle.calculateArea());
    }
}