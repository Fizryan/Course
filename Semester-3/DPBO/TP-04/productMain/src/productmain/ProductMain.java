package productmain;

/**
 *
 * @author Hafizryandin HM
 */
public class ProductMain {
    public static void main(String[] args) {
        Product Laptop = new Product("Laptop", 15000000, 10);
        Product Mouse = new Product("Mouse", 150000, 50);
        Product Keyboard = new Product("Keyboard", 300000, 30);

        ProductManager manager = new ProductManager();

        System.out.println("--- Informasi Produk 1 ---");
        manager.tampilkanInformasiProduk(Laptop);
        System.out.println("");
        System.out.println("--- Informasi Produk 2 ---");
        manager.tampilkanInformasiProduk(Mouse);
        System.out.println("");
        System.out.println("--- Informasi Produk 3 ---");
        manager.tampilkanInformasiProduk(Keyboard);
    } 
}
