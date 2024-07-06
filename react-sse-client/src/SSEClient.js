import React, { useEffect, useState } from 'react';

const SSEClient = () => {
    const [messages, setMessages] = useState([]);

    useEffect(() => {
        const eventSource = new EventSource('/events');

        eventSource.onmessage = function (event) {
            setMessages((messages) => [...messages, event.data]);
        };

        return () => {
            eventSource.close();
        };
    }, []);

    return (
        <div>
            <h1>Server-Sent Events</h1>
            <ul>
                {messages.map((message, index) => (
                    <li key={index}>{message}</li>
                ))}
            </ul>
        </div>
    );
};

export default SSEClient;