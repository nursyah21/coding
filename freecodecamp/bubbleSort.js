function bubbleSort(arr){
   var notsorted = true
   while(notsorted){
       notsorted=false
      for(var i=0, j=1; i<arr.length;i++,j++){
        if(arr[i]>arr[j]){
            var temp = arr[i]
            arr[i] = arr[j]
            arr[j] = temp
            notsorted=true
        }
      }
   }
   return arr
}

let test = [
  [1,4,2,8,345,123,43,32,5643,63,123,43,2,55,1,234,92],
  
]

test.forEach(a=>{
      //var [b,c] = a
      var d = bubbleSort(a)
      console.log(d)
    })