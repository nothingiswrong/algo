#include <bits/stdc++.h> 
using namespace std;
#define ll long long
#define mod 1000000007

void solve() {
    int n;
    cin >> n;
    int maxim = -1;
    vector<int> t(n);
    for (int i = 0; i < n; i++) {
        cin >> t[i];
        maxim = max(maxim, t[i]);
    }
    vector<ll> h(maxim + 1);
    ll pr, npr;
    h[1] = 2;
    pr = 1;
    npr = 1;
    for (int i = 2; i <= maxim; i++) {
        h[i] = (5 * pr + 3 * npr) % mod;
        ll prOld = pr;
        pr = (4 * pr + npr) % mod;
        npr = (2 * npr + prOld) % mod;
    }
    for (int i: t) {
        cout << h[i] << '\n';
    }
}





int main() {
    solve();
    return 0;
}