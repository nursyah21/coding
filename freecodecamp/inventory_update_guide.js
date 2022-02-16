//@ts-check
function updateInventory(arr1, arr2) {
    /* update value arr1 */
    arr1.forEach(e1 => {
        arr2.forEach(e2 => {
            if (e2[1] == e1[1]) e1[0] += e2[0]
        })
    })


    let oldarr = [];
    let newarr = [];

    arr1.forEach(e => {
        oldarr.push(e[1])
    });

    arr2.forEach(e => {
        newarr.push(e[1])
    })

    newarr.forEach(e => {
        if (oldarr.indexOf(e) == -1) {
            let index = newarr.indexOf(e);
            arr1.push(arr2[index])
        }
    })

    return arr1.sort((a, b) => { return a[1] > b[1] ? 1 : -1 })
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