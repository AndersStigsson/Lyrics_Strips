<script setup>
import { reactive, ref, computed, toRefs, onMounted, watch } from 'vue';
import axios from 'axios';
import Line from './Line.vue';
import TrackInfo from './TrackInfo.vue';
// import ExtraInfo from './ExtraInfo.vue';

const props = defineProps(['seedGenres'])


const line = ref({});
const trackInfo = ref({});
const state = reactive({guessing: true});
const getUrl = () => {
    return 
}
const lines = ref([]);
const lineNumber = ref(0);
const finished = ref(false);
const error = ref(false);
const loading = ref(false);
const errorText = ref('');
const totalTracks = ref(20);
const changedGenres = ref(false);
const nextSong = async (changeState) => {
    error.value = false;
    errorText.value = '';
    loading.value = true;
    if (changeState) {
        clickedLine();
    }
    try {

        const resp = await axios.post(`${import.meta.env.VITE_BACKEND_HOST}/search?total=${totalTracks.value}`,
            {
                values: props.seedGenres,
                type: ''
            }
        )
        if (!resp.data.lyrics || !resp.data.track) {
            throw new Error("Either lyrics or track not found")
        }
        trackInfo.value = resp.data.track;
        lines.value = resp.data.lyrics;
        if (resp.data.total > totalTracks.value || changedGenres.value) {
            if (changedGenres.value) {
                changedGenres.value = false;
            }
            totalTracks.value = resp.data.total;
        }
        lineNumber.value = resp.data.lineNumber;
        line.value = lines.value[lineNumber.value];
    } catch (e) {
        console.error(e);
        error.value = true;
        errorText.value = e;
    }
    loading.value = false;
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
        totalTracks.value = 20;
        changedGenres.value;
    }
)

const lineTime = computed(() => {
    if (line.value && line.value.time) {
        return line.value.time.split('.')[0]
    }
    return '00:00';
});

</script>

<template>
    <div v-if="loading">
        <div class="loading text-white">
            Loading...
        </div>
    </div>
    <div
        v-else-if="error"
    >
        <div class="text-white">
            {{ errorText }}
        </div>
        <button
            class="btn-xl text-gray-400"
            @click="nextSong(false)"
        >
            Get new
        </button>
    </div>
    <div v-else-if="finished">
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

        <!-- <ExtraInfo 
            v-if="trackInfo"
            :track="trackInfo"
            :time="lineTime"
/> -->
    </div>
</template>
