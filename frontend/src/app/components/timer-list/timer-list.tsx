import { Timer, TimerProps } from "@/components/timer/timer";

interface TimerListProps {
  timers: TimerProps[];
}

const TimerList: React.FC<TimerListProps> = ({ timers }) => {
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
    </div>
  );
};

export default TimerList;
