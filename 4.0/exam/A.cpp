#include <iostream>
#include <vector>
#include <string>

std::vector<int> mirrorZFunction(const std::string& str) {
    int n = str.size();
    std::vector<int> z(n + 1); // increase the size of the vector by 1
    int l = 0, r = 0;
    for (int i = 1; i <= n; i++) { // change the loop condition
        if (i <= r) {
            z[i] = std::min(r - i + 1, z[i - l]);
        }
        while (i + z[i] <= n && str[z[i]] == str[i + z[i] - 1]) { // adjust the index
            z[i]++;
        }
        if (i + z[i] - 1 > r) {
            l = i;
            r = i + z[i] - 1;
        }
    }
    return z;
}

int main() {
    int n;
    std::cin >> n;
    std::string str;
    std::cin >> str;

    std::vector<int> z = mirrorZFunction(str);

    for (int i = 1; i <= n; i++) { // adjust the loop condition
        std::cout << z[i] << " ";
    }
    std::cout << std::endl;

    return 0;
}