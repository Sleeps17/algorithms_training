#include <iostream>
#include <vector>
#include <string>

std::vector<int> zFunction(const std::string& s) {
    int n = s.length();
    std::vector<int> z(n, 0);
    int left = 0, right = 0;
    for (int i = 1; i < n; i++) {
        if (i <= right) {
            z[i] = std::min(right - i + 1, z[i - left]);
        }
        while (i + z[i] < n && s[z[i]] == s[i + z[i]]) {
            z[i]++;
        }
        if (i + z[i] - 1 > right) {
            left = i;
            right = i + z[i] - 1;
        }
    }
    return z;
}

int main() {
    int n;
    std::cin >> n;
    std::string s;
    std::cin >> s;

    std::vector<int> z = zFunction(s);

    for (int i = 0; i < n; i++) {
        std::cout << z[i] << " ";
    }

    return 0;
}