<template>
    <div>
        <form @submit.prevent="onSubmit">
            <div class="mb-3">
                <label class="form-label">Name</label>
                <input v-model="user.name" class="form-control" aria-describedby="emailHelp" required />
                <div id="emailHelp" class="form-text">We'll never share your email with anyone else.</div>
            </div>
            <div class="mb-3">
                <label class="form-label">Email address</label>
                <input v-model="user.email" type="email" class="form-control" aria-describedby="emailHelp" required />
                <div id="emailHelp" class="form-text">We'll never share your email with anyone else.</div>
            </div>
            <div class="mb-3">
                <label class="form-label">Password</label>
                <input v-model="user.password" type="password" class="form-control" required />
            </div>
            <button type="submit" class="btn btn-primary">Submit</button>
        </form>

        <p v-if="error" class="error">{{ error }}</p>
    </div>
</template>


<script lang="ts">
import { defineComponent, ref } from 'vue';
import { fetchUserSignUp } from '../hooks/users';
import type { User } from '../types/users.ts';
import axios from 'axios';
import { useRouter } from "vue-router";

export default defineComponent({
    setup() {
        const error = ref<string | null>(null);
        const user = ref({
            name: '',
            email: '',
            password: '',
        });
        const router = useRouter();

        const signup = async (user: User) => {
            try {
                await fetchUserSignUp(user);
                alert("ユーザーの登録が完了しました。")
                router.push('/login');
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