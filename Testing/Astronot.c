#include <stdio.h>

int main() {
    int energy = 100; // default energy
    int hari = 0; // default hari
    
    while (energy > 0) { // melakukan perulangan selama energy > 0
        hari++; // bertambahnya hari selama misi
        // melakukan print menu
        printf("Hari %d:\n", hari); 
        printf("Tingkat energi Anda: %d\n", energy);
        printf("Pilih tindakan Anda:\n");
        printf("1. Melakukan pemeriksaan peralatan\n");
        printf("2. Melakukan spacewalk\n");
        printf("3. Makan\n");
        printf("4. Beristirahat\n");
        printf("Pilih: ");
        
        int pilihan;
        scanf("%d", &pilihan); // meminta input pilihan user
        
        switch (pilihan) { // mengambil tindak sesuai pilihan
            case 1:
                energy -= 20; // melakukan pemeriksaan peralatan
                energy -= 5; // mengurangi energi tiap harinya
                break;
            case 2:
                energy -= 35; // melakukan spacewalk
                energy -= 5; // mengurangi energi tiap harinya
                break;
            case 3:
                energy += 10; // makan
                energy -= 5; // mengurangi energi tiap harinya
                break;
            case 4:
                energy += 20; // beristirahat
                energy -= 5; // mengurangi energi tiap harinya
                break; 
            default:
                printf("Pilihan tidak valid.\n"); // default jika pilihan tidak sesuai
                continue;
        }
        
        if (energy <= 0) { // berhenti jika energy <= 0
            printf("Misi selesai pada hari %d.\n", hari);
            break;
        }
    }
    
    return 0; // program selesai
}
