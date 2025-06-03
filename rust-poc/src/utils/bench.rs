use polars::prelude::DataFrame;
use std::time::Instant;

pub(crate) fn simple_bench(
    name: String,
    f: fn(&DataFrame) -> DataFrame,
    arg: &DataFrame,
) -> (DataFrame, u128) {
    let start = Instant::now();
    let result = f(arg);
    let end = Instant::now();
    let elapsed = (end - start).as_millis();
    println!("{} took {} ms", name, elapsed);
    return (result, elapsed);
}
