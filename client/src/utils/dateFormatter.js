const dateFormatter = (dateString) => {
    const options = {
        weekday: "long",
        year: "numeric",
        month: "long",
        day: "numeric"
    }

    return new Date(dateString).toLocaleDateString("en-US", options);
}

export default dateFormatter;