function f(n){
    if(n <= 0) return 1;
    return f(n-2) + f(n-1);
}

console.log(f(2));