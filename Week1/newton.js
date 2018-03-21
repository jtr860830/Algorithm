//f(x) = x^3-61x^2-860x-800
const func1 = x => x*x*x-61*x*x-860*x-800
const func2 = x => 3*x*x-122*x-860
const nt = x => {
    while(Math.abs(func1(x)-0) > 0.00001) {
        x = x-func1(x)/func2(x)
    }
    return x
}

let x, y, z

y = nt(90)
z = nt(-20)
console.log(`Answer is : ${z}\n`)
for (let i=-20; i<=90; i+=50) {
    x = nt(i)

    if (x == z) {
        continue
    } else if (x == y) {
        continue
    } else {
        console.log(`Answer is : ${x}\n`)
        break
    }
}

console.log(`Answer is : ${y}\n`)