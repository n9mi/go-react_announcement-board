import React from "react";
import dateFormatter from "../utils/dateFormatter";
import { useNavigate } from "react-router-dom";

const Announcement = ({ id, title, content, created_at, updated_at}) => {
    const navigate = useNavigate();

    return (
        <div className="announcement">
            <div className="announcement__title"> 
                { title }
                <button className="announcement__action" onClick={() => navigate(`/${id}`) } >
                    <i className="fa-regular fa-pen-to-square"></i>
                </button> 
            </div>
            <div className="announcement__date"> { dateFormatter(created_at)} </div>
            <div className="announcement__content"> { content } </div>
        </div>
    );
};

export default Announcement;