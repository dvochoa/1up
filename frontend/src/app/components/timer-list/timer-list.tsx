import { Timer, TimerProps } from "@/components/timer/timer";
import { ReactNode } from "react";

interface TimerListProps {
  timers: TimerProps[];
  children: ReactNode;
}

const TimerList: React.FC<TimerListProps> = ({ timers, children }) => {
  return (
    <div className="m-10 rounded border-solid border-2 border-red-700">
      {timers.map((timer) => (
        <Timer
          title={timer.title}
          id={timer.id}
          key={timer.id}
          backgroundColor={timer.backgroundColor}
        ></Timer>
      ))}
      {children}
    </div>
  );
};

export default TimerList;
