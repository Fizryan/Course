#include <stdio.h>

int main() {
    int jumlahHari;
    printf("Masukkan jumlah hari: ");
    scanf("%d", &jumlahHari); // input jumlah hari oleh user

    for (int hari = 1; hari <= jumlahHari; hari++) { 
        if (hari % 2 == 0) { // hari akan di bagi 2 dan di ambil sisa hasil bagi nya
            printf("Hari ke-%d: Cabai memerlukan pupuk.\n", hari); // jika hasil nya 0 maka dia akan memprint ini
        }
        if (hari % 3 == 0) { // hari akan di bagi 3 dan di ambil sisa hasil bagi nya
            printf("Hari ke-%d: Tomat memerlukan pupuk.\n", hari); // jika hasil nya 0 maka dia akan memprint ini
        }
        if (hari % 4 == 0) { // hari akan di bagi 4 dan di ambil sisa hasil bagi nya
            printf("Hari ke-%d: Terong memerlukan pupuk.\n", hari); // jika hasil nya 0 maka dia akan memprint ini
        }
    }

    return 0; // program selesai
}
