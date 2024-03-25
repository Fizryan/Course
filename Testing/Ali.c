#include <stdio.h>

int main() {
    int wktAliStart = 8 * 3600;  // Waktu Jam Ali Start convert ke detik (28.800)
    int wktBaduStart = 9 * 3600; // Waktu Jam Badu Start convert ke detik (32.400)

    int kcptAli = 5; // Kecepatan Ali 5m/detik
    int kcptBadu = 8; // Kecepatan Badu 8m/detik

    int jarakAli = 0; // Jarak Ali mulai dari titik 0
    int jarakBadu = 0; // Jarak Badu mulai dari titik 0

    int waktuBertemu;

    while (jarakBadu <= jarakAli) { // Kalau Jarak Badu <= jarak Ali Loopnya jalan
        // Nambahin jarak Ali Sama Badu
        jarakAli += kcptAli; // dari 0 ke 5
        jarakBadu += kcptBadu; // dari 0 ke 8

        // Nambahin Waktu
        wktAliStart++;
        wktBaduStart++;

        // kalau Badu mendahului Ali keluar dari loopnya
        if (jarakBadu > jarakAli) {
            waktuBertemu = wktAliStart;
            break;
        }
    }

    int jamBertemu = waktuBertemu / 3600;
    int menitBertemu = (waktuBertemu % 3600) / 60;
    int detikBertemu = waktuBertemu % 60;

    printf("Ali dan Badu bertemu pada pukul %02d:%02d:%02d\n", jamBertemu, menitBertemu, detikBertemu);

    return 0;
}
