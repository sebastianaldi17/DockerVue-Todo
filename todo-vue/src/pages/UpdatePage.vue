<template>
    <q-page>
        <q-form class="q-pa-lg" style="max-width: 500px;">
            <p class="text-h6">Edit todo:</p>
            <q-input filled v-model="content" label="Content" class="q-py-sm" />
            <q-toggle v-model="finished" :label="toggleLabel" @click="changeLabel" />
            <div class="q-py-md">
                <q-btn color="primary" label="Update" @click="updateTodo" />
                <q-btn color="red-10" label="Delete" icon="delete" class="q-ml-md" @click="deleteTodo" />
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
            todoID: 0,
            toggleLabel: "Unfinished",
        }
    },
    methods: {
        changeLabel() {
            this.toggleLabel = !this.finished ? "Unfinished" : "Finished"
        },
        debug() {
            console.log(this.content)
            console.log(this.finished)
        },

        deleteTodo() {
            axios
                .delete(`http://127.0.0.1:3000/todo/${this.todoID}`)
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

        updateTodo() {
            axios
                .put(`http://127.0.0.1:3000/todo/${this.todoID}`, {
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
        }
    },
    mounted() {
        this.todoID = this.$route.params.id
        axios
            .get(`http://127.0.0.1:3000/todo/${this.todoID}`)
            .then((res) => {
                console.log(res)
                this.content = res.data.content
                this.finished = res.data.finished
            })
            .catch((err) => {
                console.error(err)
            })
    }
}
</script>
