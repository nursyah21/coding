#include <iostream>
#include <istream>
#include <vector>


using namespace std;

void solution(string n, int c){
    char i = n[n.length()-1];
    string opt[] = {"Bob", "Alice", "nobody"};
    string vowel = "AEIOUaeiou";
    string y = "Yy";

    string result;
    result = ( ( y.find(i) != -1) ? opt[2] : (vowel.find(i) != -1) ? opt[1] : opt[0]); 


    cout << "Case #" << c << ": " << n << " is ruled by " << result << ".\n";
}

int main(){
    int N, O=1;
    string M;
    cin >> N;


    while(N--){
        cin >> M;
        solution(M,O);
        O++;
    }

    return 0;
}