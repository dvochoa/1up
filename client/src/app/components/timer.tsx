"use client";

import React, { useState, useEffect } from "react";

export interface TimerProps {
  id: number;
  title: string;
  totalTime: number; // in seconds
}

const SECONDS_IN_A_HOUR: number = 3600;
const SECONDS_IN_A_MINUTE: number = 60;

export const Timer: React.FC<TimerProps> = ({ title, totalTime }) => {
  const [time, setTime] = useState(0); // Time in seconds
  const [isRunning, setIsRunning] = useState(false);

  useEffect(() => {
    let interval: NodeJS.Timeout | undefined;

    if (isRunning) {
      // Increment time by 1s every 1000ms
      interval = setInterval(() => {
        setTime((prevTime) => prevTime + 1);
      }, 1000);
    } else if (!isRunning && interval) {
      clearInterval(interval);
    }

    return () => {
      if (interval) clearInterval(interval); // Cleanup on component unmount
    };
  }, [isRunning]);

  const formatTime = (): string => {
    const seconds = Math.floor(time % 60);
    const minutes = Math.floor((time / 60) % 60);

    return `${String(minutes).padStart(2, "0")}:${String(seconds).padStart(2, "0")}`;
  };

  const formatTotalTime = (): string => {
    const hours = Math.floor(totalTime / SECONDS_IN_A_HOUR);
    const minutes = Math.floor((totalTime % SECONDS_IN_A_HOUR) / SECONDS_IN_A_MINUTE);
    const seconds = Math.floor((totalTime % SECONDS_IN_A_HOUR) % SECONDS_IN_A_MINUTE);

    return `${hours}h ${minutes}m ${seconds}s`;
  };

  const toggleState = (): void => {
    setIsRunning(!isRunning);
  };

  const getTimerColor = (): string => {
    if (isRunning) {
      return "bg-green-500 dark:bg-green-900";
    } else if (time == 0) {
      return "bg-zinc-500/40";
    } else {
      return "bg-red-500 dark:bg-red-900";
    }
  };

  return (
    <div
      className={`my-5 rounded shadow-md dark:shadow-none ${getTimerColor()}`}
      onClick={toggleState}
    >
      <div className="m-2 py-3 pl-1">
        <span className="text-xl font-bold mr-3">{title}</span>
        <span className="text-sm font-bold text-neutral-800 dark:text-neutral-300">
          {formatTotalTime()}
        </span>

        <h1 className="text-2xl">{formatTime()}</h1>
      </div>
    </div>
  );
};
