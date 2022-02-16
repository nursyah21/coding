#include <algorithm>
#include <iostream>
#include <vector>
#include <string>


using namespace std;


void solution(string text1, string text2){
    
    bool twostring = false;
    vector<char> result;
    sort(text1.begin(), text1.end());
    sort(text2.begin(), text2.end());

    set_intersection(text1.begin(), text1.end(), text2.begin(), text2.end(), back_inserter(result));
    
    cout << (result.size() > 0 ? "YES" : "NO") << "\n";
}

int main(){
    int n;

    cin >> n;
    string text1,text2;
    while(n--){
        cin >> text1 >> text2;
        
        solution(text1, text2);

    }
}

