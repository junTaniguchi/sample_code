// interface

interface Printable {
  // 定数
  double VERSION = 1.2;
  // 抽象メソッド
  void print();
  // defaultメソッド
  public default void getInfo() {
    System.out.println("I/F ver. " + Printable.VERSION);
  }
  // staticメソッド
}

class User implements Printable {
  @Override
  public void print() {
    System.out.println("Now printing user profile...");
  }
}

public class MyApp {

  public static void main(String[] args) {
    User tom = new User();
    tom.print();
    tom.getInfo();
  }

}
