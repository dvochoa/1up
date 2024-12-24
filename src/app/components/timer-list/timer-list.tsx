import { Timer, TimerProps } from "@/components/timer/timer";

interface TimerListProps {
  timers: TimerProps[];
}

// TODO: Fix "Each child in a list should have a unique "key" prop."
// TODO: Add some more preliminary stylings
//  - More left marging
//  - More spacing in between timers
const TimerList: React.FC<TimerListProps> = ({ timers }) => {
  return (
    <div>
      {timers.map((timer) => (
        <Timer title={timer.title}></Timer>
      ))}
    </div>
  );
};

export default TimerList;
