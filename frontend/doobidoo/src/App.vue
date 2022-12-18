<script setup>
import {ref} from 'vue';
import Track from './components/Track.vue';
import Dropdown from './Dropdown.vue';
let seedGenres = ref([]);

const changedGenres = (evt) => {
    seedGenres.value = evt
};
</script>

<template>
    <div>
        <div class="w-52 mx-auto">
            <Dropdown 
                :seed-genres="seedGenres"
                @genre="changedGenres"
            />
        </div>
        <Track 
            :seed-genres="seedGenres" 
        />
    </div>
</template>

