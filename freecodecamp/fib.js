function f(n) {
    return !(n % 2)
}

var a = [1, 3, 54, 3, 2, 1, 4, 6, 4, 3, 5, 2]
let b = a.filter(f)

console.log(b);