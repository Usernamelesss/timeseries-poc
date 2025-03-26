import {getTimeseries} from "./load-file.js";
import {benchDivideBy2, benchSqrt} from "./ts-operations.js";

const main =  async () => {
    const ts = await getTimeseries()

    benchDivideBy2(ts)
    benchSqrt(ts)
}

await main()