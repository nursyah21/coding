function selectionSort(arr){
   for(var i=0; i<arr.length;i++){
      let min = i
       
      for(var j=i+1; j < arr.length ;j++){
        if(arr[min] > arr[j])min=j
        
      }
      
      if(min !== i){
         [arr[i],arr[min]] = [arr[min], arr[i]]
      }
     
   }
   return arr
}

let test = [
  [1,4,2,8,345,123,43,32,5643,63,123,43,2,55,1,234,92],
  
]

test.forEach(a=>{
      //var [b,c] = a
      var d = selectionSort(a)
      console.log(d)
    })