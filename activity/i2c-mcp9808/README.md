# i2c-mcp9808
This activity provides your flogo application the ability to read MCP9808 temp sensor on Raspberry Pi (3)

## Installation

```bash
flogo add activity github.com/stevefister/flogo-contrib/activity/i2c-mcp9808
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[ ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}
```

## Configuration Information
### MCP9808 is Output Only
The MCP9808 uses Bus 1 on the Raspberry Pi 3. The Default address is 0x18.
Wiring used: VDD, GND, SCL, SDA. See https://learn.adafruit.com/mcp9808-temperature-sensor-python-library/hardware for pinouts.

