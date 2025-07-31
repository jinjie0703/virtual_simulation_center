<template>
  <div class="member-grid">
    <MemberCard
      v-for="member in members"
      :key="member.id"
      :member="member"
      @click="onMemberClick(member)"
    />
  </div>
</template>

<script setup>
import MemberCard from './MemberCard.vue'

defineProps({
  members: {
    type: Array,
    required: true,
  },
})

const emit = defineEmits(['memberClick'])

const onMemberClick = (member) => {
  emit('memberClick', member)
}
</script>

<style scoped>
.member-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 40px;
  padding: 40px 0;
}

@media (max-width: 768px) {
  .member-grid {
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 30px;
  }
}
</style>
