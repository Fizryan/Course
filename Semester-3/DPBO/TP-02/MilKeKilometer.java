package milkekilometer;
import java.util.Scanner;
/**
 *
 * @author Hafizryandin HM
 */
public class MilKeKilometer {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        System.out.print("Enter Miles: ");
        double miles = in.nextDouble();
        double kilometers = miles * 1.6;
        System.out.println("Masukkan Jumlah Mil: " + miles + " mil" + " Setara Dengan " + kilometers + " kilometers");
    }
}
