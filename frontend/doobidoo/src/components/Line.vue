<script setup>
const props = defineProps({
  line: String
})
const emit = defineEmits(['click'])
const onkeyup = (evt) => {
    if (evt.code == 'Enter' || evt.code == 'Space') {
        emit('click')
    }
};

window.addEventListener('keyup', onkeyup);
</script>

<template>
    <div 
        class="text-2xl"
    >
        <button 
            class="btn-xl"
            @click="$emit('click')"
        >
            {{ props.line }}
        </button>
    </div>
</template>

<style scoped>
</style>
