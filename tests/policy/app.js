var express = require('express')
var app = express()
var bodyParser = require('body-parser')
app.use(bodyParser.json())

app.get('/mqtt/connect', function (req, res) {
  console.log("/mqtt/connect", req.body)
  res.send('{ }')
})

app.get('/mqtt/publish', function (req, res) {
  console.log("/mqtt/publish", req.body)
  res.send('{ }')
})

app.get('/mqtt/subscribe', function (req, res) {
  console.log("/mqtt/subscribe", req.body)
  res.send('{ }')
})

app.get('/mqtt/receive', function (req, res) {
  console.log("/mqtt/receive", req.body)
  res.send('{ }')
})

app.listen(3000, function () {
  console.log('Example app listening on port 3000!')
})
