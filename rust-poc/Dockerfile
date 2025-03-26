FROM rust

# First step: build dependencies
# Copy Cargo.toml
COPY Cargo.toml .
# Create a dummy file, to allow compilation to succeed
RUN mkdir src && touch src/lib.rs
# Run cargo build on a "dummy" project
# This only builds the dependencies
RUN cargo build

# Second step: build local crate(s)
# Copy the actual source files
COPY src src
# Run the final cargo build
# Dependencies are already built here, only local crates are built
RUN cargo build


ENTRYPOINT ["cargo", "run"]