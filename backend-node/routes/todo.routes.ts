import { Router } from "express";
import { TodoView } from "../views/todo.views";

var express = require('express');
var todo: TodoView = new TodoView();

var routes: Router = express.Router();
routes.get('/', todo.index);
routes.get('/todos', todo.getAll);
routes.post('/todo', todo.add)
routes.get('/todo/:id', todo.get);
routes.delete('/todo/:id', todo.delete);
routes.put('/todo/:id', todo.update);

export default routes