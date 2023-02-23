<template>
    <q-page>
        <div class="q-pa-md">
            <p class="text-body1 q-py-sm">List of things to do:</p>
            <q-list bordered separator style="max-width: 500px;" class="q-mb-xl">
                <q-item bordered separator clickable ripple v-for="todo in filteredTodos" :key="todo.id"
                    @click="visitUpdate(todo.id)">
                    <q-item-section>{{ todo.content }}</q-item-section>
                    <q-chip class="glossy" color="orange" text-color="white" v-if="todo.finished">
                        Finished
                    </q-chip>
                </q-item>
            </q-list>
            <div>
                <q-btn color="primary" @click="toggleHide">{{ hideFinished ? "Show all" : "Hide completed" }}</q-btn>
                <q-btn color="positive" @click="visitNew" class="q-ml-md">New todo</q-btn>
            </div>
        </div>
    </q-page>
</template>

<script>
import axios from 'axios'
export default {
    computed: {
        filteredTodos() {
            if (this.hideFinished) return this.todos.filter((todo) => { return !todo.finished })
            return this.todos
        }
    },

    data() {
        return {
            hideFinished: false,
            todos: []
        }
    },

    methods: {
        toggleHide() {
            this.hideFinished = !this.hideFinished
        },

        visitNew() {
            this.$router.push('/new')
        },

        visitUpdate(id) {
            this.$router.push(`/update/${id}`)
        }
    },

    mounted() {
        axios
            .get('http://127.0.0.1:3000/todos')
            .then((resp) => {
                this.todos = resp.data
            })
            .catch((err) => {
                console.error(err)
            })
    }
}
</script>