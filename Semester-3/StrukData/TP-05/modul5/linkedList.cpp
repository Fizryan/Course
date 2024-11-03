#include <iostream>
#include <cstring>
#include "linkedList.h"
using namespace std;

void createList_103022300158(linkedList &L) {
    L.head = nullptr;
}

Node* createNewElement_103022300158(int value) {
    Node* newNode = new Node;
    newNode -> data = value;
    newNode -> next = nullptr;
    return newNode;
}

Node* createNewElement_103022300158(string value) {
    Node* newNode = new Node;
    newNode->name = value;
    newNode->next = nullptr;
    return newNode;
}

void insertLast_103022300158(linkedList &L, int value) {
    Node* newNode = createNewElement_103022300158(value);
    if (L.head == nullptr) {
        L.head = newNode;
    } else {
        Node* temp = L.head;
        while (temp -> next != nullptr) {
            temp = temp -> next;
        }
        temp->next = newNode;
    }
}

void insertLast_103022300158(linkedList &L, string value) {
    Node* newNode = createNewElement_103022300158(value);
    if (L.head == nullptr) {
        L.head = newNode;
    } else {
        Node* temp = L.head;
        while (temp -> next != nullptr) {
            temp = temp->next;
        }
        temp -> next = newNode;
    }
}

void showAllData_103022300158(linkedList L) {
    Node* temp = L.head;
    if (temp == nullptr) {
        cout << "Output: List kosong." << endl;
        return;
    }
    while (temp != nullptr) {
        cout << temp -> data << " -> ";
        temp = temp -> next;
    }
    while (temp != nullptr) {
        if (!temp -> name.empty()) {
            cout << "Nama: " << temp -> name << " -> ";
        } else if (temp -> character != '\0') {
            cout << "Karakter: " << temp -> character << " -> ";
        } else {
            cout << "Data: " << temp -> data << " -> ";
        }
        temp = temp -> next;
    }
    cout << "NULL" << endl;
}

Node* findMin_103022300158(linkedList L) {
    if (L.head == nullptr) return nullptr;

    Node* minNode = L.head;
    Node* temp = L.head;
    while (temp != nullptr) {
        if (temp -> data < minNode -> data) {
            minNode = temp;
        }
        temp = temp -> next;
    }
    return minNode;
}

void insertMiddle_103022300158(linkedList &L, int value) {
    Node* newNode = createNewElement_103022300158(value);
    Node* temp = L.head;
    int length = 0;

    while (temp != nullptr) {
        length++;
        temp = temp -> next;
    }

    if (length % 2 == 1) {
        int middle = length / 2;
        temp = L.head;
        for (int i = 0; i < middle - 1; i++) {
            temp = temp->next;
        }
        newNode -> next = temp -> next;
        temp -> next = newNode;
    } else {
        cout << "Error: Panjang List Harus Ganjil" << endl;
    }
}

void insertMiddle_103022300158(linkedList &L, string value) {
    Node* newNode = createNewElement_103022300158(value);
    Node* temp = L.head;
    int length = 0;

    while (temp != nullptr) {
        length++;
        temp = temp -> next;
    }

    if (length % 2 == 1) {
        int middle = length / 2;
        temp = L.head;
        for (int i = 0; i < middle - 1; i++) {
            temp = temp -> next;
        }
        newNode -> next = temp -> next;
        temp -> next = newNode;
    } else {
        cout << "Error: Panjang List Harus Ganjil" << endl;
    }
}

float consonantPercentage_103022300158(linkedList L) {
    if (L.head == nullptr) return 0.0f;

    int consonantCount = 0, total = 0;
    Node* temp = L.head;

    while (temp != nullptr) {
        char c = temp->character;
        if (isalpha(c) && !strchr("aeiou", c)) {
            consonantCount++;
        }
        total++;
        temp = temp->next;
    }

    if (total == 0) return 0.0f;
    return (float)consonantCount / total * 100;
}

void showFirstK_103022300158(linkedList L, int k) {
    if (L.head == nullptr) {
        cout << "List kosong." << endl;
        return;
    }

    Node* temp = L.head;
    int count = 1;
    while (temp != nullptr && count < k) {
        temp = temp-> next;
        count++;
    }

    if (temp != nullptr) {
        cout << "Karakter pada posisi " << k << ": " << temp -> character << endl;
    } else {
        cout << "Posisi melebihi jumlah elemen list." << endl;
    }
}

Node* shortestName_103022300158(linkedList L) {
    if (L.head == nullptr) return nullptr;

    Node* shortestNode = L.head;
    Node* temp = L.head;

    while (temp != nullptr) {
        if (temp->name.length() <= shortestNode->name.length()) {
            shortestNode = temp;
        }
        temp = temp->next;
    }
    return shortestNode;
}

void showFirstKNameC_103022300158(linkedList L, int k, char c) {
    if (L.head == nullptr) {
        cout << "List kosong." << endl;
        return;
    }

    Node* temp = L.head;
    int count = 0;
    while (temp != nullptr && count < k) {
        if (temp->name[0] == c) {
            cout << "Nama: " << temp->name << " ";
            count++;
        }
        temp = temp->next;
    }
    cout << endl;
}

int selectMenu_103022300158() {
    int menu;
    cout << "========MENU========" << endl;
    cout << "1. Menambah N data baru" << endl;
    cout << "2. Menampilkan semua data" << endl;
    cout << "3. Temukan nilai terkecil dalam list" << endl;
    cout << "4. Sisipkan data di tengah list" << endl;
    cout << "5. Persentase huruf konsonan dalam list" << endl;
    cout << "6. Tampilkan karakter ke-k" << endl;
    cout << "7. Pengunjung dengan nama terpendek" << endl;
    cout << "8. Tampilkan nama pengunjung" << endl;
    cout << "0. Exit" << endl;
    cout << "Masukkan pilihan menu: ";
    cin >> menu;
    return menu;
}

void inputAngka_103022300158(int &Num, int x) {
    Num = x;
}

void inputName_103022300158(string &Name, string value) {
    Name = value;
}
