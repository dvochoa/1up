"use client";

import styles from "./timer.module.css";
import React, { useState, useEffect } from "react";
import Image from "next/image";

export interface TimerProps {
  id: number;
  title: string;
  totalTimeInSeconds: number;
  onTimerDeleted: () => void;
}

const SECONDS_IN_A_HOUR: number = 3600;
const SECONDS_IN_A_MINUTE: number = 60;
const ICON_WIDTH = 18;
const ICON_HEIGHT = 18;

export const Timer: React.FC<TimerProps> = ({
  id,
  title,
  totalTimeInSeconds: initialTotal,
  onTimerDeleted,
}) => {
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

  const toggleIsRunning = (): void => {
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

  const deleteTimer = async (e: React.MouseEvent) => {
    e.stopPropagation(); // Surrounding timer div is clickable so stop propagation of click event
    try {
      const response = await fetch(`/api/timers/${id}`, {
        method: "DELETE",
      });
      if (response.ok) {
        onTimerDeleted();
      } else {
        // TODO: Communicate the failure to the user
        console.error(
          "Error Code %d encountered when attempting to delete timer: ",
          response.status,
        );
      }
    } catch (error) {
      console.error("Parse error: ", error);
    }
  };

  const resetSession = (e: React.MouseEvent) => {
    e.stopPropagation(); // Surrounding timer div is clickable so stop propagation of click event
    setSessionTimeInSeconds(0);
    setIsRunning(false);
  };

  const commitSession = async (e: React.MouseEvent) => {
    e.stopPropagation(); // Surrounding timer div is clickable so stop propagation of click event
    try {
      const response = await fetch(`/api/timers/${id}`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
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
    <div
      className={`cursor-pointer rounded-sm shadow-md dark:shadow-none ${getTimerColor()}`}
      onClick={toggleIsRunning}
    >
      <div className={`${styles["grid-container"]} h-full p-2 pb-3 pl-3`}>
        <span className="col-1 row-1">
          <span className="mr-3 text-xl">{title}</span>
          <span className="text-xs text-neutral-700 dark:text-neutral-300">
            {formatTotalTime()}
          </span>
        </span>

        {!isRunning && (
          <button className="col-2 row-1 ml-auto cursor-pointer" onClick={deleteTimer}>
            <Image
              src="/images/delete-timer-icon.svg"
              width={ICON_WIDTH}
              height={ICON_HEIGHT}
              alt="A trash can, on click this timer is deleted"
              className="dark:invert"
            />
          </button>
        )}

        <h1 className="col-1 row-7 text-4xl font-bold">{formatTime()}</h1>

        <button
          className={`col-2 row-5 ml-auto cursor-pointer ${isRunning || sessionTimeInSeconds == 0 ? "pointer-events-none invisible" : ""}`}
          onClick={resetSession}
        >
          <Image
            src="/images/reset-session-icon.svg"
            width={ICON_WIDTH}
            height={ICON_HEIGHT}
            alt="A refresh, on click resets the current session back to zero"
            className="dark:invert"
          />
        </button>

        <button
          className={`col-2 row-9 ml-auto cursor-pointer ${isRunning || sessionTimeInSeconds == 0 ? "pointer-events-none invisible" : ""}`}
          onClick={commitSession}
        >
          <Image
            src="/images/commit-session-icon.svg"
            width={ICON_WIDTH}
            height={ICON_HEIGHT}
            alt="An arrow, on click the current session's progress is committed"
            className="dark:invert"
          />
        </button>
      </div>
    </div>
  );
};
