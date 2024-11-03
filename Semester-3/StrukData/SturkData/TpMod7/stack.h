#ifndef STACK_H_INCLUDED
#define STACK_H_INCLUDED
typedef char infotype;

struct stack {
    infotype info[15];
    int Top;
};

void createStack_103022300158(stack &S);
bool isEmpty_103022300158(stack S);
bool isFull_103022300158(stack S);
void push_103022300158(stack &S, infotype x);
infotype pop_103022300158(stack &S);
void printInfo_103022300158(stack S);

#endif // STACK_H_INCLUDED
