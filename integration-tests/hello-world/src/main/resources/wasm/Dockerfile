FROM rust:1.74

WORKDIR /usr/src/myapp
COPY . .

RUN rustup target add wasm32-unknown-unknown
RUN rustc operation.rs --target=wasm32-unknown-unknown --crate-type=cdylib -C opt-level=3 -C debuginfo=0 -o operation.wasm

ENTRYPOINT [ "cat", "operation.wasm" ]
