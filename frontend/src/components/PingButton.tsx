'use client'

import { useState } from 'react';
import { createConnectTransport } from "@connectrpc/connect-web";
import { createClient } from "@connectrpc/connect";
import { PingService } from "../gen/src/proto/ping_pb";

// Create transport and client outside component
const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

const client = createClient(PingService, transport);

export default function PingButton() {
  const [message, setMessage] = useState<string>('');
  const [isLoading, setIsLoading] = useState(false);

  const handleClick = async () => {
    setIsLoading(true);
    try {
      const response = await client.ping({
        message: "Hello from frontend!"
      });
      setMessage(response.message);
    } catch (error) {
      console.error('Error:', error);
      setMessage('Error connecting to server');
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
        className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 disabled:opacity-50"
      >
        Send Ping
      </button>

      {message && (
        <div className="p-4 bg-gray-100 rounded-md">
          <p>{message}</p>
        </div>
      )}
    </div>
  );
}