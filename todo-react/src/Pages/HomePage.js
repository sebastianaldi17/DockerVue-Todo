import Stack from '@mui/material/Stack';
import { Button, ButtonGroup } from "@mui/material";
import '../App.css'
import { useEffect, useState } from 'react';
import axios from 'axios';
import DoneIcon from '@mui/icons-material/Done';
import { useNavigate } from "react-router-dom";

function HomePage() {
    const [todos, setTodos] = useState([]);
    const [showHidden, setShowHidden] = useState(true);
    const navigate = useNavigate();

    useEffect(() => {
        axios
            .get('http://127.0.0.1:3000/todos')
            .then((resp) => {
                console.log(resp.data);
                setTodos(resp.data);
            })
            .catch((err) => {
                console.error(err);
            });
    }, []);

    return (
        <div className="margins">
            <Stack spacing={2} sx={{ width: 1 / 4 }}>
                <p>List of things to do:</p>
                {
                    todos.map((todo) => {
                        if (todo.finished) {
                            if (!showHidden) {
                                return (<></>)
                            }
                            return (
                                <Button variant="outlined" key={todo.id} startIcon={<DoneIcon />} onClick={() => { navigate(`/update/${todo.id}`) }}>{todo.content}</Button>
                            )
                        }
                        return (
                            <Button variant="outlined" key={todo.id} onClick={() => { navigate(`/update/${todo.id}`) }}>{todo.content}</Button>
                        )
                    })
                }
                <ButtonGroup>
                    <Button onClick={() => { setShowHidden(!showHidden) }}> Hide completed</Button>
                    <Button onClick={() => { navigate('/new') }}> Add todo</Button>
                </ButtonGroup >
            </Stack>
        </div >
    );
}

export default HomePage;