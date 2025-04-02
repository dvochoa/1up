"use client";

import { useEffect, useState } from "react";

import styles from "./page.module.css";
import ThemeToggle from "@/components/theme-toggle";
import { TimerProps } from "@/components/timer";
import TimerList from "@/components/timer-list";
import CreateTimer from "@/components/create-timer";

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
    <div className={`${styles["grid-container"]} grid h-lvh`}>
      <meta charSet="utf-8"></meta>
      <meta name="viewport" content="width=device-width, initial-scale=1.0"></meta>
      <ThemeToggle />

      <main className="col-start-2 row-start-2 space-y-5">
        <TimerList timers={timers} />
        <CreateTimer onTimerCreated={fetchTimers} />
      </main>
    </div>
  );
}
