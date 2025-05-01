"use client";

import styles from "./timer.module.css";
import React, { useState } from "react";
import Image from "next/image";

const ICON_WIDTH = 18;
const ICON_HEIGHT = 18;

interface NewTimerProps {
  onTimerCreated: () => void;
  onTimerDeleted: () => void;
}

export const NewTimer: React.FC<NewTimerProps> = ({ onTimerCreated, onTimerDeleted }) => {
  const [title, setTitle] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault(); // Prevent page refresh since update is handled on the client
    // TODO: Show the user some feedback if this is hit
    if (!title.trim()) return;

    try {
      const response = await fetch("/api/timers", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ ownerId: 1, title: title }),
      });

      if (response.ok) {
        setTitle("");
        onTimerCreated();
      } else {
        // TODO: Handle error appropriately
        console.error("Failed to create timer");
      }
    } catch (error) {
      console.error("Error creating timer:", error);
    }
  };

  return (
    <div className={`h-full rounded-sm bg-zinc-500/40 shadow-md dark:shadow-none`}>
      <div className={`${styles["grid-container"]} h-full p-2 pb-3 pl-3`}>
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            maxLength={10}
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            placeholder="Timer name"
            className="col-1 row-1 mr-3 text-xl outline-hidden"
          ></input>
        </form>

        <button className="col-2 row-1 ml-auto cursor-pointer" onClick={onTimerDeleted}>
          <Image
            src="/images/delete-timer-icon.svg"
            width={ICON_WIDTH}
            height={ICON_HEIGHT}
            alt="A trash can, on click this timer creation form is removed"
            className="dark:invert"
          />
        </button>

        <button className="col-2 row-9 ml-auto cursor-pointer" onClick={handleSubmit}>
          <Image
            src="/images/commit-session-icon.svg"
            width={ICON_WIDTH}
            height={ICON_HEIGHT}
            alt="An arrow, on click a new timer is created"
            className="dark:invert"
          />
        </button>
      </div>
    </div>
  );
};
