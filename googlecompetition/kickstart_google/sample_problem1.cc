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
    int N, M, O, P, Q=1;
    cin >> N;
    vector<int> temp;


    while(N--){
        cin >> M >> O;
        while(M--){
            cin >> P;
            temp.push_back(P);
        }
        solution(temp, O, Q);
        Q++;
        // for(int i : temp) cout << i << " ";
        // cout << "\n";

        temp.clear();
    }

    return 0;
}