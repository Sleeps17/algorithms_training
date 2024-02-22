#include <iostream>
#include <string>
#include <vector>
using namespace std;

const int p = 31;
const int m = 1e9 + 9;

bool is_palindrome(const string& s, int start, int end) {
    int len = end - start + 1;
    int hash_forward = 0, hash_backward = 0;
    int p_pow = 1;

    for (int i = 0; i < len; i++) {
        hash_forward = (hash_forward + (s[start + i] - 'a' + 1) * p_pow) % m;
        hash_backward = (hash_backward * p + s[start + i] - 'a' + 1) % m;
        p_pow = (p_pow * p) % m;
    }

    return hash_forward == hash_backward;
}


int count_palindromic_substrings(const string& s) {
    int count = 0;
    int n = s.length();
    for (int i = 0; i < n; i++) {
        for (int j = i; j < n; j++) {
            if (is_palindrome(s, i, j)) {
                count++;
            }
        }
    }
    return count;
}

int main() {
    string s;
    cin >> s;
    int result = count_palindromic_substrings(s);
    cout << result << endl;
    return 0;
}