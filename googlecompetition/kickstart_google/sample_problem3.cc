#include <iostream>
#include <istream>
#include <vector>
#include <algorithm>

using namespace std;

void hindex(vector<int> &arr, int n){
    int h=0;

    for(int i=0; i < n; i++){
        if (arr[i] >= n) h++;
    }

    cout << h << " ";
}

void solution(vector<int> &arr, int n, int c){
    int total = 0, h=1;

    cout << "Case #" << c << ": ";

    for(int i=0; i < arr.size(); i++){
        hindex(arr, i);
    }
        

    cout << "\n";

    //  <<(total % n) << "\n";
}

int main(){
    int N, M, O, P, Q=1;
    cin >> N;
    vector<int> temp;


    while(N--){
        cin >> M;
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