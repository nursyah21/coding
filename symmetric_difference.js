function sym(...args) {

    function symdiff(args1, args2) {
        let _intersection = new Set(args1);
        new Set(args2).forEach(e => {
            if (_intersection.has(e)) _intersection.delete(e);
            else _intersection.add(e);
        });
        return _intersection;
    }

    let realresult = [];
    args.reduce(symdiff).forEach(e => {
        realresult.push(e)
    });

    return realresult.sort();
}

console.log(sym([1, 2, 3], [5, 2, 1, 4])); //desired result [3,4,5]
console.log(sym([1, 2, 3], [5, 2, 1, 4, 5])); //desired result [3,4,5]