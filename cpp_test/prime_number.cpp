#include <iostream>
using namespace std;
int main() {
    int n;
    cout << "请输入你要列出的质数范围（1～输入的数）：" << endl;
    cin >> n;
    for (int i = 2; i <= n; i++) {
        int k = 0;
        for (int j = 2; j <= i / 2; j++) {
            if (i % j == 0) {
                k = 1;
                break;
            }
        }
        if (k == 0)
            cout << i << endl;
    }
    return 0;
}
