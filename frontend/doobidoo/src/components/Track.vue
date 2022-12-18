<script setup>
import { reactive, ref, computed, toRefs, onMounted } from 'vue';
import axios from 'axios';
import Line from './Line.vue';
import TrackInfo from './TrackInfo.vue';

const props = defineProps(['seedGenres'])

let line = ref('');
let trackInfo = ref({});
let state = reactive({guessing: true});
const getUrl = () => {
    if (props.seedGenres.length) {
        return `${import.meta.env.VITE_BACKEND_HOST}/next?seed_genres=${props.seedGenres.join(',')}`
    }
    return `${import.meta.env.VITE_BACKEND_HOST}/next`;
}
let lines = ref([]);
let lineNumber = ref(0);
let finished = ref(false);
const nextSong = async () => {
    let resp = await axios.get(getUrl())
    trackInfo = resp.data.track;
    lines = resp.data.lines;
    lineNumber = resp.data.lineNumber;
    line = lines[lineNumber].words;
    clickedLine();
}
const clickedLine = () => {
    state.guessing = !state.guessing;
}
onMounted(async () => {
    let resp = await axios.get(getUrl());
    trackInfo = resp.data.track;
    lines = resp.data.lines;
    lineNumber = resp.data.lineNumber;
    finished.value = true;


});

</script>

<template>
    <div v-if="finished">
        <Line 
            v-if="state.guessing"
            :lines="lines"
            :line-number="lineNumber"
            @guess="clickedLine"
        />
        <TrackInfo
            v-else
            :trackInfo="trackInfo"
            @click="nextSong"
        />
    </div>
    <div
        v-else
    >
        Hej Hej
    </div>
</template>
