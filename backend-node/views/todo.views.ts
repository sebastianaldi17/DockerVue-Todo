import { Request, Response } from "express";
import { addTodo, deleteTodo, fetchAllTodo, fetchTodoByID, updateTodo } from "../controllers/todo.controllers";
import TodoSchema from "../models/todo";

export class TodoView {
    public index(req: Request, res: Response) {
        res.send("Hello World!");
    }

    public getAll(req: Request, res: Response) {
        fetchAllTodo().then(result => {
            res.status(200).json(result.rows);
        }).catch(error => {
            console.error(error);
            res.status(500).send(error.message);
        })
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
        fetchTodoByID(id).then(result => {
            if (result.rows.length > 0) {
                res.status(200).json(result.rows[0]);
            } else {
                res.status(404).json({});
            }
        }).catch(error => {
            console.error(error);
            res.status(500).send(error.message);
            return;
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
        addTodo(args, params, columns).then(result => {
            res.status(200).send("Added to db");
        }).catch(error => {
            console.error(error);
            res.status(500).send(error.message);
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

        deleteTodo(id).then(result => {
            if (result.rowCount <= 0) {
                res.status(404).send("Todo not found");
                return;
            }

            res.status(200).send("Deleted from db");
        }).catch(error => {
            console.error(error);
            res.status(500).send(error.message);
        });
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

        updateTodo(args, params).then(result => {
            res.status(200).send("Updated to db");
        }).catch(error => {
            console.error(error);
            res.status(500).send(error.message);
        })
    }
}
