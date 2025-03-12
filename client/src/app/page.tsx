"use client";

import { useEffect, useState } from "react";

import styles from "./page.module.css";
import ThemeToggle from "@/components/theme-toggle";
import { TimerProps } from "@/components/timer";
import TimerList from "@/components/timer-list";

interface BackendTimer {
  id: number;
  title: string;
}

export default function HomePage() {
  const [timers, setTimers] = useState<TimerProps[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch("/api/timers");
        const jsonResponse = await response.json();
        const parsedTimers: TimerProps[] = await jsonResponse.timers.map((timer: BackendTimer) => ({
          id: timer.id,
          title: timer.title,
          totalTime: 115843,
        }));
        setTimers(parsedTimers);
      } catch (error) {
        // TODO: Handle error as desired
        console.error("Parse error: ", error);
      }
    };

    fetchData();
  }, []);

  return (
    <div className={`${styles["grid-container"]} h-lvh grid`}>
      <meta charSet="utf-8"></meta>
      <meta name="viewport" content="width=device-width, initial-scale=1.0"></meta>
      <ThemeToggle />

      <main className="row-start-2 col-start-2">
        <TimerList timers={timers}></TimerList>
      </main>
    </div>
  );
}
