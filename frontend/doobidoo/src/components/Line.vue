<script setup>
import { ref, reactive } from 'vue';

const props = defineProps({
  lines: Array,
    lineNumber: Number
})
const emit = defineEmits(['guess'])
const onkeyup = (evt) => {
    if (evt.code == 'Enter' || evt.code == 'Space') {
        emit('guess')
    }
};

let numberOfExtraLines = 0
const lengthOfTrackLines = props.lines.length;
let showLines = reactive([props.lineNumber])
const showMoreLines = () => {
    numberOfExtraLines = numberOfExtraLines + 1;
    showLines.push((props.lineNumber + numberOfExtraLines) % lengthOfTrackLines);
};


window.addEventListener('keyup', onkeyup);
</script>

<template>
    <div 
        class="text-2xl"
    >
        <p
            v-for="lineNumber in showLines"
        >
            <div
                v-if="lineNumber === 0"
                class="text-lg text-gray-400"
            >
                Beginning of song
            </div>
            <div>
                {{ props.lines[lineNumber].words }}
            </div>
        </p>
        <button
            class="btn-xl"
            @click="showMoreLines"
        >
            Show next line
        </button>
        <button 
            class="btn-xl"
            @click="$emit('guess')"
        >
            Guess
        </button>
    </div>
</template>

<style scoped>
</style>
