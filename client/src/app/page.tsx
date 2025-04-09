"use client";

import { useEffect, useState } from "react";

import styles from "./page.module.css";
import { TimerProps } from "@/components/timer";
import TimerList from "@/components/timer-list";

interface TimerOverview {
  id: number;
  title: string;
  totalTimeInSeconds: number;
}

export default function HomePage() {
  const [timers, setTimers] = useState<TimerProps[]>([]);

  const fetchTimers = async () => {
    try {
      const response = await fetch("/api/users/1/timers");
      const jsonResponse = await response.json();
      const parsedTimers: TimerProps[] = await jsonResponse.timerOverviews.map(
        (timer: TimerOverview) => ({
          id: timer.id,
          title: timer.title,
          totalTimeInSeconds: timer.totalTimeInSeconds,
        }),
      );
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
