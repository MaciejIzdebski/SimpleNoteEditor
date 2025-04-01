import { ref, reactive, useId } from 'vue'
import { defineStore } from 'pinia'

function Note(title, description, ID){
  this.title = title
  this.description = description
  this.ID = Symbol('id')
}

export const useNotesStore = defineStore('notes', () => {
  const notes = reactive([new Note("Test", "Lorem ipsum dolor sit amet consectetur adipisicing elit. Tempora, vel esse dicta quaerat amet cumque sequi id ipsum perspiciatis iusto temporibus est tenetur fugiat, officia ea facere, incidunt placeat odit!")])

  function create() {
    notes.push(new Note("test", "..."))
  }

  function remove(note){
    let noteIndex = notes.findIndex((n) => n.ID === note.ID)

    if (noteIndex === -1) {
      throw Error("Could not find an element with id: " + note.ID)
    }
    
    console.log(noteIndex)
    notes.splice(noteIndex,1)
  }

  function update(note) {
    console.log(note)
    let noteIndex = notes.findIndex((n) => n.ID === note.ID)

    if (noteIndex === -1) {
      throw Error("Could not find an element with id: " + note.ID)
    }

    notes[noteIndex] = note
  }

  return { notes, create, remove, update }
})
