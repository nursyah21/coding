//@ts-check
function updateInventory(arr1, arr2) {
    function comp(e1, e2) {
        if (e1[1] > e2[1]) return 1
        if (e1[1] < e2[1]) return -1
        return 0
    }
    arr1.sort(comp)
    arr2.sort(comp)

    let arr = []

    arr1.forEach(e1 => {
        arr.push(e1);
        arr2.forEach(e2 => {
            let temp = []
            if (e1[1] == e2[1]) {
                arr.pop()
                arr.push([e1[0] + e2[0], e1[1]])
            }
        });
    })
    var arr2list = []
    arr2.forEach(e => {
        arr2list.push(e[1])
    })

    return arr
}


// Example inventory lists
var curInv = [
    [21, "Bowling Ball"],
    [2, "Dirty Sock"],
    [1, "Hair Pin"],
    [5, "Microphone"]
];

var newInv = [
    [2, "Hair Pin"],
    [3, "Half-Eaten Apple"],
    [67, "Bowling Ball"],
    [7, "Toothpaste"]
];


console.log(updateInventory(curInv, newInv));