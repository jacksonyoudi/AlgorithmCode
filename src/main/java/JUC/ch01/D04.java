package JUC.ch01;

public class D04 {
    public static void main(String[] args) {
        long l = Runtime.getRuntime().maxMemory();


        long l1 = Runtime.getRuntime().totalMemory();

        System.out.println(l);
        System.out.println(l1);
    }
}