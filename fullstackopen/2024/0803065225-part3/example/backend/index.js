const express = require('express')
const app = express()
const cors = require('cors')

app.use(express.json())
app.use(cors())
app.use(express.static('dist'))

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
        id: "1",
        content: 'hello, world'
    }
]

const generateId = () => {
    const maxId = notes.length > 0
        ? Math.max(...notes.map(n => Number(n.id)))
        : 0
    return String(maxId + 1)
}

app.get('/api/notes', (req, resp) => {
    resp.json(notes)
})

app.get('/api/notes/:id', (req, resp) => {
    const note = notes.find(note => note.id === req.params.id)
    if (note) {
        resp.json(note)
    } else {
        resp.status(404).end()
    }
})

app.delete('/api/notes/:id', (req, resp) => {
    notes = notes.filter(note => note.id !== req.params.id)
    resp.status(204).end()
})

app.post('/api/notes', (req, resp) => {
    if (!req.body.content) {
        return resp.status(400).json({
            error: 'content missing'
        })
    }
    const note = {
        content: req.body.content,
        important: Boolean(body.important) || false,
        id: generateId(),
    }
    notes = notes.concat(note)
    resp.json(note)
})

const PORT = process.env.PORT

app.listen(PORT, () => {
    console.log(`[log] running on port ${PORT}`)
})
