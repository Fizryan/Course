#include <iostream>
#include "DLL.h"

int main()
{
    List L1, L2, L3;
    L1.first = nullptr;
    L1.last = nullptr;
    L2.first = nullptr;
    L2.last = nullptr;
    L3.first = nullptr;
    L3.last = nullptr;
    address P;

    int choice;
    string laguJudul;

    //Anggap Sudah Ada Lagu Di L2 dan L3, buat Demonstrasi.
    createNewElmt("The Corrs", P);
    insertLast(L2, P);
    createNewElmt("Radio", P);
    insertLast(L2, P);
    createNewElmt("Maroon 5", P);
    insertLast(L3, P);
    createNewElmt("Sugar", P);
    insertLast(L3, P);

    do {
        displayMenu();
        cout << "Choose an option: ";
        cin >> choice;

        switch (choice) {
            case 1:
                cout << "Masukan Judul Lagu Untuk Di InsertFirst: ";
                cin.ignore();
                getline(cin, laguJudul);
                createNewElmt(laguJudul, P);
                insertFirst(L1, P);
                break;

            case 2:
                cout << "Masukan Judul Lagu Untuk Di InsertLast: ";
                cin.ignore();
                getline(cin, laguJudul);
                createNewElmt(laguJudul, P);
                insertLast(L1, P);
                break;

            case 3:
                deleteFirst(L1, P);
                cout << "Deleted first element: " << (P != nullptr ? P -> info : "None") << endl;
                delete P;
                break;

            case 4:
                deleteLast(L1, P);
                cout << "Deleted last element: " << (P != nullptr ? P -> info : "None") << endl;
                delete P;
                break;

            case 5:
                displayList(L1);
                break;

            case 6:
                cout << "Masukan Lagu Yang Ingin Di Hapus: ";
                cin.ignore();
                getline(cin, laguJudul);
                removeLagu(laguJudul, L1);
                break;

            case 7:
                cout << "Masukan Lagu Yang Ingin Di Cari: ";
                cin.ignore();
                getline(cin, laguJudul);
                P = findLagu(laguJudul, L1);
                if (P != nullptr) {
                    cout << "Song found: " << P -> info << endl;
                } else {
                    cout << "Song not found." << endl;
                }
                break;

            case 8:
                concat(L1, L2, L3);
                cout << "Concat L1 and L2 into L3." << endl;
                displayList(L3);
                break;

            case 9:
                cout << "Exiting the program." << endl;
                break;

            default:
                cout << "Invalid choice! Please try again." << endl;
        }
    } while (choice != 9);
    return 0;
}
