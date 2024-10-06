package productmain;

/**
 *
 * @author Hafizryandin HM
 */
public class Product {
    private String namaProduk;
    private double harga;
    private int stok;
    
    public Product(String namaProduk, double harga, int stok) {
        this.namaProduk = namaProduk;
        this.harga = harga;
        this.stok = stok;
    }
    
    public String getNamaProduk() {
        return namaProduk;
    }
    
    public void setNamaProduk(String namaProduk) {
        this.namaProduk = namaProduk;
    }
    
    public double getHarga() {
        return harga;
    }
    
    public void setHarga(double harga) {
        this.harga = harga;
    }
    
    public int getStok() {
        return stok;
    }
    
    public void setStok(int stok) {
        this.stok = stok;
    }
    
    public void printProductInfo() {
        System.out.printf("Nama Produk: %s \n", namaProduk);
        System.out.printf("Harga: Rp. %.1f \n", harga);
        System.out.printf("Stok: %d Unit \n", stok);
    }
}
