import { Express } from 'express';

const cors = require('cors');

export function applyCors(app: Express): void {
    app.use(cors({
        origin: '*'
    }));
}