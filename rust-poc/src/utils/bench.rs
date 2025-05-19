use std::time::{SystemTime, UNIX_EPOCH};
use polars::prelude::DataFrame;

pub(crate) fn simple_bench(name: String, f: fn(&DataFrame) -> DataFrame, arg: &DataFrame) -> DataFrame {
    let start = SystemTime::now().duration_since(UNIX_EPOCH).unwrap();
    let result = f(arg);
    let end = SystemTime::now().duration_since(UNIX_EPOCH).unwrap();
    println!("{} took {} ms", name, (end - start).as_millis());
    return result;
}
