import TimerList from "@/components/timer-list/timer-list";

export default function HomePage() {
  const timers = [
    { title: "Coding", id: 1, backgroundColor: "bg-indigo-500/50" },
    { title: "Music Production", id: 2, backgroundColor: "bg-red-500/50" },
    { title: "DJing", id: 3, backgroundColor: "bg-green-500/50" },
    { title: "Piano", id: 4, backgroundColor: "bg-orange-500/50" },
  ];

  return (
    <div>
      <meta charSet="utf-8"></meta>
      <meta name="viewport" content="width=device-width, initial-scale=1.0"></meta>
      <main>
        <TimerList timers={timers}></TimerList>
      </main>
    </div>
  );
}
