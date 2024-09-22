package bandingbil;
import java.util.Scanner;
/**
 *
 * @author Hafizryandin HM
 */
public class BandingBil {
    public static void main(String[] args) {
        System.out.print("Input: ");
        Scanner in = new Scanner(System.in);
        int[] bilangan = new int[11];
        
        for (int i = 0; i < 11; i++) {
            bilangan[i] = in.nextInt();
        }
        
        int bilanganKe11 = bilangan[10];
        
        System.out.println("Output: ");
        for (int i = 0; i < 10; i++) {
            if (bilangan[i] < bilanganKe11) {
                System.out.println(bilangan[i] + " Lebih Kecil Dari " + bilanganKe11);
            } else if (bilangan[i] > bilanganKe11) {
                System.out.println(bilangan[i] + " Lebih Besar Dari " + bilanganKe11);
            } else {
                System.out.println(bilangan[i] + " Sama Dengan " + bilanganKe11);
            }
        }
    }   
}
