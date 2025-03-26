import {TimeSeries, toMatrix} from "./load-file.js";

export function benchDivideBy2(data: TimeSeries) {
    const m = toMatrix(data)

    let start = Date.now()
    const result = []
    for (let i = 0; i < m.length; i++) {
        const rowValues = []
        for (let j = 0; j < m[i].length; j++) {
            rowValues.push(m[i][j] / 2)
        }
        result.push(rowValues)
    }
    let end = Date.now()
    let elapsedMs = end - start
    console.log(`Took ${elapsedMs} ms`)
}

export function benchSqrt(data: TimeSeries) {
    const m = toMatrix(data)

    let start = Date.now()
    const result = []
    for (let i = 0; i < m.length; i++) {
        const rowValues = []
        for (let j = 0; j < m[i].length; j++) {
            let n = m[i][j]
            rowValues.push(
                Math.sqrt(n) + n/2 + Math.pow(n, 2) + Math.cos(n % 360)
            )
        }
        result.push(rowValues)
    }
    let end = Date.now()
    let elapsedMs = end - start
    console.log(`Took ${elapsedMs} ms`)
}
