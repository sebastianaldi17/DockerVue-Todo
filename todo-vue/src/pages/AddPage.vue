<template>
    <q-page>
        <q-form class="q-pa-lg" style="max-width: 500px;">
            <p class="text-h6">Create new todo:</p>
            <q-input filled v-model="content" label="Content" class="q-py-sm" />
            <q-toggle v-model="finished" :label="toggleLabel" @click="changeLabel" />
            <div class="q-py-md">
                <q-btn color="primary" label="Add new" @click="createTodo" />
                <q-btn color="grey-7" label="Return" class="q-ml-md" @click="returnHome" />
            </div>
        </q-form>
    </q-page>
</template>

<script>
import axios from 'axios'
import router from '@/router'
export default {
    data() {
        return {
            content: "",
            finished: false,
            toggleLabel: "Unfinished",
        }
    },
    methods: {
        changeLabel() {
            this.toggleLabel = !this.finished ? "Unfinished" : "Finished"
        },

        createTodo() {
            axios
                .post(`http://127.0.0.1:3000/todo`, {
                    content: this.content,
                    finished: this.finished
                })
                .then(() => {
                    this.$router.push('/')
                })
                .catch((err) => {
                    console.error(err)
                    alert(err)
                })
        },

        returnHome() {
            router.push("/")
        },
    },
}
</script>
