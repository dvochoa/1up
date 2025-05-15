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

  // TODO: Use a suspense or something when doesn't load?
  // TODO: style button

  return (
    <div className={`${styles["page-grid-container"]} h-[85vh]`}>
      <main className={`col-start-2 row-start-2`}>
        {timerDetails ? (
          <div className={`${styles["timer-details"]} gap-x-5`}>
            {/* Timer description + delete column */}
            <div>
              <h1 className="text-4xl">{timerDetails.timer.title}</h1>
              <p className="text-xl">
                Working on any one of my various side projects, reading books related to software,
                dev, learning about topics I am unfamiliar with, trying out new tech, etc
              </p>
              <button className="cursor-pointer text-xl">Delete</button>
            </div>

            {/* Timer stats column */}
            <div className="col-start-2">
              <p className="text-xl">
                Working on any one of my various side projects, reading books related to software,
                dev, learning about topics I am unfamiliar with, trying out new tech, etc
              </p>
            </div>
          </div>
        ) : (
          <p>Loading timer...</p>
        )}
      </main>
    </div>
  );
}
