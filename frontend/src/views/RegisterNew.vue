<template>
    <main>
        <div class="py-5 text-center">
            <h2>新しいしおり作成</h2>
            <p class="lead">Below is an example form built entirely with Bootstrap’s form controls. Each required form
                group has a validation state that can be triggered by attempting to submit the form without completing
                it.</p>
        </div>

        <form class="row g-5">
            <div class="col-md-5 col-lg-4 order-md-last">
                <h4 class="d-flex justify-content-between align-items-center mb-3">
                    <span class="text-primary">Profile</span>
                </h4>
                <div class="card flex flex-col items-center gap-6">
                    <FileUpload mode="basic" @select="onFileSelect" customUpload auto severity="secondary"
                        class="p-button-outlined" />
                    <img v-if="src" :src="src" alt="Image" class="shadow-md rounded-xl w-full sm:w-64" />
                </div>
            </div>
            <div class="col-md-7 col-lg-8">
                <div class="col-md-5 col-lg-4 order-md-last"></div>
                <h4 class="mb-3">旅のしおり</h4>
                <div class="row g-3">
                    <div class="col-12">
                        <label class="form-label">タイトル</label>
                        <input type="text" class="form-control" v-model="tripPlan.title" id="title" placeholder="北海道旅行"
                            required>
                    </div>

                    <div class="col-12">
                        <label class="form-label">目的地 </label>
                        <input type="text" class="form-control" id="destination" v-model="tripPlan.destination"
                            placeholder="Tokyo">
                    </div>

                    <div class="col-sm-6">
                        <label class="form-label">出発日</label>
                        <input type="date" class="form-control" id="startDate" v-model="tripPlan.startDate" required>
                    </div>

                    <div class="col-sm-6">
                        <label class="form-label">帰着日</label>
                        <input type="date" class="form-control" id="endDate" v-model="tripPlan.endDate" required>
                    </div>

                    <hr class="my-4">

                    <button class="w-100 btn btn-primary btn-lg" type="button" @click="onSubmit">submit</button>
                </div>
            </div>
        </form>
    </main>
</template>

<script setup>
import { ref } from "vue";
import FileUpload from 'primevue/fileupload';
import { registerTripPlan, uploadImage } from '../hooks/tripPlans'
import { useRouter } from "vue-router";
import axios from 'axios';

const src = ref(null);
const tripPlan = ref({
    title: '',
    destination: '',
    startDate: '',
    endDate: '',
});
const router = useRouter();

function onFileSelect(event) {
    const file = event.files[0];
    src.value = URL.createObjectURL(file);
}

const onSubmit = async () => {
    try {
        const tripId = await registerTripPlan(tripPlan.value);

        // ToDO:ファイルアップロード機能の追加
        // if (selectedFile.value) {
        // await uploadImage(tripId);
        // }

        await router.push('/trip-plan');
    } catch (error) {
        if (axios.isAxiosError(err) && err.response) {
            error.value = err.response.data;
            console.log(err.value)
        } else {
            error.value = '予期せぬエラーが発生しました';
            console.log(err.value)
        }
    }
};
</script>