const gd = x => x*x*x-61*x*x-860*x-800

let a = 20
let b = 90
let x, y
let t = 0.382

x = a*(1-t)+(b*t)
y = a*t+b*(1-t)

while (y-x > 0.00001) {
    if (gd(y) > gd(x)) {
        b = y
        y = x
        x = a*(1-t)+(b*t)
    } else {
        a = x
        x = y
        y = (a*t)+b*(1-t)
    }
}

console.log(`Answer is : ${x}`)