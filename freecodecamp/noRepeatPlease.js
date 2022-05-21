const test = "aab"

//function permutation(a){
   let result = []

function permutation(arr, currentSize) {
    //if (typeof(arr) == "string")arr=arr.split("")
    if (currentSize == 1) { // recursion base-case (end)
        result.push(arr.join(""));
        return;
    }
    
    for (let i = 0; i < currentSize; i++){
        
        console.log(currentSize, arr,i)
        if (currentSize % 2 == 1) {
            let temp = arr[0];
            arr[0] = arr[currentSize - 1];
            arr[currentSize - 1] = temp;
        } else {
            let temp = arr[i];
            arr[i] = arr[currentSize - 1];
            arr[currentSize - 1] = temp;
        }
        permutation(arr, currentSize - 1);
    }
}

    //permutationHelper(a,a.length)
    
    //return result
//}


//permutation(test.split(""), test.length)



function permAlone(a){
    let res = 0
    
    function norepeatword(b){
        b=b.split("")
        
        
        for(var i=1,old=0; i<b.length;i++,old++){
            if (b[old] == b[i]) return 0
        }
        return 1
    }
    
    
    for(var i=0;i<a.length;i++){
        res += norepeatword(a[i])
    }
    
    
    
    
    return res
}

//let a=permAlone(result)


let test2 = [
//"abfdefa",
"ab",
//"aaa",
"aabb",
//"abcdefa"
]

test2.forEach(e=>{
    
    result=[]
    permutation(e.split(""),e.length)
    console.log(e,permAlone(result))
    })