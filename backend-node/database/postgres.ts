import { Pool } from "pg";

const host = process.env.DBHOST || 'localhost';

const pool = new Pool({
    host: host,
    user: 'root',
    password: 'root',
    database: 'docker-todo-db',
    port: 5432
});

export default pool;