# post-date
This activity provides your flogo application the ability to post flow data to a specified URL

## Installation

```bash
flogo add activity github.com/stevefister/flogo-contrib/activity/post-data
```

## Schema
Inputs and Outputs:

```json
{
   "inputs":[
    {
      "name": "method",
      "type": "string",
      "required": "true"
    },
    {
      "name": "uri",
      "type": "string",
      "required": "true"
    },
    {
      "name": "data",
      "type": "string"
    }
    
  ],
  "outputs": [
    {
      "name": "result",
      "type": "any"
    }
  
  ]
}
```

## Configuration Information
