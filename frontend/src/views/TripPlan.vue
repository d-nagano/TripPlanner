<template>
    <main>
        <section class="py-5 text-center container">
            <div class="row py-lg-5">
                <div class="col-lg-6 col-md-8 mx-auto">
                    <h1 class="fw-light">My Library</h1>
                    <p class="lead text-body-secondary">Something short and leading about the collection below—its
                        contents,
                        the creator, etc. Make it short and sweet, but not too short so folks don’t simply skip over it
                        entirely.
                    </p>
                    <p>
                        <a href="#" class="btn btn-primary my-2">Register new</a>
                    </p>
                </div>
            </div>
        </section>

        <div class="album py-5 bg-body-tertiary">
            <div class="container">

                <div class="row row-cols-1 row-cols-sm-2 row-cols-md-3 g-3">
                    <div class="col" v-for="tripPlan in tripPlans">
                        <div class="card shadow-sm">
                            <img class="bd-placeholder-img card-img-top" width="100%" height="225"
                                src="../assets/hokkaido.png" role="img" aria-label="Placeholder: Thumbnail"
                                preserveAspectRatio="xMidYMid slice" focusable="false">
                            <div class="card-body">
                                <p class="card-text trip-plan-text">{{ tripPlan.title }}</p>
                                <p class="card-text">{{ tripPlan.departure_date }} ~ {{ tripPlan.arrival_date }}</p>
                                <div class="d-flex justify-content-between align-items-center">
                                    <div class="btn-group">
                                        <button type="button" class="btn btn-sm btn-outline-secondary">Edit</button>
                                    </div>
                                    <small class="text-body-secondary">9 mins</small>
                                </div>
                            </div>
                        </div>
                    </div>


                </div>
            </div>
        </div>

    </main>
</template>


<script lang="ts">
import { defineComponent, ref } from 'vue';
import { fetchTripPlans } from '../hooks/tripPlans';
import type { TripPlan } from "../types/tripPlans.ts";
import axios from 'axios';
import { onMounted } from 'vue'

export default defineComponent({

    setup() {
        const error = ref<string | null>(null);
        const tripPlans = ref<TripPlan[]>([]);
        const loading = ref(false);

        const fetchTripPlanList = async () => {
            loading.value = true;
            await fetchTripPlans()
                .then(data => {
                    tripPlans.value = data
                })
                .catch(err => {
                    console.log(err)
                    if (axios.isAxiosError(err) && err.response) {
                        error.value = err.response.data;
                    } else {
                        error.value = '予期せぬエラーが発生しました';
                    }
                });
        };

        onMounted(() => {
            fetchTripPlanList()
        })

        return {
            tripPlans,
            error,
        };
    },
});
</script>

<style>
.bd-placeholder-img {
    font-size: 1.125rem;
    text-anchor: middle;
    -webkit-user-select: none;
    -moz-user-select: none;
    user-select: none;
}

@media (min-width: 768px) {
    .bd-placeholder-img-lg {
        font-size: 3.5rem;
    }
}

.b-example-divider {
    width: 100%;
    height: 3rem;
    background-color: rgba(0, 0, 0, .1);
    border: solid rgba(0, 0, 0, .15);
    border-width: 1px 0;
    box-shadow: inset 0 .5em 1.5em rgba(0, 0, 0, .1), inset 0 .125em .5em rgba(0, 0, 0, .15);
}

.b-example-vr {
    flex-shrink: 0;
    width: 1.5rem;
    height: 100vh;
}

.bi {
    vertical-align: -.125em;
    fill: currentColor;
}

.nav-scroller {
    position: relative;
    z-index: 2;
    height: 2.75rem;
    overflow-y: hidden;
}

.nav-scroller .nav {
    display: flex;
    flex-wrap: nowrap;
    padding-bottom: 1rem;
    margin-top: -1px;
    overflow-x: auto;
    text-align: center;
    white-space: nowrap;
    -webkit-overflow-scrolling: touch;
}

.btn-bd-primary {
    --bd-violet-bg: #712cf9;
    --bd-violet-rgb: 112.520718, 44.062154, 249.437846;

    --bs-btn-font-weight: 600;
    --bs-btn-color: var(--bs-white);
    --bs-btn-bg: var(--bd-violet-bg);
    --bs-btn-border-color: var(--bd-violet-bg);
    --bs-btn-hover-color: var(--bs-white);
    --bs-btn-hover-bg: #6528e0;
    --bs-btn-hover-border-color: #6528e0;
    --bs-btn-focus-shadow-rgb: var(--bd-violet-rgb);
    --bs-btn-active-color: var(--bs-btn-hover-color);
    --bs-btn-active-bg: #5a23c8;
    --bs-btn-active-border-color: #5a23c8;
}

.bd-mode-toggle {
    z-index: 1500;
}

.trip-plan-text {
    font-size: 1.2rem;
}
</style>
