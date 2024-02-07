@echo off

REM Set the path to your Rust project directory
set RUST_PROJECT_DIR=C:\path\to\your\rust\project

REM Set the path to your micro:bit V2 flashing tool (e.g., uf2conv)
set UF2CONV_PATH=C:\path\to\uf2conv.exe

REM Change directory to your Rust project directory
cd %RUST_PROJECT_DIR%

REM Clean any previous build artifacts
cargo clean

REM Build the Rust code
cargo build --target thumbv7em-none-eabihf --release

REM Convert the built binary to UF2 format
%UF2CONV_PATH% target/thumbv7em-none-eabihf/release/your_project_name.uf2

REM Flash the UF2 file to the micro:bit V2 (assuming it's mounted as a removable drive)
copy /Y target/thumbv7em-none-eabihf/release/your_project_name.uf2 F:\

REM Additional commands to handle deployment if necessary

REM Optional: Open the micro:bit V2's serial port for debugging
REM mode COM4 BAUD=115200 PARITY=n DATA=8
