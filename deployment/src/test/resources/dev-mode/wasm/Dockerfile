FROM rust:1.74

WORKDIR /usr/src/myapp
COPY . .

RUN rustup target add wasm32-unknown-unknown

ARG SOURCE_FILE
RUN rustc ${SOURCE_FILE} --target=wasm32-unknown-unknown --crate-type=cdylib -C opt-level=3 -C debuginfo=0 -o math.wasm

ENTRYPOINT [ "cat", "math.wasm" ]
