<script setup>
import Note from '@/components/Note.vue';
import { useNotesStore } from '@/stores/notes';

const notesStore = useNotesStore()
</script>

<template>
  <main>
    <Note v-for="(note, index) in notesStore.notes" :note="note" :key="note.ID" />
  </main>
</template>

<style>
main {
  display: flex;
  align-content: start;
  justify-content: center;
  flex-direction: row;
  margin: 15px;
  padding: 15px;
  gap: 15px;
  flex-wrap: wrap;
}
</style>
