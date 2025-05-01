"use client";

import { useEffect, useState } from "react";

import styles from "./page.module.css";
import { Timer } from "./models/timers";
import { TimerProps } from "./components/timer";
import TimerList from "./components/timer-list";

export default function TimersOverviewPage() {
  const [timers, setTimers] = useState<TimerProps[]>([]);

  const fetchTimers = async () => {
    try {
      const response = await fetch("/api/users/1/timers");
      const jsonResponse = await response.json();
      const parsedTimers: TimerProps[] = jsonResponse.timers.map((timer: Timer) => ({
        id: timer.id,
        title: timer.title,
        totalTimeInSeconds: timer.totalTimeInSeconds,
      }));
      setTimers(parsedTimers);
    } catch (error) {
      // TODO: Handle error as desired
      console.error("Parse error: ", error);
    }
  };

  useEffect(() => {
    fetchTimers();
  }, []);

  return (
    <div className={`${styles["grid-container"]}`}>
      <meta charSet="utf-8"></meta>
      <meta name="viewport" content="width=device-width, initial-scale=1.0"></meta>
      <main className="col-start-2">
        <TimerList timers={timers} onTimerCreated={fetchTimers} onTimerDeleted={fetchTimers} />
      </main>
    </div>
  );
}
