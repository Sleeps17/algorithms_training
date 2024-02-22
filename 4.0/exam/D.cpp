#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

pair<u_int64_t, vector<u_int64_t>> findBorderBricks(u_int64_t N, vector<u_int64_t>& bricks) {
    sort(bricks.begin(), bricks.end(), greater<u_int64_t>());

    vector<u_int64_t> dp(N+1, -2); // -2 indicates uninitialized state
    vector<u_int64_t> prev(N+1, -1);
    dp[0] = 0;

    for (u_int64_t i = 1; i <= N; i++) {
        for (u_int64_t brick : bricks) {
            if (i - brick >= 0 && dp[i-brick] != -2 && dp[i-brick] + 1 > dp[i]) {
                dp[i] = dp[i-brick] + 1;
                prev[i] = brick;
            }
        }
    }

    if (dp[N] == -2 || dp[N] == 0) {
        return make_pair(0, vector<u_int64_t>()); // Return 0 if the sum of brick lengths is greater than N
    }

    u_int64_t cur = N;
    vector<u_int64_t> result;
    while (cur != 0) {
        result.push_back(prev[cur]);
        cur -= prev[cur];
    }

    return make_pair(result.size(), result);
}

int main() {
    u_int64_t N, M;
    cin >> N >> M;
    vector<u_int64_t> bricks(M);
    for (u_int64_t i = 0; i < M; i++) {
        cin >> bricks[i];
    }

    pair<u_int64_t, vector<u_int64_t>> result = findBorderBricks(N, bricks);

    if (result.first == 0) {
        cout << 0 << endl;
    } else {
        cout << result.first << endl;
        for (u_int64_t brick : result.second) {
            cout << brick << " ";
        }
        cout << endl;
    }

    return 0;
}