import { QueryResult } from "pg";
import pool from "../database/postgres";

const addTodo = function (args: string[], params: any[], columns: string[]): Promise<QueryResult> {
    return pool.query(`INSERT INTO todos(${columns.join(',')}) VALUES (${args.join(',')})`, params)
}

const deleteTodo = function (id: Number): Promise<QueryResult> {
    return pool.query(`DELETE FROM todos WHERE id = $1`, [id])
}

const fetchAllTodo = function (): Promise<QueryResult> {
    return pool.query(`SELECT * FROM todos`);
}

const fetchTodoByID = function (id: Number): Promise<QueryResult> {
    return pool.query(`SELECT * FROM todos WHERE id = $1`, [id])
}

const updateTodo = function (args: string[], params: any[]): Promise<QueryResult> {
    return pool.query(`UPDATE todos SET ${args.join(',')}, updated_at = now() WHERE id = $1`, params);
}

export { addTodo, deleteTodo, fetchAllTodo, fetchTodoByID, updateTodo };