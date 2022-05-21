#include <iostream>
#include <iterator>
#include <algorithm>
#include <string>
#include <cstdio>
#include <set>
#include <vector>

using namespace std;

int main(){
    ios::sync_with_stdio(0);
    cin.tie(0); 

    int N, a, b;
    cin >> N; 
    
    set<string> setA, setB; 
    string textA, textB;

    while(N--){
        cin >> a;
        while(a--){
            cin >> textA;
            setA.insert(textA);
        }
        
        cin >> b;
        while(b--){
            cin >> textB;
            setB.insert(textB);
        }
    }

    vector<string> result;

    set_difference(setA.begin(), setA.end(), setB.begin(), setB.end(), inserter(result, result.begin()));

    // debug 

    // cout << "i like!\n";
    // for_each(setA.begin(), setA.end(), [](string x){
    //     cout << x << " ";
    // });

    // cout << "\ni dislike!\n";

    // for_each(setB.begin(), setB.end(), [](string x){
    //     cout << x << " ";
    // });

    // cout << "\nresult!!\n";

    cout << result.size() << " ";
    for_each(result.begin(), result.end(), [](string x){
        cout << x << " ";
    });

    return 0;
}