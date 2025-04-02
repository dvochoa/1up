import { Timer, TimerProps } from "@/components/timer";

interface TimerListProps {
  timers: TimerProps[];
}

// TODO: TimerList should enforce the my-5 spacing in between timers, it shouldn't be a part of timers themselves
const TimerList: React.FC<TimerListProps> = ({ timers }) => {
  return (
    <div>
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
