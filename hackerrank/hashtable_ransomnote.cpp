#include <algorithm>
#include <iostream>
#include <vector>
#include <string>

using namespace std;

void solution(vector<string> magazine, vector<string> note){
    sort(magazine.begin(), magazine.end());
    sort(note.begin(), note.end());
    int noteword = note.size();
    vector<string> result;

    set_intersection(magazine.begin(), magazine.end(), note.begin(), note.end(), back_inserter(result));
    cout << (result.size() == noteword? "Yes" : "No");
}

int main(){
    int m,n;
    vector<string> magazine, note;
    cin >> m >> n;
    string text;
    while(m--){
        cin >> text;
        magazine.push_back(text);
    }
    while(n--){
        cin >> text;
        note.push_back(text);
    }
    solution(magazine, note);
}
