"use client";

import React, { useState, useEffect } from "react";
import Image from "next/image";

export interface TimerProps {
  id: number;
  title: string;
  totalTimeInSeconds: number;
}

const SECONDS_IN_A_HOUR: number = 3600;
const SECONDS_IN_A_MINUTE: number = 60;
const ICON_WIDTH = 30;
const ICON_HEIGHT = 30;

export const Timer: React.FC<TimerProps> = ({ id, title, totalTimeInSeconds: initialTotal }) => {
  const [totalTimeInSeconds, setTotalTimeInSeconds] = useState(initialTotal);
  const [sessionTimeInSeconds, setSessionTimeInSeconds] = useState(0);
  const [isRunning, setIsRunning] = useState(false);

  useEffect(() => {
    let interval: NodeJS.Timeout | undefined;

    if (isRunning) {
      // Increment time by 1s every 1000ms
      interval = setInterval(() => {
        setSessionTimeInSeconds((prevTime) => prevTime + 1);
      }, 1000);
    } else if (!isRunning && interval) {
      clearInterval(interval);
    }

    return () => {
      if (interval) clearInterval(interval); // Cleanup on component unmount
    };
  }, [isRunning]);

  const formatTime = (): string => {
    const seconds = Math.floor(sessionTimeInSeconds % SECONDS_IN_A_MINUTE);
    const minutes = Math.floor((sessionTimeInSeconds / SECONDS_IN_A_MINUTE) % SECONDS_IN_A_MINUTE);
    const hours = Math.floor(sessionTimeInSeconds / SECONDS_IN_A_HOUR);

    let displayTime = "";
    if (hours != 0) {
      displayTime += `${hours}:`;
    }
    displayTime += `${String(minutes).padStart(2, "0")}:${String(seconds).padStart(2, "0")}`;

    return displayTime;
  };

  const formatTotalTime = (): string => {
    const hours = Math.floor(totalTimeInSeconds / SECONDS_IN_A_HOUR);
    const minutes = Math.floor((totalTimeInSeconds % SECONDS_IN_A_HOUR) / SECONDS_IN_A_MINUTE);
    const seconds = Math.floor((totalTimeInSeconds % SECONDS_IN_A_HOUR) % SECONDS_IN_A_MINUTE);

    let displayTime = "";
    if (hours != 0) {
      displayTime += `${hours}h `;
    }
    if (minutes != 0) {
      displayTime += `${minutes}m `;
    }
    if (seconds != 0) {
      displayTime += `${seconds}s`;
    }

    return displayTime;
  };

  const toggleState = (): void => {
    setIsRunning(!isRunning);
  };

  const getTimerColor = (): string => {
    if (isRunning) {
      return "bg-green-500 dark:bg-green-900";
    } else if (sessionTimeInSeconds == 0) {
      return "bg-zinc-500/40";
    } else {
      return "bg-red-500 dark:bg-red-900";
    }
  };

  const commitTime = async () => {
    try {
      const response = await fetch(`/api/timers/${id}`, {
        method: "POST",
        body: JSON.stringify({ sessionDurationInSeconds: sessionTimeInSeconds }),
      });
      const jsonResponse = await response.json();
      if (response.ok) {
        setIsRunning(false);
        setSessionTimeInSeconds(0);
        setTotalTimeInSeconds((prevTotal) => prevTotal + jsonResponse.sessionDurationInSeconds);
      } else {
        // TODO: Communicate the failure to the user
        console.error("Failed to commit time:", jsonResponse);
      }
    } catch (error) {
      console.error("Parse error: ", error);
    }
  };

  return (
    // TODO: Change this whole thing to be a button to get better default semantics
    <div
      className={`my-5 rounded-sm shadow-md dark:shadow-none ${getTimerColor()}`}
      onClick={toggleState}
    >
      <div className={`m-2 py-3 pl-1`}>
        <span>
          <span className="mr-3 text-xl font-bold">{title}</span>
          <span className="text-sm font-bold text-neutral-800 dark:text-neutral-300">
            {formatTotalTime()}
          </span>
        </span>

        <div className="flex items-center justify-between">
          <h1 className="text-2xl">{formatTime()}</h1>
          {!isRunning && sessionTimeInSeconds != 0 && (
            <button className="ml-auto" onClick={commitTime}>
              <Image
                src="/images/arrow-icon.svg"
                width={ICON_WIDTH}
                height={ICON_HEIGHT}
                alt="An arrow, on click the current session's progress is committed"
                className="dark:invert"
              />
            </button>
          )}
        </div>
      </div>
    </div>
  );
};
