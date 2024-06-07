import { Container, Stack} from "@chakra-ui/react";
import Navbar from "./component/Navbar";
import TodoForm from "./component/TodoForm.tsx";
import TodoList from "./component/TodoList.tsx";

export const BASE_URL = import.meta.env.MODE === "development" ? "http://localhost:5000/api" : "/api";

function App() {
    return (
       <Stack h="100vh">
           <Navbar />
           <Container>
               <TodoForm />
               <TodoList />
           </Container>
       </Stack>
    )
}

export default App
