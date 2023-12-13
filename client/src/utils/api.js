import axios from "axios";
import { useEffect, useState } from "react";

const baseUrl = "http://127.0.0.1:5000/api/announcements";

const useFetch = (id = 0) => {
    let url = baseUrl;
    if (id > 0) {
        url += `/${id}`;
    }
    const [ data, setData ] = useState([]);
    const [ err, setErr ] = useState(null);
    const [ isPending, setIsPending ] = useState(true);

    useEffect(() => {
        axios.get(url)
            .then((response) => {
                return response.data.data;
            })
            .then((data) => {
                setData(data);
            })
            .catch((err) => {
                setErr(err);
            })
            .finally(() => {
                setIsPending(false);
            })
    }, [ ]);

    return { data, isPending, err };
}

const fetchData = (id = 0) => {
    let url = baseUrl;
    if (id > 0) {
        url += `/${id}`;
    }

    return axios.get(url)  
        .then((response) => (response.data));
}

const submitForm = (method, id = 0, body = {}) => {
    let url = baseUrl;
    if (id > 0) {
        url += `/${id}`;
    }

    return axios({
        method: method,
        url,
        data: body
    })
    .then((response) => (response.data));
}

export { useFetch, submitForm, fetchData };