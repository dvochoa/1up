"use client";

import { useState } from "react";
import Image from "next/image";

interface CreateTimerProps {
  onTimerCreated: () => void;
}

const ICON_WIDTH = 24;
const ICON_HEIGHT = 24;

const CreateTimer: React.FC<CreateTimerProps> = ({ onTimerCreated }) => {
  const [title, setTitle] = useState("");
  const [isCreating, setIsCreating] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault(); // Prevent page refresh since update is handled on the client
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
        setIsCreating(false);
        onTimerCreated();
      } else {
        // TODO: Handle error appropriately
        console.error("Failed to create timer");
      }
    } catch (error) {
      console.error("Error creating timer:", error);
    }
  };

  // TODO: Can I make the transition look better?
  return (
    <div>
      <div
        className={`transform transition-all duration-300 ease-in-out ${
          isCreating ? "pointer-events-none -translate-y-4 opacity-0" : "translate-y-0 opacity-100"
        }`}
      >
        <button
          onClick={() => setIsCreating(true)}
          className="flex items-center justify-center rounded-full bg-zinc-500/40 p-3 text-white transition-colors duration-200 hover:bg-zinc-500/60 focus:ring-2 focus:ring-zinc-500 focus:ring-offset-2 focus:outline-none dark:bg-zinc-500/40 dark:hover:bg-zinc-500/60"
        >
          <Image
            src="/images/plus-icon.svg"
            width={ICON_WIDTH}
            height={ICON_HEIGHT}
            alt="Add new timer"
            className="dark:invert"
          />
        </button>
      </div>

      <div
        className={`transform transition-all duration-300 ease-in-out ${
          isCreating ? "translate-y-0 opacity-100" : "pointer-events-none translate-y-4 opacity-0"
        }`}
      >
        <form onSubmit={handleSubmit}>
          <div className="flex gap-2">
            <input
              type="text"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              placeholder="Enter timer name"
              className="flex-1 rounded-sm border border-gray-400 bg-zinc-500/40 px-3 py-2 text-sm transition-colors duration-200 placeholder:text-neutral-400 focus:ring-1 focus:ring-zinc-500 focus:outline-none dark:border-gray-600 dark:text-white"
              autoFocus
            />
            <button
              type="submit"
              className="rounded-sm bg-zinc-500/40 px-4 py-2 text-sm font-medium text-neutral-800 transition-colors duration-200 hover:bg-zinc-500/60 focus:ring-2 focus:ring-zinc-500 focus:ring-offset-2 focus:outline-none dark:text-neutral-300"
            >
              Create
            </button>
            <button
              type="button"
              onClick={() => setIsCreating(false)}
              className="rounded-sm bg-zinc-500/40 px-4 py-2 text-sm font-medium text-neutral-800 transition-colors duration-200 hover:bg-zinc-500/60 focus:ring-2 focus:ring-zinc-500 focus:ring-offset-2 focus:outline-none dark:text-neutral-300"
            >
              Cancel
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default CreateTimer;
