function mergeSort(arr){
    //always divide array to half until size remain 2
    
    if (arr.length < 2)return arr
    
    //merge 2 array already sorted
    function mergeSortHelper(arr1,arr2){
        let res = []
        let i=0
        let j=0
        //throw console.log(arr1, arr2)
        while((i < arr1.length) && (j < arr2.length)){
          if(arr1[i] > arr2[j])res.push(arr2[j++])
          else res.push(arr1[i++])
        }
        return res.concat(arr1.slice(i), arr2.slice(j))
    }
    
    let half = Math.floor(arr.length / 2)
    //console.log(half,arr)
    let leftarr = arr.slice(0,half)
    
    let rightarr = arr.slice(half, arr.length)
    
    
    return mergeSortHelper(mergeSort(leftarr), mergeSort(rightarr))
}

let test = [
  [1,4,2,8,345,123,43,32,5643,63,123,43,2,55,1,234,92],
  
]

test.forEach(a=>{
      //var [b,c] = a
      var d = mergeSort(a)
      console.log(d)
    })