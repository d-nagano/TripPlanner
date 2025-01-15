<template>
    <div>
        ログイン完了しました。
    </div>
    <div>
        <form @submit.prevent="onSubmit">
            <button type="submit" class="btn btn-primary">Logout</button>
        </form>

        <p v-if="error" class="error">{{ error }}</p>
    </div>
</template>


<script lang="ts">
import { defineComponent, ref } from 'vue';
import { fetchUserLogout } from '../hooks/users';
import { useRouter } from "vue-router";

export default defineComponent({
    setup() {
        const error = ref<string | null>(null);
        const router = useRouter();

        const logout = async () => {
            try {
                await fetchUserLogout();
                alert("ログアウト完了しました。")
                router.push('/');
            } catch (err) {
                error.value = '予期せぬエラーが発生しました';
            }
        };

        const onSubmit = () => {
            logout();
        };

        return {
            error,
            onSubmit,
        };
    },
});
</script>