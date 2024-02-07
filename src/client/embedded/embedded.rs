// Example Rust code to blink an LED on the micro:bit V2

#![no_std]
#![no_main]

use cortex_m_rt::entry;
use microbit::hal::gpio::Level;
use microbit::hal::prelude::*;

#[entry]
fn main() -> ! {
    if let Some(p) = microbit::Peripherals::take() {
        let gpio = p.GPIO.split();

        let mut led = gpio.pin24.into_push_pull_output(Level::Low);

        loop {
            led.set_high().unwrap();
            cortex_m::asm::delay(1_000_000);
            led.set_low().unwrap();
            cortex_m::asm::delay(1_000_000);
        }
    }

    loop {}
}
