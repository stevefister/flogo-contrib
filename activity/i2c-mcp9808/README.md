# i2c-mcp9808
This activity provides your flogo application the ability to read MCP9808 temp sensor on Raspberry Pi

## Installation

```bash
flogo add activity github.com/TIBCOSoftware/flogo-contrib/activity/i2c-mcp9808
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "method",
      "type": "string",
      "required": true,
      "allowed" : ["Direction", "Read State", "Pull"]
    },
    {
      "name": "pinNumber",
      "type": "integer",
      "required": true
    },
    {
      "name": "direction",
      "type": "string",
      "allowed" : ["Input", "Output"]
    },
    {
      "name": "state",
      "type": "string",
      "allowed" : ["High", "Low"]
    },

    {
      "name": "Pull",
      "type": "string",
      "allowed" : ["Up", "Down", "Off"]
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "float"
    }
  ]
}
```
## Settings
| Setting     | Description    |
|:------------|:---------------|
| method      | The method to take action for GPIO|         
| pinNumber   | The pin number   |
| direction   | The direction of pin number, either Input or Output |
| state       | The state of pin number, either high or low |
| Pull        | Pull the pin number to Up, Down and Off |


## Configuration Examples
### Get pin state
Get specific pin 23's state
```json
  "attributes": [
          {
            "name": "method",
            "value": "Read State",
            "type": "string"
          },
          {
            "name": "pinNumber",
            "value": "23",
            "type": "integer"
          }
        ]
```

### Change pin's direction
Change pin's direction to Output
```json
  "attributes": [
          {
            "name": "method",
            "value": "Direction",
            "type": "string"
          },
          {
            "name": "pinNumber",
            "value": "23",
            "type": "integer"
          },
          {
            "name": "direction",
            "value": "Output",
            "type": "string"
          }
        ]
```
