<script setup>
import { RouterView } from 'vue-router'
import { useNotesStore } from './stores/notes';

const notes = useNotesStore()
</script>

<template>
  <header>
    <button @click="notes.create()" id="create-note-button">Dodaj notatkę +</button>
  </header>

  <RouterView />

  <footer>
    <h6 class="">Note editor 2025 - Maciej Izdebski</h6>
  </footer>
</template>

<style scoped>
header {
  display: flex;
  flex-direction: column;
  justify-items: center;
  align-items: center;
}

footer {
  width: 100%;
  text-align: center;
  color: wheat;
}

#create-note-button {
  background-color: var(--color-primary);
  color: var(--color-text);
  padding: 10px;
  margin: 5px;
  border-radius: 5px;
  border: none;

  font-family: Roboto, Verdana, Geneva, Tahoma, sans-serif;

  &:hover {
   background-color: var(--color-secondary);
   cursor: pointer;
   align-self: center;
  }
}
</style>
