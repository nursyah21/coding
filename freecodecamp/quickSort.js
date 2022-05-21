function quickSort(arr){
   if(arr.length <= 1)return arr
    
   let pivot = arr[0]
   let greater = []
   let lesser = []
   
    
   for(var i=1; i<arr.length;i++){
      if(arr[i] > pivot)
         greater.push(arr[i])
      else 
         lesser.push(arr[i])
   }
   
   sorted = [...quickSort(lesser), pivot, ...quickSort(greater)]
   return sorted
}

let test = [
  [1,4,2,8,345,123,43,32,5643,63,123,43,2,55,1,234,92],
  
]

test.forEach(a=>{
      //var [b,c] = a
      var d = quickSort(a)
      console.log(d)
    })