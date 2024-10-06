#include <iostream>
#include "list.h"

using namespace std;

void createList(List &L) {
    first(L) = NULL;
}

address allocate(infotype x) {
    address p = new elmlist;
    info(p) = x;
    next(p) = NULL;
    return p;
}

void insertFirst(List &L, address P){
    next(P) = first(L);
    first(L) = P;
}

void printInfo(List L){
    address p = first(L);
    while (p != NULL) {
        cout << info(p) << ", ";
        p = next(p);
    }
    cout << endl;
}

void insertLast(List &L, address P) {
    if (first(L) == nullptr) {
        first(L) = P;
    } else {
        address last = first(L);
        while (next(last) != nullptr) {
            last = next(last);
        }
        next(last) = P;
    }
}

void insertAfter(List &L, address Prec, address P) {
    if (Prec != nullptr) {
        next(P) = next(Prec);
        next(Prec) = P;
    }
}

void deleteLast(List &L, address &P) {
    if (first(L) != nullptr) {
        address last = first(L);
        address prevLast = nullptr;

        if (next(last) == nullptr) {
            P = last;
            first(L) = nullptr;
        } else {
            while (next(last) != nullptr) {
                prevLast = last;
                last = next(last);
            }
            P = last;
            next(prevLast) = nullptr;
        }
    }
}

void deleteAfter(List &L, address Prec, address &P) {
    if(Prec != nullptr && next(Prec) != nullptr) {
        P = next(Prec);
        next(Prec) = next(P);
        next(P) = nullptr;
    }
}

address searchInfo(List L, int info) {
    address P = first(L);

    while (P != nullptr) {
        if (info(P) == info) {
            return P;
        }
        P = next(P);

    }
    return nullptr;
}
