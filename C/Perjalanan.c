#include <stdio.h>

int main() {
    int Start = (9 * 3600); // default awal jam dimulai pukul 09:00

    int speedAli = 5; // kecepatan Ali 5m/detik
    int speedBadu = 8; // kecepatan Badu 8m/detik

    int jarakAli = 18000; // jarak Ali sudah 18000 meter di karnakan Ali sudah lebih dulu berangkat 1 jam lebih dulu, jadi 5 * 3600 = 1800
    int jarakBadu = 0; // jarak Badu masih 0 dikarnakan Badu belum berangkat

    int bertemu = 0;

    while (jarakBadu <= jarakAli) { // ketika jarak Badu <= jarak ali, program akan selalu berulang / loop
        jarakAli += speedAli;  // jarak Ali akan di tambahkan kecepatan Ali
        jarakBadu += speedBadu; // jarak Badu akan di tambahkan kecepatan Badu

        Start++; // jam bertambah / berjalan

        if (jarakBadu > jarakAli) { // jika jarak Badu > jarak Ali maka program ini akan jalan
            bertemu = Start; // bertemu akan mengambil waktu/jam saat Badu sudah mendahului Ali
            break; // perulangan/loop akan berhenti
        }
    }
    int jamBertemu = (bertemu / 3600); // konversi detik ke jam
    int menitBertemu = (bertemu % 3600) / 60; // konversi ke menit
    int detikBertemu = (bertemu % 60); // konversi ke detik

    printf("Titik Ali %02dm titik Badu %02dm\n", jarakAli, jarakBadu); // output titik Ali dan titik Badu
    printf("Badu mendahului Ali pada pukul %02d:%02d:%02d\n", jamBertemu, menitBertemu, detikBertemu); // output hasil
    return 0; // program selesai
}
