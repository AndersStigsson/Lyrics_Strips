<script setup>
const props = defineProps({
  trackInfo: Object
})
const emit = defineEmits(['click']);
const onkeyup = (evt) => {
    if (evt.code == 'Enter' || evt.code == 'Space') {
        emit('click')
    }
};

window.addEventListener('keyup', onkeyup);

</script>

<template>
    <div>
        <div
            v-for="artist in props.trackInfo.artists"
            class="text-2xl text-zinc-200"
        >
            {{ artist.name }}
        </div>
        <div 
            class="text-xl text-zinc-400"
        >
            {{ props.trackInfo.name }}
        </div>
        <button
            class="btn-xl text-gray-400"
            @click="$emit('click')"
        >
            Next
        </button>
    </div>
</template>
