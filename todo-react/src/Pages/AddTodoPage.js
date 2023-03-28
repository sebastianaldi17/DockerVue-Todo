import { Button, ButtonGroup, FormControlLabel, Switch, TextField } from "@mui/material";
import { useState } from "react";
import axios from 'axios'
import { Stack } from "@mui/system";
import { useNavigate } from "react-router-dom";

function AddTodoPage() {
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

    function addTodo() {
        axios
            .post(`http://127.0.0.1:3000/todo`, {
                content: content,
                finished: finished
            })
            .then(() => {
                navigate('/')
            })
            .catch((err) => {
                console.error(err)
                alert(err)
            })
    }

    return (
        <div className="margins">
            <Stack spacing={2} sx={{ width: 1 / 4 }}>
                <TextField id="content" label="Content" variant="outlined" value={content} onChange={(e) => {
                    setContent(e.target.value);
                }} />
                <FormControlLabel control={<Switch onChange={toggleFinished} checked={finished} />} label={switchText} />
                <ButtonGroup>
                    <Button onClick={addTodo}> Submit</Button>
                    <Button onClick={() => { navigate('/') }}> Return</Button>
                </ButtonGroup >
            </Stack >
        </div >
    )
}

export default AddTodoPage;