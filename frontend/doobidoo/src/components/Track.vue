<script setup>
import { reactive, ref } from 'vue';
import axios from 'axios';
import Line from './Line.vue';
import TrackInfo from './TrackInfo.vue';

let line = ref('');
let trackInfo = ref({});
let state = reactive({guessing: true});

const URL = `${import.meta.env.VITE_BACKEND_HOST}/next`;
let resp = await axios.get(URL);
trackInfo = ref(resp.data.track);
let lines = ref(resp.data.lines);
let lineNumber = ref(resp.data.lineNumber);

const clickedLine = () => {
    state.guessing = !state.guessing;
}

const nextSong = async () => {
    resp = await axios.get(URL)
    trackInfo = resp.data.track;
    lines = resp.data.lines;
    lineNumber = resp.data.lineNumber;
    line = lines[lineNumber].words;
    clickedLine();
}

</script>

<template>
    <div>
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
</template>
