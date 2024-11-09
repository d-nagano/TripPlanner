<template>
    <div>
        <form>
            <div class="mb-3">
                <label class="form-label">Email address</label>
                <input v-model="user.email" type="email" class="form-control" aria-describedby="emailHelp" required />
                <div id="emailHelp" class="form-text">We'll never share your email with anyone else.</div>
            </div>
            <div class="mb-3">
                <label class="form-label">Password</label>
                <input v-model="user.password" type="password" class="form-control" required />
            </div>
            <button type="submit" class="btn btn-primary">Login</button>
        </form>

        <p v-if="error" class="error">{{ error }}</p>
    </div>
</template>


<script lang="ts">
import { defineComponent, ref } from 'vue';

export default defineComponent({
    setup() {
        const error = ref<string | null>(null);
        const user = ref({
            email: '',
            password: '',
        });

        return {
            user,
            error,
        };
    },
});
</script>
