const mongoose = require('mongoose')

if (process.argv.length<3) {
    console.log('give password as argument')
    process.exit(1)
}

const password = process.argv[2]

const url = `mongodb+srv://sjt:${password}@fso.9rgsa.mongodb.net/noteApp?retryWrites=true&w=majority&appName=fso`

mongoose.set('strictQuery',false)

mongoose.connect(url)

const noteSchema = new mongoose.Schema({
    content: String,
    important: Boolean,
})

const Note = mongoose.model('Note', noteSchema)

const note = new Note({
    content: 'HTML is easy',
    important: true,
})

const Save = ({note}) => {
    note.save().then(result => {
        console.log('note saved!')
        mongoose.connection.close()
    })
}

// e.g. query = {important: true}
const Find = ({query}) => {
    Note.find(query).then(result => {
        result.forEach(note => {
            console.log(note)
        })
        mongoose.connection.close()
    })
}

Find({})
