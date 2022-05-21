#include <iostream>
#include <istream>
#include <vector>


using namespace std;

void solution(vector<int> arr, int n, int c){
    int total = 0;
    for(int i: arr) total += i;
    cout << "Case #" << c << ": " <<(total % n) << "\n";
}

int main(){
    int N, M, O, P, Q;
    cin >> N;
    
    vector<string> bin, forbidden;
    string ice;

    while(N--){
        cin >> M >> O >> P;
        while(M--){
            cin >> ice;
            bin.push_back(ice);
        }
        while(O--){
            cin >> ice;
            forbidden.push_back(ice);
        }
        // solution(temp, O, Q);
        Q++;
        // for(int i : temp) cout << i << " ";
        // cout << "\n";

    }

    return 0;
}