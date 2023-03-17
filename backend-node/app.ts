import express, { Express } from 'express';

import routes from './routes/todo.routes';
import { applyCors } from './middleware/cors';
import { applySampleLogging } from './middleware/sample';

const app: Express = express();

applyCors(app);
applySampleLogging(app);
app.use(express.json());
app.use(routes);

export default app;