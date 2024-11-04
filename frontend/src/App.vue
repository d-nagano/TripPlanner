<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <div class="wrapper">
      <nav>
        <RouterLink to="/prefecture">都道府県</RouterLink>
        <RouterLink to="/signup">ユーザー登録</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView />
</template>

