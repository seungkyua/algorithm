package lang_java;

import java.util.Scanner;

public class standard_input {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int a, b;
        a = sc.nextInt();
        b = sc.nextInt();
        System.out.printf("%d\n", a + b);
        sc.close();
    }
}