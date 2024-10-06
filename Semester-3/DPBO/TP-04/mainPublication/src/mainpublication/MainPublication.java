package mainpublication;

import java.text.NumberFormat;
import java.util.Locale;

/**
 *
 * @author Hafizryandin HM
 */
public class MainPublication {
    public static void main(String[] args) {
        Publication threeKingdom = new Publication("Three Kingdom", 5000, 500);
        Publication aPie = new Publication("a Pie", 8400, 200);

        threeKingdom.printHeader();
        threeKingdom.printBookInfo();
        aPie.printBookInfo();

        System.out.println("============================");

        int pembelianThreeKingdom = 67;
        int pembelianPie = 82;

        Locale myIndonesianLocale = new Locale("in", "ID");
        NumberFormat formatter = NumberFormat.getCurrencyInstance(myIndonesianLocale);

        System.out.println(threeKingdom.getTitle() + " - Pembelian Jilid: " + pembelianThreeKingdom);
        System.out.println(aPie.getTitle() + " - Pembelian Jilid: " + pembelianPie);
        System.out.println("==============================");

        int totalThreeKingdom = pembelianThreeKingdom * threeKingdom.getPrice();
        int totalPie = pembelianPie * aPie.getPrice();
        int totalPembelian = totalThreeKingdom + totalPie;

        System.out.println(threeKingdom.getTitle() + " " + pembelianThreeKingdom + " Jilid (" + formatter.format(totalThreeKingdom) + ")");
        System.out.println(aPie.getTitle() + " " + pembelianPie + " Jilid (" + formatter.format(totalPie) + ")");
        System.out.println("==============================");

        System.out.println("Jumlah Total Pembelian: " + formatter.format(totalPembelian));
    }  
}
