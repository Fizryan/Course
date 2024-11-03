#ifndef LINKEDLIST_H_INCLUDED
#define LINKEDLIST_H_INCLUDED
#endif
#include <string>
using namespace std;

struct Node {
    int data;
    char character;
    string name;
    Node* next;
};

struct linkedList {
    Node* head;
};

void createList_103022300158(linkedList &L);
Node* createNewElement_103022300158(int value);
Node* createNewElement_103022300158(string value);
void insertLast_103022300158(linkedList &L, int value);
void insertLast_103022300158(linkedList &L, string value);
void showAllData_103022300158(linkedList L);

Node* findMin_103022300158(linkedList L);
void insertMiddle_103022300158(linkedList &L, int value);
void insertMiddle_103022300158(linkedList &L, string value);

float consonantPercentage_103022300158(linkedList L);
void showFirstK_103022300158(linkedList L, int k);

Node* shortestName_103022300158(linkedList L);
void showFirstKNameC_103022300158(linkedList L, int k, char c);

int selectMenu_103022300158();
void inputAngka_103022300158(int &Num, int x);
void inputName_103022300158(string &Name, string value);
