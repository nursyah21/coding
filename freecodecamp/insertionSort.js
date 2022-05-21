function insertionSort(arr){
   for(var i=1; i<arr.length;i++){
       
       let temp = arr[i]
       let j
       
      for(j=i-1; (j>=0 && arr[j]>temp) ;j--){
        arr[j+1] = arr[j]
      }
      arr[j+1] = temp
   }
   return arr
}

let test = [
  [1,4,2,8,345,123,43,32,5643,63,123,43,2,55,1,234,92],
  
]

test.forEach(a=>{
      //var [b,c] = a
      var d = insertionSort(a)
      console.log(d)
    })