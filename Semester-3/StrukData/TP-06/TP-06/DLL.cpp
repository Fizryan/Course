#include "DLL.h"

bool isEmpty(List L) {
    return L.first == nullptr;
}

void createNewElmt(infotype X, address &P){
    P = new elmList;
    if (P != nullptr) {
        P -> info = X;
        P -> prev = nullptr;
        P -> next = nullptr;
    }
}

void insertFirst(List &L, address P) {
    if (isEmpty(L)){
        L.first = P;
        L.last = P;
    } else {
        P -> next = L.first;
        L.first -> prev = P;
        L.first = P;
    }
}

void insertAfter(List &L, address Prec, address P){
    if (Prec != nullptr && Prec -> next != nullptr){
        P -> next = Prec -> next;
        P -> prev = Prec;
        Prec -> next -> prev = P;
        Prec -> next = P;
    } else if (Prec == L.first){
        insertLast(L, P);
    }
}

void insertLast(List &L, address P){
    if (isEmpty(L)) {
        insertFirst(L, P);
    } else {
        L.last -> next = P;
        P -> prev = L.last;
        L.last = P;
    }
}

void deleteFirst(List &L, address &P){
    P = L.first;
    if (L.first -> next != nullptr){
        L.first = L.first -> next;
        L.first -> prev = nullptr;
    } else {
        L.first = nullptr;
        L.last = nullptr;
    }
    P -> next = nullptr;
}

void deleteAfter(List &L, address Prec, address &P){
    if (Prec != nullptr && Prec -> next != nullptr){
        P = Prec -> next;
        Prec -> next = P -> next;
        if (P -> next != nullptr){
            P -> next -> prev = Prec;
        } else {
            L.last = Prec;
        }
        P -> next = nullptr;
        P -> prev = nullptr;
    }
}

void deleteLast(List &L, address &P){
    P = L.last;
    if (L.first == L.last) {
        L.first = nullptr;
        L.last = nullptr;
    } else {
        L.last = L.last -> prev;
        L.last -> next = nullptr;
    }
}

void concat(List L1, List L2, List &L3){
    L3.first = nullptr;
    L3.last = nullptr;
    address P = L1.first;
    while (P != nullptr){
        address newElmt;
        createNewElmt(P -> info, newElmt);
        insertLast(L3, newElmt);
        P = P -> next;
    }

    P = L2.first;
    while (P != nullptr) {
        address newElmt;
        createNewElmt(P -> info, newElmt);
        insertLast(L3, newElmt);
        P = P -> next;
    }
}

address findLagu(string Judul, List &L){
    address P = L.first;
    while (P != nullptr && P -> info != Judul){
        P = P -> next;
    }
    return P;
}

void removeLagu(string Judul, List &L){
    address P = findLagu(Judul, L);
    if (P != nullptr) {
        if (P == L.first){
            deleteFirst(L, P);
        } else if (P == L.last){
            deleteLast(L, P);
        } else {
            deleteAfter(L, P -> prev, P);
        }
        delete P;
    }
}

void displayList(List L) {
    address Q = L.first;
    if (isEmpty(L)) {
        cout << "List is empty." << endl;
        return;
    }
    cout << "List contents: ";
    while (Q != nullptr) {
        cout << Q->info << " ";
        Q = Q->next;
    }
    cout << endl;
}


void displayMenu() {
    cout << "Menu:" << endl;
    cout << "1. Insert First" << endl;
    cout << "2. Insert Last" << endl;
    cout << "3. Delete First" << endl;
    cout << "4. Delete Last" << endl;
    cout << "5. Display List" << endl;
    cout << "6. Remove Lagu" << endl;
    cout << "7. Find Lagu" << endl;
    cout << "8. Concat Lists" << endl;
    cout << "9. Exit" << endl;
}


