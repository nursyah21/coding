#include <bits/stdc++.h>
using namespace std;



//ONLY CHANGE 3 LINE CODE, DON'T ADD OR REMOVE LINE CODE
void findZigZagSequenceHackerrank(vector < int > a, int n){
    sort(a.begin(), a.end()); // 1 2 3 4 7 8 10 -> 7
    int mid = ((n + 1)/2) - 1; // 3 fix mid for array --> change 1
    swap(a[mid], a[n-1]); // 1 2 3 10 7 8 4

    int st = mid + 1; // 4
    int ed = n -2 ; // 5 --> change 2
    while(st <= ed){ //iterate only last mid array
        swap(a[st], a[ed]);   // reverse
        //iteration-1 1 2 3 4 10 7 8 -> 5, 6
        //iteration-2 1 2 3 10 4 8 7 -> 6 6
        st = st + 1;
        ed = ed - 1; //--> change 3
    }
    for(int i = 0; i < n; i++){
        if(i > 0) cout << " ";
        cout << a[i];
    }
    cout << endl;
}

// this is will failed, if we run in hackerrank eventhough the output is same



// func caesarCipher(s string, k int32) string {
//     // Write your code here
// 	var res []string
// 	//fmt.Printf("%d %d %d %d %d",rune('-'), rune('A'),rune('Z'), rune('a'),rune('z')) //know rune - A Z
// 	for _,v := range []rune(s) {
	
// 		if v >= 65 && v <= 90 { // A-Z
// 			v += k
// 			if v > 'Z' {v -= 26} // back to beginning
// 		}
// 		if v >= 97 && v <= 122 { // a-z
// 			v += k
// 			if v > 'z' {v -= 26} // back to beginning
// 		}

// 		res = append(res, string(v))
// 	}
// 	return strings.Join(res, "")
// }

void findZigZagSequence(vector < int > a, int n){ 
    // nlogn
    sort(a.begin(), a.end()); //sorted array 1, 2, 3, 4, 5
    int mid = (n + 1)/2; //find mid (5+1)/2 -> 3
 
    // 1, 2, 5, 3, 4 // how to like this
    // 1, 2, 5, 4, 3

    swap(a[mid-1], a[n-1]); // change mid arr 1, 2, 5, 4, 3
    for(int i=0; i < mid; i++){ // 0..3
        cout << a[i] << " ";
    }

    vector<int> dec;
    for(int i=mid; i <n; i++){
        dec.push_back(a[i]);
    }
    sort(dec.begin(), dec.end()); //nlogn
    reverse(dec.begin(), dec.end()); //logn

    for(int i=0; i <dec.size(); i++){ //0.. (5-3)
        cout << dec[i];
        if (i != dec.size() - 1) cout << " "; // fixing so in last element didn't show " "
    }
    cout << endl;
}

void testCase() {
    vector<int> a = {2,3,5,1,4}; //5

    findZigZagSequenceHackerrank(a, 5); //vector, size vector (odd)
    //expected 1,2,3,5,4,

    vector<int> b = {1, 2, 3, 4, 5, 6, 7}; //7

    findZigZagSequenceHackerrank(b, 7); //vector, size vector (odd)
    //expected 1,2,3,5,4,
}

void testCase2(){
    int n, x;
    int test_cases;
    cin >> test_cases;

    for(int cs = 1; cs <= test_cases; cs++){
        cin >> n;
        vector < int > a;
        for(int i = 0; i < n; i++){
            cin >> x;
            a.push_back(x);
        }
        findZigZagSequence(a, n);
    }
}

void caesar(string s, int k){
    //s[0] = s[0]+1;
    //if (s[0] > 65) {cout << 'A';}
    for (int i=0; i < s.size(); i++ ){
        //int v = s[i];
        // if (s[i] >= 65 && s[i] <= 90 ){ // A-Z
        //     s[i] = s[i] + k;
        // }
        //s[i] += k;
        if (s[i] >= 65 && s[i] <= 90) {
            s[i] += k;
            if (s[i] > 90) s[i] -= 26;
            int a = s[i];
            cout << a << " ";
        }
        if (s[i] >= 97 && s[i] <= 122) {
            int a = s[i];
            s[i] += k;
            cout << a << " ";
            if (s[i] > 122) s[i] -= 26;
        }
        // (s[i] >= 97 && s[i] <= 122)){ //A-z
        //     s[i] += k;
        // }
        //if (s[i] > 122) s[i] -= 26;

        // if (s[i] >= 65 && s[i] <= 90 ){ // A-Z
        //     s[i] = s[i] + k;
		// 	if (s[i] > 'Z') {s[i] -= 26;} // back to beginning
		// }
		// if (s[i] >= 97 && s[i] <= 122) { // a-z
		// 	s[i] += k;
		// 	if (s[i] > 'z') {s[i] -= 26;} // back to beginning
		// }
                
    }
    cout << s ;
}

int main() {
    caesar("Ciphering.", 26);
    // string j = "zaya";
    // j[0] += 26;
    // char i = 'C';
    // i+26;
    // cout << i  << j;
    //testCase2();
    return 0;
}



