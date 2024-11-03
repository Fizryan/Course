#include <iostream>
#include "stack.h"
using namespace std;

int main()
{
    stack S;
    createStack_103022300158(S);
    int digitAkhirNIM = 8 % 4;
    if (digitAkhirNIM == 0) {
        cout << "Output Nim Mod 4 Sisa 0" << endl;
        char test1[] = "AYAJBALFI";
        for (int i = 0; i < 9; i++) {
            push_103022300158(S, test1[i]);
        }
        cout << "Isi Stack Awal: ";
        printInfo_103022300158(S);
        for (int i = 0; i < 5; i++) {
            pop_103022300158(S);
        }
        cout << "Isi Stack Setelah Pop: ";
        printInfo_103022300158(S);
        cout << endl;
    } else if (digitAkhirNIM == 1){
        cout << "Output Nim Mod 4 Sisa 1" << endl;
        char test2[] = "GNUDNABOLAH";
        for (int i = 0; i < 11; i++) {
            push_103022300158(S, test2[i]);
        }
        cout << "Isi Stack Awal: ";
        printInfo_103022300158(S);
        for (int i = 0; i < 4; i++) {
            pop_103022300158(S);
        }
        cout << "Isi Stack Setelah Pop: ";
        printInfo_103022300158(S);
        cout << endl;
    } else if (digitAkhirNIM == 2){
        cout << "Output Nim Mod 4 Sisa 2" << endl;
        char test3[] = "IRIDAYACREP";
        for (int i = 0; i < 11; i++) {
            push_103022300158(S, test3[i]);
        }
        cout << "Isi Stack Awal: ";
        printInfo_103022300158(S);
        for (int i = 0; i < 7; i++) {
            pop_103022300158(S);
        }
        cout << "Isi Stack Setelah Pop: ";
        printInfo_103022300158(S);
        cout << endl;
    } else if (digitAkhirNIM == 3){
        cout << "Output Nim Mod 4 Sisa 3" << endl;
        char test4[] = "ATADRUTKURTS";
        for (int i = 0; i < 12; i++) {
            push_103022300158(S, test4[i]);
        }
        cout << "Isi Stack Awal: ";
        printInfo_103022300158(S);
        for (int i = 0; i < 8; i++) {
            pop_103022300158(S);
        }
        cout << "Isi Stack Setelah Pop: ";
        printInfo_103022300158(S);
        cout << endl;
    }
    return 0;
}
