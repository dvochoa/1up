import { Timer, TimerProps } from "@/components/timer";
import CreateTimer from "@/components/create-timer";

interface TimerListProps {
  timers: TimerProps[];
  onTimerCreated: () => void;
  onTimerDeleted: () => void;
}

const TimerList: React.FC<TimerListProps> = ({ timers, onTimerCreated, onTimerDeleted }) => {
  return (
    <div className="grid grid-flow-col grid-cols-5 grid-rows-6 gap-x-25 gap-y-10">
      {timers.map((timer) => (
        <Timer
          key={timer.id}
          id={timer.id}
          title={timer.title}
          totalTimeInSeconds={timer.totalTimeInSeconds}
          onTimerDeleted={onTimerDeleted}
        ></Timer>
      ))}
      <CreateTimer onTimerCreated={onTimerCreated} />
    </div>
  );
};

export default TimerList;
