import { Timer, TimerProps } from "@/components/timer";
import CreateTimer from "@/components/create-timer";

interface TimerListProps {
  timers: TimerProps[];
  onTimerCreated: () => void;
  onTimerDeleted: () => void;
}

const TimerList: React.FC<TimerListProps> = ({ timers, onTimerCreated, onTimerDeleted }) => {
  return (
    <div className="grid gap-y-5 lg:grid-flow-col lg:grid-cols-4 lg:grid-rows-5 lg:gap-x-25 lg:gap-y-10 2xl:grid-cols-5 2xl:grid-rows-6">
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
