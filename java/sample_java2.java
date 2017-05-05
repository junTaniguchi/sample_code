class User {
  String name;
  //コンストラクタ
  User(String name) {
    this.name = name;
  }

  void sayHi() {
    System.out.println("hi! " + this.name);
  }
}

class AdminUser extends User {
  //コンストラクタ
  AdminUser(String name) {
    super(name);
  }

  void sayHello() {
    System.out.println("hello! " + this.name);
  }

  // オーバーライド
  @Override
  void sayHi() {
    System.out.println("[admin] hi! " + this.name);
  }
}

public class MyApp {

  public static void main(String[] args) {
    User tom = new User("tom");
    System.out.println(tom.name);
    tom.sayHi();

    AdminUser bob = new AdminUser("bob");
    System.out.println(bob.name);
    bob.sayHi();
    bob.sayHello();
  }

}
