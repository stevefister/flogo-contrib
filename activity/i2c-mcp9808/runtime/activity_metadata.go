package mcp9808

var jsonMetadata = `{
  "name": "mcp9808",
  "version": "0.0.1",
  "title": "Read MCP9808",
  "description": "Read Temperature from MCP9808",
  "homepage": "https://github.com/TIBCOSoftware/flogo-contrib/tree/master/activity/i2c-mcp9808",

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

`
