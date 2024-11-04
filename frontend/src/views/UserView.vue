<template>
    <div>
      <form @submit.prevent="onSubmit">
        <input v-model="user.name" placeholder="名前" required />
        <input v-model="user.email" type="email" placeholder="メールアドレス" required />
        <input v-model="user.password" type="password" placeholder="パスワード" required />
        <button type="submit">サインアップ</button>
      </form>
  
      <p v-if="error" class="error">{{ error }}</p>
    </div>
</template>
  
  
  <script lang="ts">
  import { defineComponent, ref } from 'vue';
  import { fetchUserSignUp } from '../hooks/user';
  import type { User } from '../types/types';
  import axios from 'axios';
  
  export default defineComponent({
    setup() {
        const error = ref<string | null>(null);
        const user = ref({
            name: '',
            email: '',
            password: '',
        });
  
        const signup = async (user: User) => {
            try {
                await fetchUserSignUp(user);
            } catch (err) {
                if (axios.isAxiosError(err) && err.response) {
                    error.value = err.response.data;
                } else {
                    error.value = '予期せぬエラーが発生しました';
                }
            }
        };
  
        const onSubmit = () => {
            signup(user.value);
        };
      
        return {
            user,
            error,
            onSubmit,
        };
    },
  });
  </script>