"use client";

import React, { useState, useEffect } from "react";

export interface TimerProps {
  id: number;
  title: string;
  totalTime: number;
}

export const Timer: React.FC<TimerProps> = ({ title, totalTime }) => {
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

  const formatTime = (): string => {
    const milliseconds = (time % 1000) / 10;
    const seconds = Math.floor((time / 1000) % 60);
    const minutes = Math.floor((time / (1000 * 60)) % 60);

    return `${String(minutes).padStart(2, "0")}:${String(seconds).padStart(2, "0")}.${String(milliseconds).padStart(2, "0")}`;
  };

  const toggleState = (): void => {
    setIsRunning(!isRunning);
  };

  const getTimerColor = (): string => {
    if (isRunning) {
      return "bg-green-900";
    } else if (time == 0) {
      return "bg-zinc-400/40";
    } else {
      return "bg-red-900";
    }
  };

  return (
    <div className={`m-2 rounded ${getTimerColor()}`} onClick={toggleState}>
      <div className="m-2 py-3 pl-1">
        <span className="font-bold mr-2">{title}</span>
        <span className="text-xs text-neutral-300">{totalTime}</span>

        <h1>{formatTime()}</h1>
      </div>
    </div>
  );
};
