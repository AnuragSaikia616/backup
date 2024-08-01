---
link: https://leetcode.com/discuss/interview-question/3828150/OOPS-Cheatsheet-for-Interviews-or-30-questions
---
#### Theory Questions
1. What is Object-Oriented Programming (OOP)?
Answer: Object-Oriented Programming is a programming paradigm that organizes code into objects, which are instances of classes. It emphasizes concepts like inheritance, encapsulation, polymorphism, and abstraction to improve code reusability and maintainability.

2. What is a class?
Answer: A class is a blueprint or template that defines the structure and behavior of objects. It serves as a user-defined data type and encapsulates data (attributes) and methods (functions) that operate on that data.

3. What is an object?
Answer: An object is an instance of a class. It represents a real-world entity and can hold its own state (attributes) and behavior (methods) defined by the class.

4. What is inheritance?
Answer: Inheritance is a mechanism where one class (subclass or derived class) acquires properties and behaviors from another class (superclass or base class). It promotes code reusability and establishes an "is-a" relationship between classes.

5. Explain the different types of inheritance.
Answer: The main types of inheritance are:
Single Inheritance: A subclass inherits from a single superclass.
Multiple Inheritance: A subclass inherits from multiple superclasses (supported by some languages like C++).
Multilevel Inheritance: A class is derived from another class, which, in turn, is derived from a base class.
Hierarchical Inheritance: Multiple subclasses inherit from a single superclass.

6. What is encapsulation?
Answer: Encapsulation is the bundling of data (attributes) and methods (functions) that operate on the data within a single unit (i.e., the class). It restricts access to the data by providing public interfaces (public methods) to interact with the object's state. Encapsulation helps in data hiding and protecting the internal state of an object.

7. What is polymorphism?
Answer: Polymorphism allows objects of different classes to be treated as objects of a common superclass. It enables a single interface to represent multiple types of objects. Polymorphism is achieved through method overloading (compile-time polymorphism) and method overriding (runtime polymorphism).

8. What is abstraction?
Answer: Abstraction is the process of hiding the implementation details of an object and exposing only the relevant features or behavior. It allows developers to work with high-level concepts and ignore low-level implementation complexities.

9. What is the difference between method overloading and method overriding?
Answer: Method overloading occurs when a class has multiple methods with the same name but different parameter lists. The appropriate method is determined at compile time based on the method signature. Method overriding occurs when a subclass provides a specific implementation for a method that is already defined in its superclass. The decision on which method to call is made at runtime based on the actual object type.

10. What are abstract classes?
Answer: An abstract class is a class that cannot be instantiated directly and may contain one or more abstract methods. Abstract methods are declared without implementation and must be implemented by subclasses.

11. What is an interface?
Answer: An interface is a blueprint that defines a set of methods that a class must implement. It provides a way to achieve multiple inheritance in languages that do not support it directly.

12. What is the "this" keyword?
Answer: The "this" keyword refers to the current object instance within a class. It is used to access instance variables or call instance methods of the class.

13. What is the "super" keyword?
Answer: The "super" keyword is used to call the superclass's constructor or refer to the superclass's methods or variables from a subclass.

14. What is method hiding?
Answer: Method hiding occurs when a subclass defines a static method with the same name and signature as a static method in the superclass. The subclass's method hides the superclass's method, and the method invoked depends on the reference type rather than the object type.

15. What are access modifiers, and what are their purposes?
Answer: Access modifiers define the visibility and accessibility of class members (attributes, methods, constructors). The main access modifiers are public, private, protected, and package-private/default. They control how members can be accessed from other classes and packages.

16. Explain the "final" keyword.
Answer: In OOP, the "final" keyword can be applied to a class, method, or variable. A final class cannot be subclassed, a final method cannot be overridden, and a final variable cannot be reassigned once initialized.

17. What is a constructor?
Answer: A constructor is a special method that is automatically called when an object of a class is created. It is used to initialize the object's state and perform setup tasks.

18. What is a destructor?
Answer: A destructor is a special method that is called when an object is destroyed or goes out of scope. It is used to release resources and perform cleanup operations.

19. Explain the concept of static members
Answer: Static members (attributes or methods) belong to the class rather than individual objects. They are shared among all instances of the class and can be accessed using the class name.

20. What is the "instanceof" operator used for?
Answer: The "instanceof" operato is used to test whether an object is an instance of a particular class or implements a specific interface. It checks the object's type at runtime and returns a boolean value indicating whether the object is an instance of the specified class or interface.

21. What is a constructor chaining?
Answer: Constructor chaining is the process of calling one constructor from another within the same class or between base and derived classes. It allows constructors to reuse code and perform common initialization tasks.

22. What is the difference between composition and inheritance?
Answer: Composition represents a "has-a" relationship, where a class contains objects of other classes as its members. Inheritance represents an "is-a" relationship, where a subclass inherits properties and behaviors from a superclass.

23. What is method visibility, and how is it controlled?
Answer: Method visibility refers to the accessibility of methods from other classes. It is controlled by access modifiers like "public," "private," "protected," and "package-private/default."

24. Explain the concept of dynamic method dispatch.
Answer: Dynamic method dispatch is the mechanism that enables a method call on an object to be resolved at runtime based on the actual type of the object, rather than the reference type.

25. What is the difference between shallow copy and deep copy?
Answer: Shallow copy creates a new object but copies only the references of the original object's members. Deep copy, on the other hand, creates a new object and copies all the content, including recursively copying nested objects, if any.

