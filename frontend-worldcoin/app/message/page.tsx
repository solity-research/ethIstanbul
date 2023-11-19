"use client"

import { useState } from "react";

interface Message {
    side: number;
    message: string;
}

const MessageScreen = () => {
    const [messages, setMessages] = useState([{ side: 0, message: "How can I help you ?" }])

    return (
        <>
            <div className="flex-1 h-full p-10 flex flex-col bg-gray-200">
                <div className="flex-1 overflow-y-auto p-4">
                    <div className="flex items-end mb-4">
                        {messages.map((message) => {
                            if (message.side == 0) {
                                return (<>
                                    <div className="flex flex-col space-y-2 text-xs max-w-xs mx-2 order-2 items-end">
                                       <div><span className="px-4 py-2 rounded-lg inline-block rounded-bl-none bg-gray-600 text-white ">{message.message}</span></div> 
                                    </div>
                                    <img src="path_to_sender_pic.jpg" alt="Sender" className="w-6 h-6 rounded-full order-1" />
                                </>)
                            } else {
                                return (
                                    <>
                                        <div className="flex flex-col space-y-2 text-xs max-w-xs mx-2 order-2 items-start">
                                           <div><span className="px-4 py-2 rounded-lg inline-block rounded-bl-none bg-blue-600 text-white ">{message.message}</span></div> 
                                        </div>
                                        <img src="path_to_sender_pic.jpg" alt="Sender" className="w-6 h-6 rounded-full order-1" />
                                    </>
                                )
                            }
                        })

                        }

                    </div>
                </div>
                <div className="border-t p-4">
                    <input type="text" placeholder="Write a message..." className="w-full p-2 border rounded" />
                </div>
            </div>
        </>
    )
}

export default MessageScreen