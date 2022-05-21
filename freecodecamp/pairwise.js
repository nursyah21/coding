function pairwise(arr, arg){
    let result = []
    /*
    arr.forEach((e,idx)=>{
         arr.forEach((f,idx2)=>{
             if(idx==idx2){}
             else if(e+f == arg){
                 result.push(e)
                 result.push(f)
                
             }
         })
       })*/
    for(var i=0;i<arr.length;i++){
      for(var j=i+1;j<arr.length;j++){
          if (arr[i] + arr[j] == arg){
              result.push(i,j)
          }
      }
    }
    result = result.reduce((a,b)=>{return a+b})
    console.log(result)
}

let test = [
  [[1, 4, 2, 3, 0, 5],7],
  [[1,3,2,4],4],
]

test.forEach(a=>{
      var [b,c] = a
      pairwise(b,c)
    })