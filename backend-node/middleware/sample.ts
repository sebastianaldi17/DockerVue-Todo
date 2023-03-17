import { Express, NextFunction, Request, Response } from "express";

function logCurrentTime(req: Request, res: Response, next: NextFunction): void {
    console.log(Date.now());
    next();
}

export function applySampleLogging(app: Express): void {
    app.use(logCurrentTime);
}