<script setup>
import { reactive, ref } from 'vue';
import axios from 'axios';
import Line from './Line.vue';
import TrackInfo from './TrackInfo.vue';

let line = ref('');
let trackInfo = ref({});
let state = reactive({guessing: true});

const URL = 'http://localhost:10010/next';
let resp = await axios.get(URL);
trackInfo = ref(resp.data.track);
line = ref(resp.data.line.words)

const clickedLine = () => {
    state.guessing = !state.guessing;
}

const nextSong = async () => {
    resp = await axios.get(URL)
    trackInfo = resp.data.track;
    line = resp.data.line.words
    
    clickedLine();
}

</script>

<template>
    <div>
        <Line 
            v-if="state.guessing"
            :line="line"
            @click="clickedLine"
        />
        <TrackInfo
            v-else
            :trackInfo="trackInfo"
            @click="nextSong"
        />
        {{ guessing }}
    </div>
</template>
