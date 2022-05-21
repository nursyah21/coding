#include <algorithm>
#include <iostream>
#include <vector>
#include <string>


using namespace std;


void solution(string text1, string text2){
    
    bool twostring = false;
    for(int i=0; i < text1.length(); i++){
        for(int j=0; j < text2.length(); j++){
            if(text2[j] == text1[i]){
                twostring = true;
                break;
            }
        }
        if(twostring) break;
    }
    cout << (twostring ? "YES" : "NO") << "\n";
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