26. What are virtual functions (or virtual methods)?
Answer: Virtual functions allow late binding or runtime polymorphism, enabling a subclass to provide its implementation for a method declared in a base class.

27. What is the Liskov Substitution Principle (LSP)?
Answer: The Liskov Substitution Principle states that objects of a superclass should be replaceable with objects of its subclasses without affecting the correctness of the program. Subclasses must maintain the behavior expected by the superclass.

Bonus Coding Questions to implement what we've studied:

Note: I code in C++ only, @aniket123456789 has contributed all the codes in Java

#### Implementation Questions

1. Create a class "BankAccount" with attributes account number and balance. Implement methods to deposit and withdraw funds, and a display method to show the account details.

C++
```c++
#include <iostream>
#include <string>

class BankAccount {
private:
    std::string account_number;
    double balance;

public:
    BankAccount(std::string acc_num, double initial_balance) : account_number(acc_num), balance(initial_balance) {}

    void deposit(double amount) {
        balance += amount;
    }

    void withdraw(double amount) {
        if (balance >= amount) {
            balance -= amount;
        } else {
            std::cout << "Insufficient balance!" << std::endl;
        }
    }

    void display() {
        std::cout << "Account Number: " << account_number << std::endl;
        std::cout << "Balance: " << balance << std::endl;
    }
};

int main() {
    BankAccount account("123456789", 1000);
    account.deposit(500);
    account.withdraw(200);
    account.display();

    return 0;
}
```
		


2. Create a base class "Shape" with methods to calculate the area and perimeter (pure virtual). Implement derived classes "Rectangle" and "Circle" that inherit from "Shape" and provide their own area and perimeter calculations.
C++:
```c++
#include <iostream>

class Shape {
public:
    virtual double area() const = 0;
    virtual double perimeter() const = 0;
};

class Rectangle : public Shape {
private:
    double length;
    double width;

public:
    Rectangle(double l, double w) : length(l), width(w) {}

    double area() const override {
        return length * width;
    }

    double perimeter() const override {
        return 2 * (length + width);
    }
};

class Circle : public Shape {
private:
    double radius;

public:
    Circle(double r) : radius(r) {}

    double area() const override {
        return 3.141592653589793 * radius * radius;
    }

    double perimeter() const override {
        return 2 * 3.141592653589793 * radius;
    }
};

int main() {
    Shape* shape1 = new Rectangle(5, 3);
    Shape* shape2 = new Circle(4);

    std::cout << "Rectangle Area: " << shape1->area() << std::endl;
    std::cout << "Rectangle Perimeter: " << shape1->perimeter() << std::endl;

    std::cout << "Circle Area: " << shape2->area() << std::endl;
    std::cout << "Circle Perimeter: " << shape2->perimeter() << std::endl;

    delete shape1;
    delete shape2;

    return 0;
}


```


3. Create a class "Person" with a static member variable "count" that keeps track of the number of instances created.
C++:
```c++
#include <iostream>
#include <string>

class Person {
private:
    std::string name;
    static int count;

public:
    Person(std::string n) : name(n) {
        count++;
    }

    static int getCount() {
        return count;
    }

    std::string getName() const {
        return name;
    }
};

int Person::count = 0;

int main() {
    Person person1("Alice");
    Person person2("Bob");

    std::cout << "Total Persons: " << Person::getCount() << std::endl;
    std::cout << person1.getName() << std::endl;
    std::cout << person2.getName() << std::endl;

    return 0;
}


```

4. Create a class "Employee" with attributes name and salary. Implement overloaded operators + and - to combine and compare employees based on their salaries.
C++:
```c++
#include <iostream>
#include <string>

class Employee {
private:
    std::string name;
    double salary;

public:
    Employee(std::string n, double s) : name(n), salary(s) {}

    double getSalary() const {
        return salary;
    }

    bool operator<(const Employee& other) const {
        return salary < other.salary;
    }

    bool operator>(const Employee& other) const {
        return salary > other.salary;
    }

    Employee operator+(const Employee& other) const {
        return Employee("Combined", salary + other.salary);
    }

    Employee operator-(const Employee& other) const {
        return Employee("Difference", salary - other.salary);
    }
};

int main() {
    Employee emp1("Alice", 5000);
    Employee emp2("Bob", 6000);

    if (emp1 < emp2) {
        std::cout << "Bob has a higher salary than Alice." << std::endl;
    } else {
        std::cout << "Alice has a higher salary than Bob." << std::endl;
    }

    Employee combined = emp1 + emp2;
    std::cout << "Combined Salary: " << combined.getSalary() << std::endl;

    Employee difference = emp1 - emp2;
    std::cout << "Salary Difference: " << difference.getSalary() << std::endl;

    return 0;
}

```



5. Create a class "Time" with attributes hours and minutes. Implement the << operator to display time in the format "hh:mm".
C++:
```c++
#include <iostream>

class Time {
private:
    int hours;
    int minutes;

public:
    Time(int h, int m) : hours(h), minutes(m) {}

    friend std::ostream& operator<<(std::ostream& os, const Time& time) {
        os << time.hours << ":" << (time.minutes < 10 ? "0" : "") << time.minutes;
        return os;
    }
};

int main() {
    Time time(14, 30);
    std::cout << "Time: " << time << std::endl;

    return 0;
}


```

