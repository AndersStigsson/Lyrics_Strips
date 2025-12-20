<script setup>
import {ref} from 'vue';
import Track from './components/Track.vue';
import Dropdown from './Dropdown.vue';
let seedGenres = ref(["swedish"]);

const changedGenres = (evt) => {
    seedGenres.value = evt
};
</script>

<template>
    <div>
        <div style="background-color: grey" class="w-52 mx-auto">
            <Dropdown 
                @genre="changedGenres"
            />
        </div>
        <Track 
            :seed-genres="seedGenres" 
        />
    </div>
</template>

