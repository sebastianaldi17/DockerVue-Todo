import { Button, ButtonGroup, FormControlLabel, Switch, TextField } from "@mui/material";
import { useEffect, useState } from "react";
import { useLoaderData } from "react-router-dom";
import axios from 'axios'
import { Stack } from "@mui/system";
import { useNavigate } from "react-router-dom";
export async function loader({ params }) {
    return params.id;
}

export function TodoPage() {
    const id = useLoaderData()
    const navigate = useNavigate();
    const [content, setContent] = useState('');
    const [finished, setFinished] = useState(false);
    const [switchText, setSwitchText] = useState('Unfinished');

    function toggleFinished(e) {
        setFinished(e.target.checked);
        if (e.target.checked) {
            setSwitchText('Finished');
        } else {
            setSwitchText('Unfinished');
        }
    }

    function updateTodo() {
        axios
            .put(`http://127.0.0.1:3000/todo/${id}`, {
                content: content,
                finished: finished
            })
            .then(() => {
                navigate('/');
            })
            .catch((err) => {
                console.error(err);
                alert(err);
            })
    }

    function deleteTodo() {
        axios
            .delete(`http://127.0.0.1:3000/todo/${id}`)
            .then(() => {
                navigate('/');
            })
            .catch((err) => {
                console.error(err);
                alert(err);
            })
    }

    useEffect(() => {
        axios
            .get(`http://127.0.0.1:3000/todo/${id}`)
            .then((res) => {
                console.log(res);
                setContent(res.data.content);
                setFinished(res.data.finished);
                if (res.data.finished) {
                    setSwitchText('Finished');
                }
            })
            .catch((err) => {
                console.error(err);
            })
    }, [id])

    return (
        <div className="margins">
            <Stack spacing={2} sx={{ width: 1 / 4 }}>
                <TextField id="content" label="Content" variant="outlined" value={content} onChange={(e) => {
                    setContent(e.target.value);
                }} />
                <FormControlLabel control={<Switch onChange={toggleFinished} checked={finished} />} label={switchText} />
                <ButtonGroup>
                    <Button onClick={updateTodo}> Submit</Button>
                    <Button onClick={deleteTodo}> Delete</Button>
                    <Button onClick={() => { navigate('/') }}> Return</Button>
                </ButtonGroup >
            </Stack >
        </div >
    )
}
