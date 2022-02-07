#include <iostream>
using namespace std;
int main() {
    int a[10] = {1, 23, 45, 0, 4325, 66, 2, 45, 4, 65};
    int min = a[0];
    for (int i = 0; i < 10; ++i) {
        if (a[i] <= min) {
            min = a[i];
        }
    }
    cout << "最小的数字:" << endl;
    cout << min << endl;
}
