import { Request, Response } from "express";
import pool from "../database/postgres";
import TodoSchema from "../models/todo";

export class TodoController {
    public index(req: Request, res: Response) {
        res.send("Hello World!");
    }

    public getAll(req: Request, res: Response) {
        pool.query(`SELECT * FROM todos`, (error, results) => {
            if (error) {
                console.error(error);
                res.status(500).send(error.message);
                return;
            }
            res.status(200).json(results.rows);
        });
    }

    public get(req: Request, res: Response) {
        let id = Number(req.params.id)
        if (isNaN(id)) {
            res.status(400).send("ID must be integer");
            return;
        }
        if (id <= 0) {
            res.status(400).send("ID must be bigger than 0");
            return;
        }
        pool.query(`SELECT * FROM todos WHERE id = $1`, [id], (error, results) => {
            if (error) {
                console.error(error);
                res.status(500).send(error.message);
                return;
            }

            if (results.rows.length > 0) {
                res.status(200).json(results.rows[0]);
            } else {
                res.status(404).json({});
            }
        })
    }

    public add(req: Request, res: Response) {
        let parseResult = TodoSchema.safeParse(req.body);
        if (!parseResult.success) {
            res.status(400).send("Invalid request");
            return;
        }

        const { content, status, finished }: { content: string, status: number, finished: boolean } = req.body;
        if (content === undefined || content.length <= 0) {
            res.status(400).send("content must not be empty");
            return;
        }
        let columns = ['content'];
        let args = ['$1'];
        let params: any[] = [content];
        let counter = 2;
        if (status !== undefined) {
            args.push(`$${counter}`);
            params.push(status);
            columns.push('status');
            counter += 1;
        }
        if (finished !== undefined) {
            args.push(`$${counter}`);
            params.push(finished);
            columns.push('finished');
            counter += 1;
        }
        pool.query(`INSERT INTO todos(${columns.join(',')}) VALUES (${args.join(',')})`, params, (error, result) => {
            if (error) {
                console.error(error);
                res.status(500).send(error.message);
                return;
            }

            res.status(200).send("Added to db");
        });
    }

    public delete(req: Request, res: Response) {
        let id = Number(req.params.id)
        if (isNaN(id)) {
            res.status(400).send("ID must be integer");
            return;
        }
        if (id <= 0) {
            res.status(400).send("ID must be bigger than 0");
            return;
        }

        pool.query(`DELETE FROM todos WHERE id = $1`, [id], (error, result) => {
            if (error) {
                console.error(error);
                res.status(500).send(error.message);
                return;
            }

            if (result.rowCount <= 0) {
                res.status(404).send("Todo not found");
                return;
            }

            res.status(200).send("Deleted from db");
        })
    }

    public update(req: Request, res: Response) {
        let id = Number(req.params.id);
        if (isNaN(id)) {
            res.status(400).send("ID must be integer");
            return;
        }
        if (id <= 0) {
            res.status(400).send("ID must be bigger than 0");
            return;
        }

        let parseResult = TodoSchema.safeParse(req.body);
        if (!parseResult.success) {
            res.status(400).send("Invalid request");
            return;
        }

        const { content, status, finished }: { content: string, status: number, finished: boolean } = req.body;
        let args = [];
        let params: any[] = [id];
        let counter = 2;
        if (content !== undefined) {
            args.push(`content = $${counter}`);
            params.push(content);
            counter += 1;
        }
        if (status !== undefined) {
            args.push(`status = $${counter}`);
            params.push(status);
            counter += 1;
        }
        if (finished !== undefined) {
            args.push(`finished = $${counter}`);
            params.push(finished);
            counter += 1;
        }

        pool.query(`UPDATE todos SET ${args.join(',')}, updated_at = now() WHERE id = $1`, params, (error, result) => {
            if (error) {
                console.error(error);
                res.status(500).send(error.message);
                return;
            }

            res.status(200).send("Updated to db");
        });
    }
}
