package buah;
import java.util.Scanner;
/**
 *
 * @author Hafizryandin HM
 */
public class Buah {
    public static void main(String[] args) {
        System.out.print("Banyak Buah: ");
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int apel = 0, jeruk = 0, mangga = 0;
        for (int i = 0; i < n; i++) {
            System.out.print("Input Jenis Buah: ");
            String buah = in.next();
            switch (buah) {
                case "apel":
                    apel++;
                    break;
                case "jeruk":
                    jeruk++;
                    break;
                case "mangga":
                    mangga++;
                    break;
                default:
                    break;
            }
        }
        System.out.println("Output: " + apel + " " + jeruk + " " + mangga);
    }
}
