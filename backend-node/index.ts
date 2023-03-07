import express, { Express, Request, Response } from 'express';
import dotenv from 'dotenv';
import routes from './routes/todo.routes';

dotenv.config();

const app: Express = express();

const port = process.env.PORT;

const cors = require('cors');
app.use(cors({
    origin: '*'
}));
app.use(express.json());
app.use(routes);

app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
});
