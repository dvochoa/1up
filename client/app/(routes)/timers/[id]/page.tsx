"use client";

import { useEffect, useState } from "react";
import { use } from "react";

import styles from "./page.module.css";
import { TimerDetails } from "../../../models/timers";

export default function TimersOverviewPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);

  const [timerDetails, setTimerDetails] = useState<TimerDetails>();

  const fetchTimerDetails = async () => {
    try {
      const response = await fetch(`/api/timers/${id}`);
      const jsonResponse = await response.json();
      const timerDetails: TimerDetails = {
        timer: jsonResponse.timer,
        timerSessions: jsonResponse.timerSessions,
      };
      console.log(timerDetails);
      setTimerDetails(timerDetails);
    } catch (error) {
      // TODO: Handle error as desired
      console.error("Parse error: ", error);
    }
  };

  useEffect(() => {
    fetchTimerDetails();
  }, []);

  // row-start-2 not working since <main> doesn't have display:grid
  return (
    <div className={`${styles["grid-container"]}`}>
      <meta charSet="utf-8"></meta>
      <meta name="viewport" content="width=device-width, initial-scale=1.0"></meta>
      <main className="col-start-2">
        {timerDetails ? (
          <div className="row-start-2">
            <h1 className="text-4xl">{timerDetails.timer.title}</h1>
            <p className="text-xl">Add a description</p>
            <button className="cursor-pointer text-xl">Delete</button>
          </div>
        ) : (
          <p>Loading timer...</p>
        )}
      </main>
    </div>
  );
}
