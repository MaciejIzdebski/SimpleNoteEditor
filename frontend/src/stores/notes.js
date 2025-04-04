import { ref, reactive, useId } from 'vue'
import { defineStore } from 'pinia'

function Note(title, description, ID){
  this.title = title
  this.description = description
  this.ID = Symbol('id')
}

export const useNotesStore = defineStore('notes', () => {

  const notes = reactive([])

  async function getAll() {
    let response;
    try {
      response = await fetch("/api/notes/", { method: "GET" })
    } catch (err) {
      window.alert(err)
      return
    }

    return await response.json()
  }

  getAll().then((noteResp) => notes.push(...noteResp))

  async function create() {
    let response;
    try {
      response = await fetch("/api/notes/", { method: "POST", body: "{}" })
    } catch (err) {
      window.alert(err)
      return
    }
    notes.push(await response.json())
  }

  async function remove(note){
    let response;
    try {
      response = await fetch(`/api/notes/${note.ID}`, { method: "DELETE", body: JSON.stringify(note) })
    } catch (err) {
      window.alert(err)
      return
    }

    notes.findIndex((el) => el.ID === note.ID)
  }

  async function update(note) {
    let response;
    try {
      response = await fetch(`/api/notes/${note.ID}`, { method: "PUT", body: JSON.stringify(note) })
    } catch (err) {
      window.alert(err)
      return
    }
  }

  return { notes, create, remove, update }
})
