<template>
    <div class="note">
        <div class="note-header">
            <h2 @input="(e) => note.Title = e.target.innerText" contenteditable="true">{{ props.note.Title }}</h2>
            <div style="width-"></div>
            <svg @click="noteStore.remove(note)" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><!--!Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path fill="currentColor" d="M135.2 17.7L128 32 32 32C14.3 32 0 46.3 0 64S14.3 96 32 96l384 0c17.7 0 32-14.3 32-32s-14.3-32-32-32l-96 0-7.2-14.3C307.4 6.8 296.3 0 284.2 0L163.8 0c-12.1 0-23.2 6.8-28.6 17.7zM416 128L32 128 53.2 467c1.6 25.3 22.6 45 47.9 45l245.8 0c25.3 0 46.3-19.7 47.9-45L416 128z"/></svg>
        </div>
        <p @input="(e) => note.Description = e.target.innerText" contenteditable="true">{{ props.note.Description }}</p>
    </div>
</template>

<script setup>
import { useNotesStore } from '@/stores/notes';
import { reactive, watch } from 'vue';

const props = defineProps(["note"])

const note = reactive(props.note)

const noteStore = useNotesStore()

watch(note, () => {
    noteStore.update(note)
}, {
    deep: true,
    flush: "post"
})


</script>

<style lang="scss" scoped>
.note {
    background-color: var(--color-secondary-background);
    color: var(--color-text);

    padding: 10px;

    border-radius: 15px;
    box-shadow: -6px 6px var(--color-border);

    width: 300px;

    &-header {
        display: flex;
        flex-direction: row;
        & > h2 {
            width: 100%;
            padding: 3px 3px;
            margin-top: 0px;
            margin-right: 5px;

            &:focus{
                outline: solid 3px var(--color-outline);
            }
        }
        & > svg {
            width: 15px;
            height: 15px; 
            color: var(--color-primary);
            &:hover {
                cursor: pointer;
                color: var(--color-secondary);
            }
        }
    }

    & > p {
        text-wrap: wrap;
        padding: 4px;
        &:focus{
            outline: solid 3px var(--color-outline);
            outline-offset: 3px 3px;
        }
    }

}

</style>