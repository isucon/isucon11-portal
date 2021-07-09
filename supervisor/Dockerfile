FROM rust:1.53

RUN rustup component add rustfmt

WORKDIR /app

COPY ./supervisor /app
COPY ./proto /proto

RUN cargo install --path . --bin isuxportal-supervisor

CMD isuxportal-supervisor cat
