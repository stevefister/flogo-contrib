package post-data

var jsonMetadata = `{
  "name": "postdata",
  "version": "0.0.1",
  "title": "Post Data to a URL:PORT",
  "description": "Post Data to a URL & Port",
  "homepage": "https://github.com/stevefister/flogo-contrib/activity/post-data",

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
      "name": "port",
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
}`