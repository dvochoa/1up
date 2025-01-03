"use client";

import { useEffect, useState } from "react";

import { TimerProps } from "@/components/timer/timer";
import TimerList from "@/components/timer-list/timer-list";

interface BackendTimer {
  id: number;
  title: string;
}

export default function HomePage() {
  const [timers, setTimers] = useState<TimerProps[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch("http://localhost:8080/timers");
        const jsonResponse = await response.json();
        const parsedTimers: TimerProps[] = await jsonResponse.timers.map((timer: BackendTimer) => ({
          id: timer.id,
          title: timer.title,
          backgroundColor: "bg-orange-500/50",
        }));
        setTimers(parsedTimers);
      } catch (error) {
        // TODO: Handle error as desired
        console.log(error);
      }
    };

    fetchData();
  }, []);

  return (
    <div>
      <meta charSet="utf-8"></meta>
      <meta name="viewport" content="width=device-width, initial-scale=1.0"></meta>
      <main>
        <TimerList timers={timers}></TimerList>
      </main>
    </div>
  );
}
