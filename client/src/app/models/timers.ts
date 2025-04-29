export interface Timer {
  id: number;
  title: string;
  totalTimeInSeconds: number;
}

export interface TimerDetails {
  timer: Timer;
  timerSessions: TimerSession[];
}

interface TimerSession {
  id: number;
  createdAd: Date;
  sessionDurationInSeconds: number;
}
