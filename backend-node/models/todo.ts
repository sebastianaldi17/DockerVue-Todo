import { z } from "zod";

const TodoSchema = z.object({
    content: z.string().nullish(),
    status: z.number().nullish(),
    finished: z.boolean().nullish(),
});

export default TodoSchema;