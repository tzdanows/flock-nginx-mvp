'use client'
import { useState } from 'react';
import { createConnectTransport } from "@connectrpc/connect-web";
import { createClient } from "@connectrpc/connect";
import { PingService } from "../gen/src/proto/ping_pb";

// setup baseURL as localhost:8080 or kubernetes secret url
const baseUrl = process.env.NEXT_PUBLIC_BACKEND_URL || "http://localhost:8080";
]const transport = createConnectTransport({
  baseUrl: baseUrl,
});

const client = createClient(PingService, transport);

export default function PingButton() {
  const [messages, setMessages] = useState<string[]>([]); // Changed to array
  const [isLoading, setIsLoading] = useState(false);

  const handleClick = async () => {
    setIsLoading(true);
    try {
      const response = await client.ping({
        message: "Hello from frontend!"
      });
      setMessages(prevMessages => [response.message, ...prevMessages]); // Add to array
    } catch (error) {
      console.error('Error:', error);
      setMessages(prevMessages => ['Error connecting to server', ...prevMessages]);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="space-y-4">
      <button
        type="button"
        onClick={handleClick}
        disabled={isLoading}
        className="px-4 py-2 bg-gray-800 text-white rounded-md hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 disabled:opacity-50 transition-colors duration-200"
      >
        {isLoading ? 'Sending...' : 'Send Ping'}
      </button>

      <div className="space-y-2">
        {messages.map((message, index) => (
          <div 
            key={index} 
            className="p-4 bg-gray-50 rounded-md border border-gray-200 shadow-sm"
          >
            <p className="text-gray-700 text-sm">{message}</p>
          </div>
        ))}
      </div>
    </div>
  );
}