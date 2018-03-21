//f(x)=x^3-61x^2-860x-800
//[-2,2]
let b = -2
let a = 2
let x

while (a-b > 0.00001) {
    x = (a+b)/2
    if ((a*a*a-61*a*a-860*a-800) * (x*x*x-61*x*x-860*x-800) < 0)
        b = x
    else
        a = x
}

console.log(`Answer is : ${x}`)