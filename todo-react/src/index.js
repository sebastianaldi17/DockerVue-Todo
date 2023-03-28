import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import {
    createBrowserRouter,
    RouterProvider,
} from "react-router-dom";
import HomePage from './Pages/HomePage';
import Header from './Components/Header';
import { TodoPage, loader as todoLoader } from './Pages/TodoPage';
import AddTodoPage from './Pages/AddTodoPage';
import Footer from './Components/Footer';

const router = createBrowserRouter([
    {
        path: '/',
        element: <HomePage />
    },
    {
        path: '/new',
        element: <AddTodoPage />
    },
    {
        path: '/update/:id',
        element: <TodoPage />,
        loader: todoLoader
    },
]);

ReactDOM.createRoot(document.getElementById("root")).render(
    <React.StrictMode>
        <Header />
        <RouterProvider router={router} />
        <Footer />
    </React.StrictMode>
);