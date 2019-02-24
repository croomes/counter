# Counter

Counter is a simple API for controlling an [Adafruit 4-digit 7-segment
display](https://www.adafruit.com/product/1270).

Tested on Arduino Uno and Raspberry Pi Gen 1.

## Usage

Running `./counter` will start an HTTP server running on port `8001`.

### Raspberry Pi

```bash
./counter
```

### Arduino

If connecting to the board via an Arduino, you must pass Arduino's device:

```bash
./counter -device /dev/tty.usbmodem1421
```

See [Gobot Arduino docs](https://gobot.io/documentation/platforms/arduino/) for
more info.

## API

There is no security/authentication on the API.

### Set Number

This sets number from 0 - 9999 on the display:

```bash
curl -X POST http://localhost:8001/number -d '{"value": 1234}'
```

### Set Single Digit

Setting a single digit requires the position on the LED, and the digit's value.
Positions are from 0 - 3, from left to right. Other digits will keep their
existing value.

```bash
curl -X POST http://localhost:8001/digit -d '{"position": 1, "value": 9}'
```

The example above would set the hundreds digit (2nd from left) to 9.

### Brightness

Set the display brightness.  Takes a number fronm 0 - 15, 0 being off, 15 full
power.

```bash
curl -X POST http://localhost:8001/brightness -d '{"value": 10}'
```

### Colon

Toggles display of the colon between the two sets of digits.

```bash
curl -X POST http://localhost:8001/colon -d '{"value": true}'
curl -X POST http://localhost:8001/colon -d '{"value": false}'
```

### Clear Display

Clear the display contents:

```bash
curl -X DELETE http://localhost:8001/
```

### Power

Controls turning display power on/off:

```bash
curl -X POST http://localhost:8001/power -d '{"value": false}'
curl -X POST http://localhost:8001/power -d '{"value": true}'
```

## Building

Requires the HT16K33 Gobot driver, not currently upstreamed (still some failing
tests).

```bash
cd ~/go/src
mkdir -p gobot.io/x
cd gobot.io/x
git clone https://github.com/croomes/gobot.git
cd gobot
git checkout ht16k33
```

To generate for multiple platforms:

```bash
make build
```

Alternatively, run `go build` on the target platform.