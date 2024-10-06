#include <iostream>
#include "list.h"
using namespace std;

int main()
{
    List L;
    List T;
    createList(L);
    createList(T);

    address P1 = allocate(1);
    address P2 = allocate(5);
    address P3 = allocate(8);

    insertFirst(L, P1);
    insertFirst(L, P2);
    insertFirst(L, P3);

    cout << "INI YANG BAWAAN" << endl;
    printInfo(L);
    cout << "" << endl;

    int NIM[10];
    cout << "Masukkan NIM per digit: " << endl;
    for (int i = 0; i < 10; i++){
        cout << "Digit " << i + 1 << " : ";
        cin >> NIM[i];
    }

    for (int i = 0; i < 10; i++){
        address P = allocate(NIM[i]);
        insertLast(T, P);
    }

    cout << "Isi List: ";
    printInfo(T);
    return 0;
}
