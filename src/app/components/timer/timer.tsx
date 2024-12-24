"use client";

import React, { useState, useEffect } from "react";

export interface TimerProps {
  id: number;
  title: string;
  backgroundColor: string;
}

export const Timer: React.FC<TimerProps> = ({ title, backgroundColor }) => {
  const [time, setTime] = useState(0); // Time in milliseconds
  const [isRunning, setIsRunning] = useState(false);

  useEffect(() => {
    let interval: NodeJS.Timeout | undefined;

    if (isRunning) {
      // Increment time by 10ms every 10ms
      interval = setInterval(() => {
        setTime((prevTime) => prevTime + 10);
      }, 10);
    } else if (!isRunning && interval) {
      clearInterval(interval);
    }

    return () => {
      if (interval) clearInterval(interval); // Cleanup on component unmount
    };
  }, [isRunning]);

  const formatTime = () => {
    const milliseconds = (time % 1000) / 10;
    const seconds = Math.floor((time / 1000) % 60);
    const minutes = Math.floor((time / (1000 * 60)) % 60);

    return `${String(minutes).padStart(2, "0")}:${String(seconds).padStart(2, "0")}.${String(milliseconds).padStart(2, "0")}`;
  };
  return (
    <div className={`m-2 rounded ${backgroundColor}`}>
      <div className="m-2 py-1">
        <h1 className="font-bold">{title}</h1>
        <h1>{formatTime()}</h1>

        <button
          onClick={() => {
            setIsRunning(!isRunning);
          }}
          className={`mr-1 ${isRunning ? "text-red-700" : "text-green-700"}`}
        >
          <p>{isRunning ? "Stop" : "Start"}</p>
        </button>

        <button
          onClick={() => {
            setIsRunning(false);
            setTime(0);
          }}
        >
          <p>Reset</p>
        </button>
      </div>
    </div>
  );
};
