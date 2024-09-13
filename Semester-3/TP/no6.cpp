#include <iostream>

using namespace std;

int main() {
    int a, b, bilangan;
    cout << "Masukan Batasan Bawah: ";
    cin >> a;
    cout << "Masukan Batas Atas: ";
    cin >> b;
    for (bilangan = a; bilangan <= b; bilangan++) {
        cout << "Bilangan " << bilangan << endl;
    }
    return 0;
}