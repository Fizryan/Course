package tigabilanganbulat;
import java.util.Scanner;
/**
 *
 * @author Hafizryandin HM
 */
public class TigaBilanganBulat {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        System.out.print("Input Tiga Bilangan Bulat: ");
        int a = in.nextInt();
        int b = in.nextInt();
        int c = in.nextInt();
        
        if (a > b) {
            int temp = a; a = b; b = temp;
        }
        if (a > c) {
            int temp = a; a = c; c = temp;
        }
        if (b > c) {
            int temp = b; b = c; c = temp;
        }
        System.out.println("Output: " + a + " " + b + " " + c);
    }
}
