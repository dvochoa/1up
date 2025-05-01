"use client";

import { useState } from "react";
import Image from "next/image";
import { NewTimer } from "./new-timer";

interface CreateTimerProps {
  onTimerCreated: () => void;
}

const ICON_WIDTH = 24;
const ICON_HEIGHT = 24;

const CreateTimer: React.FC<CreateTimerProps> = ({ onTimerCreated }) => {
  const [isCreating, setIsCreating] = useState(false);

  const handleTimerCreated = async () => {
    onTimerCreated();
    setIsCreating(false);
  };

  return (
    <div className="relative flex flex-col">
      <div
        className={`absolute inset-0 flex items-center transition-opacity duration-300 ease-in-out ${
          isCreating ? "pointer-events-none opacity-0" : "opacity-100"
        }`}
      >
        <button
          onClick={() => setIsCreating(true)}
          className="cursor-pointer rounded-full bg-zinc-500/40 p-3 text-white transition-colors duration-200 hover:bg-zinc-500/60 focus:ring-2 focus:ring-zinc-500 focus:ring-offset-2 focus:outline-none dark:bg-zinc-500/40 dark:hover:bg-zinc-500/60"
        >
          <Image
            src="/images/new-timer-icon.svg"
            width={ICON_WIDTH}
            height={ICON_HEIGHT}
            alt="Add new timer"
            className="dark:invert"
          />
        </button>
      </div>

      <div
        className={`h-full transition-opacity duration-300 ease-in-out ${
          isCreating ? "opacity-100" : "pointer-events-none opacity-0"
        }`}
      >
        <NewTimer
          onTimerCreated={handleTimerCreated}
          onTimerDeleted={() => setIsCreating(false)}
        ></NewTimer>
      </div>
    </div>
  );
};

export default CreateTimer;
