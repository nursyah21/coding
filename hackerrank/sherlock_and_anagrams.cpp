#include <algorithm>
#include <iostream>
#include <vector>
#include <string>
#include <map>

using namespace std;


void solution(string text){
    
    int result = 0;
    map<string, int> test;

    for(int i=0; i < text.length(); i++){
        string temp = "";
        for(int j=0; j < text.length() - i; j++){
            temp +=  text[i+j];
            string temp2 = temp;
            sort(temp2.begin(), temp2.end());
            test[temp2]++;
            //cout << temp2 << " ";
        }
        //cout << "\n";
    }
    for(auto i : test){
        //cout << i.first << ": " << i.second << "\n";
        if(i.second > 1) result += i.second * (i.second -1) / 2;
    }
    cout << result << "\n";
} 
/*
  5 * 2 = 10
  4 * 1.5= 6
  3 * 1  = 3
  2 *.5= 1
*/



int main(){
    int n;

    cin >> n;
    string text;
    while(n--){
        cin >> text;
        solution(text);
    }
}

