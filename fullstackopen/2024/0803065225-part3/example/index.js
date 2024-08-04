const express = require('express')
const exp = express()
/* const http = require('http')
const app = http.createServer((request, response) => {
  response.writeHead(200, { 'Content-Type': 'text/plain' })
  response.end('Hello World')
})
let notes = {}
const app = http.createServer((request, response) => {
  response.writeHead(200, {'Content-Type': 'application/json'})
  response.end(JSON.stringify(notes))
})*/

let notes = [
  {
    id: 1,
    content: 'hello, world'
  }
]

exp.get('/', (req, resp) => {
  resp.send('<h1>hello, world</h1>')
})

exp.get('/api/notes', (req, resp) => {
  resp.json(notes)
})

const PORT = 3001

exp.listen(PORT, () => {
  console.log(`[log] running on port ${PORT}`)
})
