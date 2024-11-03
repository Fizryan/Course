#include <iostream>
#include "linkedList.h"
using namespace std;

int main() {
    linkedList L;
    createList_103022300158(L);

    int menu = 0;

    do {
        menu = selectMenu_103022300158();

        switch (menu) {
            case 1: {
                int Num = 0, x, N, nimDigit9;
                string Name;
                char c;
                string nim;
                cout << "Jumlah data yang akan ditambahkan: ";
                cin >> N;
                for (int i = 0; i < N; i++) {
                    cout << "Masukkan NIM ke-" << i + 1 << ": ";
                    cin >> nim;
                    nimDigit9 = nim[8] - '0';
                    int modResult = nimDigit9 % 3;
                    if (modResult == 0) {
                        Num = stoi(nim);
                        inputAngka_103022300158(Num, x);
                        insertLast_103022300158(L, Num);
                        cout << nimDigit9 << " % 3 = " << modResult << " Data Disimpan adalah integer" << endl;
                    } else if (modResult == 1) {
                        c = nim[8];
                        insertLast_103022300158(L, c);
                        cout << nimDigit9 << " % 3 = " << modResult << " Data Disimpan adalah char" << endl;
                    } else {
                        insertLast_103022300158(L, nim);
                        cout << nimDigit9 << " % 3 = " << modResult << " Data Disimpan adalah string" << endl;
                    }
                }
                break;
            }

            case 2: {
                showAllData_103022300158(L);
                break;
            }

            case 3: {
                Node* minNode = findMin_103022300158(L);
                if (minNode != nullptr) {
                    cout << "Output: Nilai terkecil: " << minNode -> data << endl;
                } else {
                    cout << "Output: List kosong." << endl;
                }
                break;
            }

            case 4: {
                int Num;
                string Name;
                char c;
                string nim;
                cout << "Masukkan Data Baru Yang Akan Disisipkan" << endl;
                cout << "Masukkan NIM Anda: ";
                cin >> nim;
                int nimDigit9 = nim[8] - '0';
                int modResult = nimDigit9 % 3;
                if (modResult == 0) {
                    Num = stoi(nim);
                    insertMiddle_103022300158(L, Num);
                    cout << nimDigit9 << " % 3 = " << modResult << " Data Disisipkan adalah integer" << endl;
                } else if (modResult == 1) {
                    c = nim[8];
                    insertMiddle_103022300158(L, c);
                    cout << nimDigit9 << " % 3 = " << modResult << " Data Disisipkan adalah char" << endl;
                } else {
                    insertMiddle_103022300158(L, nim);
                    cout << nimDigit9 << " % 3 = " << modResult << " Data Disisipkan adalah string" << endl;
                }
                break;
            }

            case 5: {
                float percentage = consonantPercentage_103022300158(L);
                cout << "Output: Persentase huruf konsonan: " << percentage << "%" << endl;
                break;
            }

            case 6: {
                int k;
                cout << "Masukkan posisi ke-k yang ingin ditampilkan: ";
                cin >> k;
                showFirstK_103022300158(L, k);
                break;
            }

            case 7: {
                Node* shortest = shortestName_103022300158(L);
                if (shortest != nullptr) {
                    cout << "Output: Pengunjung dengan nama terpendek: " << shortest -> name << endl;
                } else {
                    cout << "Output: List kosong." << endl;
                }
                break;
            }

            case 8: {
                int k;
                char c;
                cout << "Masukkan jumlah pengunjung yang ingin ditampilkan: ";
                cin >> k;
                cout << "Masukkan huruf awal nama: ";
                cin >> c;
                showFirstKNameC_103022300158(L, k, c);
                break;
            }

            case 0: {
                cout << "Anda telah keluar dari program" << endl;
                break;
            }

            default: {
                cout << "Menu tidak valid!" << endl;
                break;
            }
        }
        cout << "" << endl;
    } while (menu != 0);
    return 0;
}
