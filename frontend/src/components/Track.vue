<script setup>
import { reactive, ref, computed, toRefs, onMounted, watch } from 'vue';
import axios from 'axios';
import Line from './Line.vue';
import TrackInfo from './TrackInfo.vue';

const props = defineProps(['seedGenres'])


const pageNumber = ref(0);
const line = ref('');
const trackInfo = ref({});
const state = reactive({guessing: true});
const getUrl = () => {
    return 
}
const lines = ref([]);
const lineNumber = ref(0);
const finished = ref(false);
const nextSong = async (changeState) => {
    const resp = await axios.post(`${import.meta.env.VITE_BACKEND_HOST}/search?page=${pageNumber.value+1}`,
        {
            values: props.seedGenres,
            type: ''
        }
    )
    pageNumber.value += 1;
    trackInfo.value = resp.data.track;
    lines.value = resp.data.lyrics;
    lineNumber.value = resp.data.lineNumber;
    line.value = lines.value[lineNumber.value].words;
    if (changeState) {
        clickedLine();
    }
}
const clickedLine = () => {
    state.guessing = !state.guessing;
}
onMounted(async () => {
    // const resp = await axios.get(getUrl());
    // trackInfo = resp.data.track;
    // lines = resp.data.lines;
    // lineNumber = resp.data.lineNumber;
    await nextSong();
    finished.value = true;


});
watch(
    () => props.seedGenres,
    () => {
        pageNumber.value = 0
    }
)

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
            @click="nextSong(true)"
        />
    </div>
    <div
        v-else
    >
        Hej Hej
    </div>
</template>
