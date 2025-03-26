import {asyncBufferFromFile, parquetReadObjects} from "hyparquet";

const projectRoot = process.env.PROJECT_ROOT || "";
const fileName = `${projectRoot}/fixtures/sample_001.parquet`


export type DataRow = {
    param_0: number,
    param_1: number,
    param_2: number,
    param_3: number,
    param_4: number,
    param_5: number,
    param_6: number,
    param_7: number,
    param_8: number,
    param_9: number
}

export type TimeSeries = DataRow[]

export function toMatrix(ts: TimeSeries): number[][] {
    const result = []

    for (let i = 0; i < ts.length; i++) {
        result[i] = [ts[i].param_0, ts[i].param_1, ts[i].param_2, ts[i].param_3, ts[i].param_4, ts[i].param_5,
            ts[i].param_6, ts[i].param_7, ts[i].param_8, ts[i].param_9]
    }
    return result
}

export async function getTimeseries(): Promise<TimeSeries> {
    const file = await asyncBufferFromFile(fileName)
    const data = await parquetReadObjects({file}) as TimeSeries
    console.log(`Loaded ${data.length} datapoint`)
    return data
}

