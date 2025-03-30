import { ref, reactive } from 'vue'
import { defineStore } from 'pinia'

function Note(title, description){
  this.title = title
  this.description = description
}

export const useNotesStore = defineStore('notes', () => {
  const notes = reactive([new Note("Test", "Lorem ipsum dolor sit amet consectetur adipisicing elit. Tempora, vel esse dicta quaerat amet cumque sequi id ipsum perspiciatis iusto temporibus est tenetur fugiat, officia ea facere, incidunt placeat odit!")])

  function create() {
    notes.push(new Note("test", "..."))
  }

  function remove(note){
    
  }

  function update(note){

  }

  return { notes, create, remove, update }
})
