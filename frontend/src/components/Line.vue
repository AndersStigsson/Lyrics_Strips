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
    } else if (evt.code == 'KeyN') {
        showMoreLines()
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
        <div
            v-for="lineNumber in showLines"
            class="text-white"
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
        </div>
        <button
            class="btn-xl text-gray-400"
            @click="showMoreLines"
        >
            Show next line
        </button>
        <button 
            class="btn-xl text-gray-400"
            @click="$emit('guess')"
        >
            Guess
        </button>
    </div>
</template>
