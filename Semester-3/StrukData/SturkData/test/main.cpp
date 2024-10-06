#include <iostream>
#include <cmath>
using namespace std;

class Cone {
private:
    double radius;
    double height;

public:
    Cone(double r, double h) : radius(r), height(h) {}
    double volume() {
        return (1.0 / 3.0) * M_PI * radius * radius * height;
    }
    double surfaceArea() {
        double slantHeight = sqrt(radius * radius + height * height);
        return M_PI * radius * (radius + slantHeight);
    }
};

int main() {
    Cone cone(5.0, 10.0);
    cout << "Volume kerucut: " << cone.volume() << endl;
    cout << "Luas permukaan kerucut: " << cone.surfaceArea() << endl;
    return 0;
}
