#include <iostream>
#include "stack.h"
using namespace std;

void createStack_103022300158(stack &S){
    S.Top = 0;
}

bool isEmpty_103022300158(stack S){
    return S.Top == 0;
}

bool isFull_103022300158(stack S){
    return S.Top == 15;
}

void push_103022300158(stack &S, infotype x){
     if (!isFull_103022300158(S)) {
        S.Top++;
        S.info[S.Top] = x;
     }
}

infotype pop_103022300158(stack &S){
    if (!isEmpty_103022300158(S)) {
        infotype x = S.info[S.Top];
        S.Top--;
        return x;
    }
}

void printInfo_103022300158(stack S){
    for (int i = S.Top; i >= 0; i--) {
        cout << S.info[i] << " ";
    }
    cout << endl;
}
