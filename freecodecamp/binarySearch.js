function binarySearch(arr,target){
    let res = []
    
    function binarySearchRecursive (arr, x, low = 0, high = arr.length - 1) {
      const mid = Math.floor(low + (high - low) / 2)

      if (high >= low) {
        if (arr[mid] === x) {
        // item found => return its index
          res.push(arr[mid])
          return
         }

        if (x < arr[mid]) {
          // arr[mid] is an upper bound for x, so if x is in arr => low <= x < mid
          res.push(arr[mid])
          binarySearchRecursive(arr, x, low, mid - 1)
        } else {
          res.push(arr[mid])
          // arr[mid] is a lower bound for x, so if x is in arr => mid < x <= high
          binarySearchRecursive(arr, x, mid + 1, high)
        } 
        res = []
        return
      }
    }
    
    binarySearchRecursive(arr, target, 0, arr.length)
   // throw console.log(res)
    return res.length > 0 ? res : "Value Not Found"
}

let test = [
  [[1,4,2,8,342],3],
  
]

test.forEach(a=>{
      var [b,c] = a
      var d = binarySearch(b, c)
      console.log(d)
    })