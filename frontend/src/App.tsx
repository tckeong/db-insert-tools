import { BrowserRouter, Routes, Route } from "react-router-dom";
import Index from "./pages";
import InsertPage from "./pages/insertPage";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Index />} />
                <Route path="/insert" element={<InsertPage />} />
            </Routes>
        </BrowserRouter>
    )
}

export default App
