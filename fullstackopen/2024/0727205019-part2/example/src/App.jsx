import Note from './components/Note'
import { useState, useEffect } from 'react'
import db from './services/db'
import './index.css'
import Footer from './components/Footer'

const App = () => {
  const [notes, setNotes] = useState([])
  const [newNote, setNewNote] = useState('')
  const [showAll, setShowAll] = useState(true)
  useEffect(() => db.getAll().then(initialNotes => setNotes(initialNotes)), [])
  const addNote = (event) => {
    event.preventDefault()
    var note = {
      id: notes[notes.length-1].id + 1,
      content: newNote,
      important: false
    }
    db.create(note).then(returnedNote => setNotes(notes.concat(returnedNote)))
    setNewNote('')
  }
  const handleNoteChange = (event) => {
    setNewNote(event.target.value)
  }
  const notesToShow = showAll ? notes : notes.filter(note => note.important) //=== true
  const toggleImportanceOf = id => {
    const note = notes.find(n => n.id === id)
    const newNote = {...note, important: !note.important}
    db.update(id, newNote).then(returnedNote => setNotes(notes.map(note => note.id !== id ? note : returnedNote)))
  }

  return (
    <div>
      <h1>Notes</h1>
      <div>
        <button onClick={() => setShowAll(!showAll)}>
          show {showAll ? 'important' : 'all'}
        </button>
      </div>
      <ul>
        {notesToShow.map(note => 
        <Note key={note.id} note={note} toggleImportance={() => toggleImportanceOf(note.id)} />)}
        <form onSubmit={addNote}>
          <input value={newNote} onChange={handleNoteChange} />
          <button type="submit">save</button>
        </form>
      </ul>
      <Footer />
    </div>
  )
}

export default App