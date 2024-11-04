<template>
  <div>
    <ul>
      <li v-for="prefecture in prefectures" :key="prefecture.id">
        {{ prefecture.name }}
      </li>
    </ul>
  </div>
</template>


<script lang="ts">
import { defineComponent, ref } from 'vue';
import { fetchPrefectures } from '../hooks/prefectures';
import type { Prefecture } from '../types/types';

export default defineComponent({
  setup() {
      const prefectures = ref<Prefecture[]>([]);
      const error = ref<string | null>(null);

      const getPrefectures = async () => {
          try {
              prefectures.value = await fetchPrefectures();
          } catch (err) {
              error.value = 'データの取得に失敗しました。';
          }
      };

      getPrefectures();
    
      return {
          prefectures,
          error,
      };
  },
});
</script>
