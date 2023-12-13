import React from "react";
import { Routes, Route } from "react-router-dom";
import NavigationBar from "./NavigationBar";
import AnnouncementBoard from "./AnnouncementBoard";
import Save from "./Save";
import { submitForm } from "../utils/api";

const App = () => {
    return (
        <div className="App">
            <div className="content">
                <NavigationBar />
                <Routes>
                    <Route path="/" 
                        element={<AnnouncementBoard 
                        pageTitle={ "Announcement Board" }/>}>
                    </Route> 
                    <Route path="/create" 
                        element={<Save apiHandler={ submitForm }
                        pageTitle={ "Create Announcement "} />}>
                        </Route> 
                    <Route path="/:id" 
                        element={<Save apiHandler={ submitForm } 
                        pageTitle={ "Update Announcement" } />}>
                    </Route> 
                </Routes>
            </div>
        </div>
    );
}

export default App;