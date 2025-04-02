import { Timer, TimerProps } from "@/components/timer";

interface TimerListProps {
  timers: TimerProps[];
}

const TimerList: React.FC<TimerListProps> = ({ timers }) => {
  return (
    <div className="space-y-5">
      {timers.map((timer) => (
        <Timer
          key={timer.id}
          id={timer.id}
          title={timer.title}
          totalTimeInSeconds={timer.totalTimeInSeconds}
        ></Timer>
      ))}
    </div>
  );
};

export default TimerList;
