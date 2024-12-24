import TimerList from "@/components/timer-list/timer-list";

export default function HomePage() {
  const timers = [
    { title: "Coding" },
    { title: "Music Production" },
    { title: "DJing" },
    { title: "Piano" },
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
