<script setup>
import Note from '@/components/Note.vue';
import { useNotesStore } from '@/stores/notes';

const notesStore = useNotesStore()
</script>

<template>
  <main>
    <Note v-for="note in notesStore.notes" :model="note" />
    
  </main>
</template>

<style>
main {
  display: flex;
  align-items: center center;
  flex-direction: column;
  margin: 15px;
  padding: 15px;
  gap: 7px;

}
</style>
