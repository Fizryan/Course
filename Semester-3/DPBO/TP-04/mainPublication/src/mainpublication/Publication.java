package mainpublication;

import java.text.NumberFormat;
import java.util.Locale;

/**
 *
 * @author Hafizryandin HM
 */
public class Publication {
    private String title;
    private int price;
    private int page;

    public Publication(String title, int price, int page) {
        this.title = title;
        this.price = price;
        this.page = page;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public int getPrice() {
        return price;
    }

    public void setPrice(int price) {
        this.price = price;
    }

    public int getPage() {
        return page;
    }

    public void setPage(int page) {
        this.page = page;
    }

    public void printHeader() {
        System.out.println("===== Informasi Produk =====");
    }
    
    public void printBookInfo() {
        Locale myIndonesianLocale = new Locale("in", "ID");
        NumberFormat formatter = NumberFormat.getCurrencyInstance(myIndonesianLocale);
        System.out.println(getTitle() + " (Price List: " + formatter.format(getPrice()) + ")");
    }
}
