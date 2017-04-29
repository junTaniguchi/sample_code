public class MyApp {

  public static void main(String[] args) {
    // if文
    int score = 95;
    // if (score > 80) {
    //   System.out.println("Great!");
    // } else if (score > 60) {
    //   System.out.println("Good!");
    // } else {
    //   System.out.println("so so ... !");
    // }
    String msg = score > 80 ? "Great!" : "so so ... !";
    System.out.println(msg);

    // switch文
    String signal = "green";
    switch (signal) {
      case "red":
        System.out.println("stop!");
        break;
      case "blue":
      case "green":
        System.out.println("go!");
        break;
      case "yellow":
        System.out.println("caution!");
        break;
      default:
        System.out.println("wrong signal!");
        break;
    }

    // for文
    for (int i = 0; i < 10; i++) {
      if (i == 5) {
        // break;
        continue;
      }
      System.out.println(i);
    }

    // 配列
    // int[] sales;
    // sales = new int[3];
    // int[] sales;
    // sales = new int[] {700, 400, 500};
    int[] sales = {700, 400, 500};
    // sales[0] = 700;
    // sales[1] = 400;
    // sales[2] = 500;
    System.out.println(sales[1]); // 400
    sales[1] = 500;
    System.out.println(sales[1]); // 500

    // 配列
    int[] sales = {700, 400, 500};
    for (int sale : sales) {
      System.out.println(sale);
    }

  }

}
