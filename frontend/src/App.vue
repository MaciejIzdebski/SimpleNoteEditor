<script setup>
import { RouterView } from 'vue-router'
import { useNotesStore } from './stores/notes';

const notes = useNotesStore()
</script>

<template>
  <header>
    <h3 class="">Note editor 2025 - Maciej Izdebski</h3>
    <button @click="notes.create()" id="create-note-button">Dodaj notatkę +</button>
  </header>

  <RouterView />
</template>

<style scoped>
header {
  display: flex;
  flex-direction: row;
  line-height: 1.5;
  max-height: 100vh;
}

#create-note-button {
  background-color: yellowgreen;
  color: black;
  padding: 10px;
  margin: 5px;
  border-radius: 5px;
  border: none;

  font-family: Roboto, Verdana, Geneva, Tahoma, sans-serif;
  &:hover {
   filter: brightness(60%);
   cursor: pointer;
   align-self: center;
  }
}
</style>
